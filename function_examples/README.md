# Netlify Functions code examples

## Basic examples

### [Random Word](https://github.com/netlify/code-examples/tree/master/function_examples/random-word)

A basic usage case that triggers a function to generate a random word. No dependencies or build tools.

### [Use Env](https://github.com/netlify/code-examples/tree/master/function_examples/use-env)

An example that uses Netlify build environment variables, with no dependencies or build tools.

### [Rest](https://github.com/netlify/code-examples/tree/master/function_examples/rest)

A function for calling an external API with `node-fetch`, using `netlify-lambda` to bundle dependencies.

### [Token Hider](https://github.com/netlify/code-examples/tree/master/function_examples/token-hider)

A function that uses `axios` to access an external API and uses Netlify build environment variables to hide the API token; bundled with `netlify-lambda`.

## Further exploration

For a more comprehensive tutorial, this enlightening [example](https://github.com/alexmacarthur/netlify-lambda-function-example) will walk you through the thought process step by step. Thanks to our user [alexmacarthur](https://github.com/alexmacarthur/netlify-lambda-function-example/commits?author=alexmacarthur).

If you're curious about integrating Functions with other Netlify add-on features, such as [Identity](https://www.netlify.com/docs/identity/) and [Form handling](https://www.netlify.com/docs/form-handling/), Checkout these tutorials:

* [JAMstack architecture on Netlify: How Identity and Functions work together](https://www.netlify.com/blog/2018/03/29/jamstack-architecture-on-netlify-how-identity-and-functions-work-together/), which is based on our [Create React App Lambda](https://github.com/netlify/create-react-app-lambda) boiler plate.

* [Creating your own URL shortener with Netlify Forms and Functions](https://medium.com/netlify/creating-your-own-url-shortener-with-netlify-forms-and-functions-a4286f55ea6c)
