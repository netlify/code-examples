# code-examples

Code snippets for Netlify users

# Netlify AWS Lambda Function Examples

These next functions are in subfolders, but are meant to be deployed from their own repo. They each include their own netlify.toml files.

## rest

This is a basic example of a function that when called, makes a request to a rest API using node-fetch, and returns a portion of the response back to the client. It includes an index.html that makes the call and displays the response

## random-word

This is a basic example of a function that when called, returns a random word from an array of words back to the user. This function has no build process and does not use any external packages

## use-env

This is a basic example of a function that uses a build environment variable that's specified in the users netlify site's deploy settings.
