const coffeeMiddleware = require('coffee-middleware')
const compression = require('compression')
const express = require('express')
const sassMiddleware = require('node-sass-middleware')
const qr = require('qr-image')

process.on('uncaughtException', console.error)

const index = require('./routes/index')

const app = express()
app.use(compression())

// view engine setup
app.set('views', './resources/views')
app.set('view engine', 'pug')
app.use(sassMiddleware({
  src: 'resources',
  dest: './public',
  debug: false,
  indentedSyntax: true,
  outputStyle: 'compressed'
}))
app.use(coffeeMiddleware({
  compress: true,
  debug: false,
  src: 'resources'
}))

app.use(express.static('./public'))

app.use('/', index)
app.get('/qrcode/svg/:str', (req, res) => {
  res.header({ 'Content-Type': 'image/svg+xml' })
  res.end(qr.imageSync(req.params.str, { type: 'svg' }))
})
app.get('/qrcode/png/:str', (req, res) => {
  res.header({ 'Content-Type': 'image/png' })
  res.end(qr.imageSync(req.params.str))
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
