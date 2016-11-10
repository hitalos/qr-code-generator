str = document.querySelector('#str')
img = document.querySelector('#img')
lnk = document.createElement('a')
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
      link = root + encodeURI(latest)
      img.parentNode.insertBefore(lnk, img)
      img.innerHTML = '<img src="' + link + '">'
    , 2000)
  return
)
