str = document.querySelector('#str')
img = document.querySelector('#img')
timer = null
latest = ''
str.addEventListener('keyup', () ->
  if (str.value.trim().length >= 3 && str.value.trim() != latest)
    img.innertHTML = ''
    clearTimeout timer
    timer = setTimeout(() ->
      request = new XMLHttpRequest()
      request.open('POST', '/qrcode/svg', true)
      request.setRequestHeader(
        'Content-Type',
        'application/x-www-form-urlencoded; charset=UTF-8'
      )
      latest = str.value.trim()
      request.send("str=" + latest)

      request.onload = () ->
        if (request.status >= 200 && request.status < 400)
          resp = request.responseText
          img.innerHTML = resp
        else window.alert 'Erro!'
        return
    , 2000)
  return
)
