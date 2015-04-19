# Maintainer: GreenRaccoon23 <GreenRaccoon a t gmail d o t com>

pkgname=archdroid-icon-theme
pkgver=r49.db29455
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
md5sums=("01621aef6f4a6b3eb3a6f88679f65cf1")

package() {
	msg2 "Installing ${pkgname}..." ;
	cd ${pkgname} ;
  	install -dm 755 "${pkgdir}"/usr/share/icons
  	cp -drf --no-preserve='ownership' . "${pkgdir}"/usr/share/icons/
}

