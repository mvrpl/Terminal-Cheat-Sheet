yum install rpmdevtools
echo '%_topdir %(pwd)' > ~/.rpmmacros
rpmbuild -ba SPECS/chsht.spec
