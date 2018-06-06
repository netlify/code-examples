# Sentry release example

This is an example of using [Netlify Functions](https://www.netlify.com/docs/functions/) to create a release in Sentry every time Netlify publishes a new Production deploy.

## Configuration

When you deploy your site to Netlify, you'll need to define the following [build environment variables](https://www.netlify.com/docs/continuous-deployment/#build-environment-variables) in the UI.

1- Authentication token to talk with Sentry's API:

* `SENTRY_AUTH_TOKEN` = "secret token issued by sentry"

2- Your Sentry's organization slug:

* `SENTRY_ORG_SLUG` = "netlify"

3- Your Sentry's project slug:

* `SENTRY_PROJECT_SLUG` = "code-examples"
