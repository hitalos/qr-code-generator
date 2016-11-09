const express = require('express')

const router = express.Router()

/* GET home page. */
router.get('/', (req, res, next) => {
  res.render('index', { title: 'QR code generator' })
})

module.exports = router
