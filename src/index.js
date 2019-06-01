var express = require('express')

var app = express()

app.get('/', (req, rsp) => {
  rsp.send('OK')
})

app.listen('3000', () => console.log('ok!'))