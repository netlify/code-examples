# Token hider

This is an example of using [Netlify Functions](https://www.netlify.com/docs/functions/) to hide the API token/key for API calls using [Google Maps](https://developers.google.com/maps/documentation/javascript/tutorial).

You may want to do this because you don't want every user to see your API key. If they had that, they could make malicious use of Google Maps API, while pretending to be you.

[Live demo](https://hzdf-maps.netlify.com/)

## Configuration

When you deploy your site to Netlify, you'll need to define the following [build environment variables](https://www.netlify.com/docs/continuous-deployment/#build-environment-variables) in the UI.

* `API_URL` = "https://maps.googleapis.com/maps/api/geocode/json"

* `API_TOKEN` = "your Google Maps API key" (To get a Google Maps Web API key, visit [here](https://developers.google.com/maps/documentation/javascript/get-api-key))

## Localhost

### 1. Installation
Clone the repo, from the command line, go to `function_examples/token-hider` folder and run `npm install`.

### 2. Run Lambda Functions
`API_URL="https://maps.googleapis.com/maps/api/geocode/json" API_TOKEN="your Google Maps API key" npm run lambda:dev`
This will start a local dev server. Requests to `http://localhost:9000/getapi?address=<the address>` will now be handled by `functions/getapi.js`.
e.g. try this URL in your browser:
`localhost:9000/getapi?address=220+Stanhope+St,+Brooklyn,+NY`

#### Proxying for local development
Examine the `server.js` file. It is already configured to proxy all requests destined for `/.netlify/functions` to the server running on port 9000.

### 3. Test the web site
Open a new console and type in: `npm run site:dev`
The site is now available at `http://localhost:1234/`

_Notes: Make sure you run Functions first before you run the site._

Add a query to retrieve a specific data object, e.g.:
`localhost:9000/getapi?address=220+Stanhope+St,+Brooklyn,+NY`

Open the site in your browser and verify that the API key is not anywhere on the client side.