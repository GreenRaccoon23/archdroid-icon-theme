# Maintainer: GreenRaccoon23 <GreenRaccoon a t gmail d o t com>

pkgname=archdroid-icon-theme
pkgver=r19.63668ef
pkgrel=1
pkgdesc="Port of Android 5.0 Lollipop's material design icons to Arch."
arch=('any')
url="https://github.com/GreenRaccoon23/${pkgname}"
license=('GPL3')
makedepends=('intltool' 'librsvg' 'gtk-update-icon-cache' 'xz')
provides=("${pkgname}")
conflicts=("${pkgname}")
#options=('!strip')
install="${pkgname}.install"
source=("https://github.com/GreenRaccoon23/${pkgname}/raw/master/${pkgname}.tar.xz")
md5sums=("cc0f4e8c0c3b083d87189670cc3dcaa3")

package() {
	cd ${pkgname}
  	install -dm 755 "$pkgdir"/usr/share/icons
  	cp -drf --no-preserve='ownership' . "$pkgdir"/usr/share/icons/
}

