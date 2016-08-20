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

	db, err := sql.Open("sqlite3", os.Getenv("HOME")+"/.cheat_sheets.db")
	check(err)
	dbtables, err := db.Query("select name from sqlite_master where type = 'table'")
	check(err)

	tables := []string{}

	for dbtables.Next() {
		var name string
		_ = dbtables.Scan(&name)
		tables = append(tables, name)
	}

	if len(os.Args) <= 1 {
		fmt.Println("Uso: $ chsht python | less -r, para atualizar: $ chsht --update")
		os.Exit(0)
	}

	if _, err := os.Stat(os.Getenv("HOME") + "/.cheat_sheets.db"); os.IsNotExist(err) {
		updateDB()
	}

	if os.Args[1] == "--update" {
		updateDB()
		fmt.Println("Database updated!")
		os.Exit(0)
	}

	if os.Args[1] == "--version" {
		fmt.Println("chsht 1.0")
		os.Exit(0)
	}

	if os.Args[1] == "--help" {

		opt_update := []string{"--update", "Update database, internet connection is required."}
		fmt.Println(strings.Join(opt_update, "\n\t"))
		opt_ls := []string{"--list", "List sofwares with cheat sheets in database."}
		fmt.Println(strings.Join(opt_ls, "\n\t"))

		os.Exit(0)
	}

	if os.Args[1] == "--list" {
		fmt.Println("Available:")
		for _, soft := range tables {
			fmt.Println("* " + soft)
		}
		os.Exit(0)
	}

	if strings.HasPrefix(os.Args[1], "--") == true {
		fmt.Println("Invalid option. See: --help")
		os.Exit(0)
	}

	Program := os.Args[1]

	if containsStr(tables, Program) != true {
		fmt.Println("Sofware not in list.")
		os.Exit(0)
	}

	rows, err := db.Query("SELECT * FROM " + Program)
	check(err)

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
