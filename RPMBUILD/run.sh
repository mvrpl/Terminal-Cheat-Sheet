yum install rpmdevtools -y
echo '%_topdir %(pwd)' > ~/.rpmmacros
rpmbuild -ba SPECS/chsht.spec
