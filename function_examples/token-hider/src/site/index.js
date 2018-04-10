const api = {
  encode: () => {
    const address = document.getElementById("address")
    const encoded = address.value.split(' ').join('+')
    const params = "?address=" + encoded
    return params
  },
  send: () => {
    encode = api.encode
    api.get(encode)
  },
  show: (result) => {
    const location = result[0].geometry.location
    const coordinates = document.getElementById("coordinates")
    const map = document.getElementById("map")
    console.log(JSON.stringify(result))
    map.innerHTML = JSON.stringify(location)
    map.hidden = false
  },
  get: (params) => {
  fetch("/.netlify/functions/getapi" + api.encode())
    .then(response => response.json())
    .then(result => api.show(result.results))
    .catch(err => console.log(err))
  },
}

//Create eventlistener
const submit = document.getElementById('submit')
submit.addEventListener('click', api.send, false)
