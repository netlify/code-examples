exports.handler = function(event, context, callback) {
  // Currently this only works with variables that are defined in the UI and not in the netlify.toml
  // Using ES6 destructuring to pull the variable called FAKE_TOKEN out of the process.env object
  // You can then use the FAKE_TOKEN anywhere in your code just as you would any string variable
  const { FAKE_TOKEN } = process.env;
  console.log(FAKE_TOKEN);
  // Amazon callback function without any custom headers
  callback(null, {
    statusCode: 200,
    body: JSON.stringify({ token: FAKE_TOKEN }),
  });
};
