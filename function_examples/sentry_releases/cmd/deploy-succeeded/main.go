package main

import (
	"bytes"
	"encoding/json"
	"os"

	sentry "github.com/atlassian/go-sentry-api"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type deploy struct {
	CommitRef string `json:"commit_ref"`
	Context   string `json:"context"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var d deploy

	if err := json.NewDecoder(bytes.NewReader([]byte(request.Body))).Decode(&d); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	if d.Context != "production" {
		// Ignore deploys that don't go to production
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       "I ran after a deploy was created",
		}, nil
	}

	c, err := sentry.NewClient(os.Getenv("SENTRY_AUTH_TOKEN"), nil, nil)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	orgSlug := os.Getenv("SENTRY_ORG_SLUG")
	org := sentry.Organization{
		Slug: &orgSlug,
	}

	proSlug := os.Getenv("SENTRY_PROJECT_SLUG")
	pro := sentry.Project{
		Slug: &proSlug,
	}

	rel := sentry.NewRelease{
		Version: d.CommitRef,
		Ref:     &d.CommitRef,
	}

	if _, err := c.CreateRelease(org, pro, rel); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "I ran after a deploy was created",
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
