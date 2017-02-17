yum install rpmdevtools wget gcc -y
echo '%_topdir %(pwd)' > ~/.rpmmacros
wget https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz
tar -xvf go1.8.linux-amd64.tar.gz -C /usr/local
mkdir ~/.go
export GOPATH=~/.go
export GOROOT=/usr/local/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
go get github.com/mattn/go-sqlite3
go get golang.org/x/text/runes
cd /Terminal-Cheat-Sheet
go build -o chsht
tar cvfz chsht-0.1.4.tar.gz chsht
mv chsht-0.1.4.tar.gz /Terminal-Cheat-Sheet/RPMBUILD/RPMS/x86_64
cd /Terminal-Cheat-Sheet/RPMBUILD
rpmbuild -ba SPECS/chsht.spec
