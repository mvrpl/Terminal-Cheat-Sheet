gover="1.8"
chshtver="0.1.4"
find . -regex ".*chsht-[0-9]\.[0-9]\.[0-9].*" | xargs rm -Rf
yum install rpmdevtools wget gcc -y
echo '%_topdir %(pwd)' > ~/.rpmmacros
wget https://storage.googleapis.com/golang/go${gover}.linux-amd64.tar.gz
tar -xvf go${gover}.linux-amd64.tar.gz -C /usr/local
mkdir ~/.go
export GOPATH=~/.go
export GOROOT=/usr/local/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
go get github.com/mattn/go-sqlite3
go get golang.org/x/text/runes
cd /Terminal-Cheat-Sheet
go build -o chsht
mkdir chsht-${chshtver}
mv chsht chsht-${chshtver}
tar cvfz chsht-${chshtver}.tar.gz chsht-${chshtver}/
mv chsht-${chshtver}.tar.gz /Terminal-Cheat-Sheet/RPMBUILD/SOURCES
cd /Terminal-Cheat-Sheet/RPMBUILD
rpmbuild -ba SPECS/chsht.spec
rm -Rf go${gover}.linux-amd64.tar.gz
