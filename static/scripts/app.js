const str = document.querySelector('#str')
const img = document.querySelector('#img')
const alert = document.createElement('div')
const lnk = document.createElement('a')

alert.classList.add('alert')
alert.appendChild(lnk)

let timer
let latest = ''

const getImage = () => {
	if (str.value.trim().length >= 3 && str.value.trim() !== latest) {
		clearTimeout(timer)
		timer = setTimeout(() => {
			latest = str.value.trim()
			const root = `${window.location}qrcode/png/`
			const link = root + encodeURIComponent(latest)
			lnk.setAttribute('href', link)
			lnk.innerHTML = root + latest
			img.parentNode.insertBefore(alert, img)
			const png = document.createElement('img')
			png.setAttribute('src', link)
			img.textContent = ''
			img.appendChild(png)
		}, 1000)
	}
}

str.addEventListener('keyup', getImage)
str.addEventListener('blur', getImage)
