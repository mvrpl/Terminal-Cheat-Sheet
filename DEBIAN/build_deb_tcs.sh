#!/bin/sh

apt-get install build-essential dh-make devscripts

# Export DEB variables
export DEBEMAIL="mvrpl@icloud.com"
export DEBFULLNAME="Marcos Lima"

SOURCEBINPATH=/home/mvro_dev/Terminal-Cheat-Sheet/compiled/Ubuntu64 # Set path of bin
SOURCEBIN=chsht # Bin name
DEBFOLDER=~/chsht # Package directory in user HOME
DEBVERSION=0.1.5 # Set version of package
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

sed -i -e "s/chsht-data/chsht/g" -e "s/\(Section:\s\).*/\1misc/g" -e "s/\(Homepage:\s\).*/\1http:\/\/www.mvrpl.com.br\//g" -e "s/\(Depends:\s.*\)/\1, less/g" debian/control

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
