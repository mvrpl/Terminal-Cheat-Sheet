package main

import (
	"database/sql"
	"flag"
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

func main() {

	if os.Args[1] == "-update" {
		db, _ := os.Create(os.Getenv("HOME") + "/.cheat_sheets.db")
		defer db.Close()
		resp, _ := http.Get("https://github.com/mvrpl/Terminal-Cheat-Sheet/blob/master/cheat_sheets.db?raw=true")
		defer resp.Body.Close()
		io.Copy(db, resp.Body)
		fmt.Println("Banco de dados atualizado com sucesso!")
		os.Exit(0)
	}

	Program := flag.String("program", "", "Cheat Sheet for the program. (Required)")
	flag.Parse()

	db, _ := sql.Open("sqlite3", os.Getenv("HOME")+"/.cheat_sheets.db")
	rows, err := db.Query("SELECT * FROM " + *Program)
	dbtables, err := db.Query("select name from sqlite_master where type = 'table'")

	tables := []string{}
	for dbtables.Next() {
		var name string
		_ = dbtables.Scan(&name)
		tables = append(tables, name)
	}

	if containsStr(tables, *Program) != true {
		fmt.Println("Program not exists in database!")
		os.Exit(1)
	}

	if err != nil {
		panic(err)
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
		table.SetHeader([]string{"Command", "Description"})
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
