const express = require('express')
const qr = require('qr-image')

const router = express.Router()

router.get('/:type(png|svg)/:str', (req, res) => {
  let ct = 'image/svg+xml'
  if (req.params.type === 'png') ct = 'image/png'
  res.header({ 'Content-Type': ct })
  res.end(qr.imageSync(req.params.str, { type: req.params.type }))
})

module.exports = router
