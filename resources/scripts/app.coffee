str = document.querySelector('#str')
img = document.querySelector('#img')
alert = document.createElement('div')
alert.className = "alert alert-success"
lnk = document.createElement('a')
lnk.className += ' alert-link'
alert.appendChild(lnk)
timer = null
latest = ''
str.addEventListener('keyup', () ->
  if (str.value.trim().length >= 3 && str.value.trim() != latest)
    img.innertHTML = ''
    lnk.innertHTML = ''
    clearTimeout timer
    timer = setTimeout(() ->
      latest = str.value.trim()
      root = window.location + 'qrcode/svg/'
      link = root + encodeURIComponent(latest)
      lnk.setAttribute('href', link)
      lnk.innerHTML = root + latest
      img.parentNode.insertBefore(alert, img)
      img.innerHTML = '<img src="' + link + '">'
    , 2000)
  return
)
