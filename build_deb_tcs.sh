#!/bin/sh

# Export DEB variables
export DEBEMAIL="mvrpl@icloud.com"
export DEBFULLNAME="Marcos Lima"

SOURCEBINPATH=~/Documentos/Terminal-Cheat-Sheet/compiled/Ubuntu64 # Set path of bin
SOURCEBIN=chsht # Bin name
DEBFOLDER=~/chsht-data # Package directory in user HOME
DEBVERSION=2.0 # Set version of package
DEBFOLDERNAME=$DEBFOLDER-$DEBVERSION # Complete packge folder name in user HOME

# Create your scripts source dir
mkdir $DEBFOLDERNAME

# Copy your script to the source dir
cp $SOURCEBINPATH/$SOURCEBIN $DEBFOLDERNAME 
cd $DEBFOLDERNAME

# Create the packaging skeleton (debian/*)
dh_make --native

# Remove make calls
grep -v makefile debian/rules > debian/rules.new 
mv debian/rules.new debian/rules 

# debian/install must contain the list of scripts to install 
# as well as the target directory
echo $SOURCEBIN usr/bin > debian/install 

# Remove the example files
rm debian/*.ex

# Build package
dpkg-buildpackage -us -uc

# Build .deb with GPG signature
debuild -S -k$(gpg --list-secret-keys --keyid-format LONG | awk '{if(index($0, "sec   ") != 0) {print gensub(/(sec).*\/(.*)\s.*/, "\\2", "g", $0);}}')

cd $HOME

# Upload change to launchpad.net
dput ppa:mvrpl/chsht chsht-data_2.0_source.changes

# Running in Bash Shell
# chmod 755 build_tcs.sh
# . ./build_tcs.sh
