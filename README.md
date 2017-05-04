[![Build Status](https://travis-ci.org/mvrpl/Terminal-Cheat-Sheet.svg?branch=master)](https://travis-ci.org/mvrpl/Terminal-Cheat-Sheet)
[![Go Report Card](https://goreportcard.com/badge/github.com/mvrpl/Terminal-Cheat-Sheet)](https://goreportcard.com/report/github.com/mvrpl/Terminal-Cheat-Sheet)
[![Code Climate](https://codeclimate.com/github/mvrpl/Terminal-Cheat-Sheet/badges/gpa.svg)](https://codeclimate.com/github/mvrpl/Terminal-Cheat-Sheet)
# Terminal Cheat Sheets

Cheat Sheets on terminal for Unix Like. Written in GO with the following packages below:

  - [GO Sqlite3](https://github.com/mattn/go-sqlite3)

## Dependencies
  - [GNU Less](https://www.gnu.org/software/less/)

## Use
##### Help
```sh
chsht --help
```
##### Show cheat sheet for software
```sh
chsht tmux
```
##### Update database
```sh
chsht --update
```
##### List softwares
```sh
chsht --list
```
##### Version
```sh
chsht --version
```

## Version

0.1.5

## Screenshot

![](https://s11.postimg.org/h5j75navn/Captura_de_tela_2016_10_09_14_06_28.png)

## Install formula for homebrew
```sh
brew install mvrpl/chsht/chsht
```

## Install from RPM 
#### For CentOS 6+
```sh
sudo yum update
sudo yum install https://cdn.rawgit.com/mvrpl/Terminal-Cheat-Sheet/515e6f99/RPMBUILD/RPMS/x86_64/chsht-0.1.5-1.x86_64.rpm
```
## See database
[SQLite Viewer Online](https://sqliteonline.com/#fiddle-58f75932c55f5b9gj1oym9lq)

## License

MIT
