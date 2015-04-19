# Maintainer: GreenRaccoon23 <GreenRaccoon a t gmail d o t com>

pkgname=archdroid-icon-theme
pkgver=r47.9e31441
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
md5sums=("cc6d2b0bead603f74bfa73fb60411f14")

package() {
	msg2 "Installing ${pkgname}..." ;
	cd ${pkgname} ;
  	install -dm 755 "${pkgdir}"/usr/share/icons
  	cp -drf --no-preserve='ownership' . "${pkgdir}"/usr/share/icons/
}

