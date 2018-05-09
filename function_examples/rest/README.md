# Rest

This is a basic example of a function that when called, makes a request to a REST API using [`node-fetch`](https://www.npmjs.com/package/node-fetch), and returns a portion of the response back to the client.

During the build process, [`netlify-lambda`](https://www.netlify.com/docs/functions/#tools-for-building-javascript-functions) takes the source files and package.json from `/func` and builds the function as a single file into a new folder, `/functions`.

In the `index.html` file, a click event fires up the lambda function and displays the response object in the browser.

Data source: https://jsonplaceholder.typicode.com/.
