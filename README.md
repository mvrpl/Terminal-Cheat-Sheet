# Terminal Cheat Sheets

Cheat Sheets on terminal for Unix Like. Written in GO with the following packages below:

  - [GO Sqlite3](https://github.com/mattn/go-sqlite3)

## Demo
##### User: marcos | Pass: 0312
[http://ec2-54-70-189-64.us-west-2.compute.amazonaws.com:4200](http://ec2-54-70-189-64.us-west-2.compute.amazonaws.com:4200)

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

2.0

## Screenshot

![](https://s11.postimg.org/h5j75navn/Captura_de_tela_2016_10_09_14_06_28.png)

## Install from apt-get
```sh
sudo add-apt-repository ppa:mvrpl/chsht && sudo apt-get update
sudo apt-get install chsht
```

## Download compiled

##### Choose your version: [Centos65|MacSierra|Ubuntu64]
```sh
wget https://github.com/mvrpl/Terminal-Cheat-Sheet/blob/master/compiled/[Centos65|MacSierra|Ubuntu64]/chsht?raw=true -O chsht
chmod 755 chsht
./chsht tmux
```
Optional: Add **chsht** to your PATH

## License

FREE
