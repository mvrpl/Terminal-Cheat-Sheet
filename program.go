package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/olekukonko/tablewriter"
	"io"
	"net/http"
	"os"
	"strings"
)

func containsStr(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}
	_, ok := set[item]
	return ok
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func updateDB() {
	db, err := os.Create(os.Getenv("HOME") + "/.cheat_sheets.db")
	check(err)
	defer db.Close()
	resp, err := http.Get("https://github.com/mvrpl/Terminal-Cheat-Sheet/blob/master/cheat_sheets.db?raw=true")
	check(err)
	defer resp.Body.Close()
	io.Copy(db, resp.Body)
}

func main() {

	if len(os.Args) <= 1 {
		fmt.Println("Uso: $ chsht python | less -r, para atualizar: $ chsht -update")
		os.Exit(0)
	}

	if _, err := os.Stat(os.Getenv("HOME") + "/.cheat_sheets.db"); os.IsNotExist(err) {
		updateDB()
		fmt.Println("Banco de dados instalado com sucesso!")
	}

	if os.Args[1] == "-update" {
		updateDB()
		fmt.Println("Banco de dados atualizado com sucesso!")
		os.Exit(0)
	}

	Program := os.Args[1]

	db, _ := sql.Open("sqlite3", os.Getenv("HOME")+"/.cheat_sheets.db")
	rows, _ := db.Query("SELECT * FROM " + Program)
	dbtables, err := db.Query("select name from sqlite_master where type = 'table'")
	check(err)

	tables := []string{}
	for dbtables.Next() {
		var name string
		_ = dbtables.Scan(&name)
		tables = append(tables, name)
	}

	if containsStr(tables, Program) != true {
		fmt.Println("Sofware nao existe na base da dados.")
		fmt.Println("Disponiveis:")
		for _, soft := range tables {
			fmt.Println("  - " + soft)
		}
		os.Exit(0)
	}

	var data map[string][]map[string]string
	data = make(map[string][]map[string]string)

	for rows.Next() {
		var command string
		var about string
		var session string
		_ = rows.Scan(&command, &about, &session)
		commands := map[string]string{
			command: about,
		}
		data[session] = append(data[session], commands)
	}

	for key, value := range data {
		fmt.Println("+" + strings.Repeat("-", len(key)+2) + "+")
		fmt.Println("| \033[1m" + strings.ToUpper(key) + "\033[0m |")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Comando", "Descricao"})
		table.SetRowLine(true)
		table.SetRowSeparator("-")
		for _, v := range value {
			for cmd, desc := range v {
				outline := []string{cmd, desc}
				table.Append(outline)
			}
		}
		table.Render()
	}
}
