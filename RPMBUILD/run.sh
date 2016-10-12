sudo yum install rpmdevtools
echo '%_topdir %(echo "$HOME")/RPMBUILD' > ~/.rpmmacros
rpmbuild -ba SPECS/chsht.spec
