# Netlify Functions Code Examples

The following examples are under `function_examples` folder. Each example showcases the basic usage of [Netlify Functions](https://www.netlify.com/docs/functions/) with slightly different approaches.

## Token hider

An example of using Functions to hide API key during API calls. [Learn more](https://github.com/netlify/code-examples/tree/update-readme/function_examples/token-hider).

## rest

This is a basic example of a function that when called, makes a request to a rest API using node-fetch, and returns a portion of the response back to the client. It includes an index.html that makes the call and displays the response. [Learn more](https://github.com/netlify/code-examples/tree/update-readme/function_examples/rest)

## random-word

This is a basic example of a function that when it gets called, returns a random word from an array of words back. This function has no build process and does not use any external packages. [Learn more](https://github.com/netlify/code-examples/tree/update-readme/function_examples/random-word).

## use-env

This is a basic example of a function that uses a build environment variable that's specified in the users Netlify site's deploy settings. [Learn more](https://github.com/netlify/code-examples/tree/update-readme/function_examples/use-env).
