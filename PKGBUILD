Maintainer="Barış İnandıoğlu"
pkgname=barley
pkgver=r32.5e864b2
pkgrel=1
epoch=
pkgdesc="A wrapper for touch that generates boilerplate files."
arch=(x86_64)
url="https://github.com/baris-inandi/boil"
license=('GPLv2')
groups=()
depends=()
makedepends=(git go)
checkdepends=()
optdepends=()
provides=(barley)
conflicts=()
replaces=()
backup=()
options=()
install=
changelog=
source=("git+$url")
noextract=()
md5sums=('SKIP')
validpgpkeys=()

pkgver() {
    cd $pkgname
    printf "${_pkgver}r%s.%s" "$(git rev-list --count HEAD)" "$(git rev-parse --short HEAD)"
}

build() {
    cd $pkgname
    make
}

package() {
    cd $pkgname
    make install
}
