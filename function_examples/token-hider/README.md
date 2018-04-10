# token-hider
A token/key obscuring function for API calls using Netlify functions.

To test: /.netlify/functions/passthrough?address=4640+Hensley+Dr,+Cookeville,+TN

## requirements

2 ENV VARs need to be set in the UI (https://app.netlify.com/sites/hzdf-maps/settings/deploys#build-environment-variables) for this example to work.

API_URL = https://maps.googleapis.com/maps/api/geocode/json

API_TOKEN = <your Google Maps API key here>
  
To get a Maps Web API key, visit here: https://developers.google.com/maps/documentation/javascript/get-api-key
