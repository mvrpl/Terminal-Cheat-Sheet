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

0.1.3

## Screenshot

![](https://s11.postimg.org/h5j75navn/Captura_de_tela_2016_10_09_14_06_28.png)

## Install from apt-get
```sh
sudo add-apt-repository ppa:mvrpl/chsht -y && sudo apt-get update
sudo apt-get install chsht
```

## Install formula for homebrew
```sh
brew install chsht
```

## Install from RPM 
####For CentOS 6+
```sh
sudo yum update
sudo yum install https://rawgit.com/mvrpl/Terminal-Cheat-Sheet/master/RPMBUILD/RPMS/x86_64/chsht-0.1.3-1.x86_64.rpm
```

## License

FREE
