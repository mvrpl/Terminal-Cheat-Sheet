# Terminal Cheat Sheets

Cheat Sheets on terminal for Unix Like. Written in GO with the following packages below:

  - [GO Sqlite3](https://github.com/mattn/go-sqlite3)
  - [ASCII Table Writer](https://github.com/olekukonko/tablewriter)

## Use
##### Help
```sh
chsht --help
```
##### Show cheat sheet for software
```sh
chsht tmux | less -r
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

1.0

## Download compiled

##### Choose your version: [Centos65|MacElCapitan|Ubuntu64]
```sh
wget https://github.com/mvrpl/Terminal-Cheat-Sheet/blob/master/compiled/[Centos65|MacElCapitan|Ubuntu64]/chsht?raw=true -O chsht
chmod 755 chsht
./chsht tmux
```
Optional: Add **chsht** to your PATH

## License

FREE
