const fetch = require('node-fetch');

exports.handler = function(event, context, callback) {
  const respond = ({ status, body }) => {
    callback(null, {
      statusCode: status,
      body: JSON.stringify({ body }),
    });
  };

  (() => {
    fetch('https://jsonplaceholder.typicode.com/posts/1')
      .then(response => response.json())
      .then(json => {
        respond({ status: 200, body: json.body });
      })
      .catch(err => {
        respond({ status: 422, body: "Couldn't get the data" });
      });
  })();
};
