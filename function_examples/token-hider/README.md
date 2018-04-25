# Token Hider

A token/key obscuring function for API calls using Netlify Functions.
Go to [Functions documentation](https://www.netlify.com/docs/functions/) for detailed instructions about Netlify Functions.

### Example

https://hzdf-maps.netlify.com/

## Get your own

You can clone this repo to your own Github account and create a new site on Netlify to make your own by clicking the _Deploy to Netlify button_ below.

[![Deploy to Netlify](https://www.netlify.com/img/deploy/button.svg)](https://app.netlify.com/start/deploy?repository=https://github.com/depadiernos/token-hider)

After you have created a new site on Netlify via the button above, you will then need to perform the configuration steps below.

## Configuration

You'll need to define the following Environment Variables in the Netlify UI to make the example work.

* `API_URL` = "https://maps.googleapis.com/maps/api/geocode/json"

* `API_TOKEN` = "<your Google Maps API key>"(To get a Google Maps Web API key, visit [here](https://developers.google.com/maps/documentation/javascript/get-api-key))

## Testing locally

Follow the steps below to test the example locally.

### 1. Run Lambda Functions

Type in the following script in your console:
`export API_URL="https://maps.googleapis.com/maps/api/geocode/json" && export API_TOKEN="<your Google Maps API token>" && npm run lambda:dev`
This will start a local dev server. Requests to `http://localhost:9000/getapi` will route to `functions/getapi.js`.

### 2. Run site

Open a new console and type in: `npm run site:dev`
The site is now available at `http://localhost:1234/`

_Notice: Make sure you run Lambda Functions first before you run the site._

Add a query to retrieve a specific data object, e.g.:
`localhost:9000/getapi?address=220+Stanhope+St,+Brooklyn,+NY`
