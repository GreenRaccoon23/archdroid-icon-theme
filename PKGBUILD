# Maintainer: GreenRaccoon23 <GreenRaccoon a t gmail d o t com>

pkgname=archdroid-icon-theme
pkgver=r51.c00b07b
pkgrel=1
pkgdesc="Port of Android 5.0 Lollipop's material design icons to Arch."
arch=('any')
url="https://github.com/GreenRaccoon23/${pkgname}"
license=('GPL3')
makedepends=('intltool' 'librsvg' 'gtk-update-icon-cache')
provides=("${pkgname}")
conflicts=("${pkgname}")
options=('!strip')
install="${pkgname}.install"
source=("https://github.com/GreenRaccoon23/${pkgname}/raw/master/${pkgname}.tar.xz")
md5sums=("0ec9ea8ebde4ab59df15797d4b86b584")

package() {
	msg2 "Installing ${pkgname}..." ;
	cd ${pkgname} ;
  	install -dm 755 "${pkgdir}"/usr/share/icons
  	cp -drf --no-preserve='ownership' . "${pkgdir}"/usr/share/icons/
}

