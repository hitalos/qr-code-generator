const str = document.querySelector('#str')
const img = document.querySelector('#img')
const alert = document.createElement('div')
const lnk = document.createElement('a')

alert.classList.add('alert')
alert.appendChild(lnk)

let timer
let latest = ''

str.addEventListener('keyup', () => {
  if (str.value.trim().length >= 3 && str.value.trim() !== latest) {
    img.innertHTML = ''
    lnk.innertHTML = ''
    clearTimeout(timer)
    timer = setTimeout(() => {
      latest = str.value.trim()
      const root = `${window.location}qrcode/svg/`
      const link = root + encodeURIComponent(latest)
      lnk.setAttribute('href', link)
      lnk.innerHTML = root + latest
      img.parentNode.insertBefore(alert, img)
      img.innerHTML = `<img src="${link}">`
    }, 2000)
  }
})
