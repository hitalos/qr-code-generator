const express = require('express')
const qr = require('qr-image')
const logger = require('pino')()

process.on('uncaughtException', logger.error)

const app = express()

// view engine setup
app.set('views', './resources/views')
app.set('view engine', 'pug')

app.use(express.static('./public'))

app.get('/', (_, res) => {
  res.render('index', { title: 'QR code generator' })
})

app.get('/qrcode/:type(png|svg)/:str', (req, res) => {
  let ct = 'image/svg+xml'
  if (req.params.type === 'png') ct = 'image/png'
  res.header({ 'Content-Type': ct })
  res.end(qr.imageSync(req.params.str, { type: req.params.type }))
})

// catch 404 and forward to error handler
app.use((req, res, next) => {
  const err = new Error('Not Found')
  err.status = 404
  next(err)
})

app.use((err, req, res, next) => {
  res.status(err.status || 500)
  if (app.get('env') === 'development') {
    res.render('error', { err })
  } else {
    res.render('error', { error: err.message })
  }
  next()
})

module.exports = app
