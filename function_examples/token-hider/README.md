# Token Hider

This is an example of hiding the API token/key for API calls using [Google Maps API](https://developers.google.com/maps/documentation/javascript/tutorial) and [Netlify Functions](https://www.netlify.com/docs/functions/).

[Live demo](https://hzdf-maps.netlify.com/)

## Configuration

When you deploy your site to Netlify, you'll need to define the following Environment Variables in the UI.

* `API_URL` = "https://maps.googleapis.com/maps/api/geocode/json"

* `API_TOKEN` = "your Google Maps API key" (To get a Google Maps Web API key, visit [here](https://developers.google.com/maps/documentation/javascript/get-api-key))

## Testing locally

### 1. Install

Clone the repo and go to `function_examples/token-hider` and `npm install`.

### 2. Run Lambda Functions

Type in the following script in your console:
`export API_URL="https://maps.googleapis.com/maps/api/geocode/json" && export API_TOKEN="<your Google Maps API token>" && npm run lambda:dev`
This will start a local dev server. Requests to `http://localhost:9000/getapi` will route to `functions/getapi.js`.

### 3. Run site

Open a new console and type in: `npm run site:dev`
The site is now available at `http://localhost:1234/`

_Notes: Make sure you run Functions first before you run the site._

Add a query to retrieve a specific data object, e.g.:
`localhost:9000/getapi?address=220+Stanhope+St,+Brooklyn,+NY`
