package main

import (
	"bytes"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

const (
	version = "0.1.5"
)

func NormalizeString(text string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	s, _, _ := transform.String(t, text)
	return s
}

func rightPad2Len(str, pad string, lenght int) string {
	for {
		str += pad
		if utf8.RuneCountInString(str) > lenght {
			return str
		}
	}
}

func rightPad2LenNorm(s string, padStr string, overallLen int) string {
	s = NormalizeString(s)
	var padCountInt int
	padCountInt = 1 + ((overallLen - utf8.RuneCountInString(padStr)) / utf8.RuneCountInString(padStr))
	var retStr = s + strings.Repeat(padStr, padCountInt)
	return retStr[:overallLen]
}

func OutLess(output string) {
	var in io.WriteCloser
	var cmd *exec.Cmd
	os.Setenv("LESSCHARSET", "utf-8")
	cmd = exec.Command("/usr/bin/less", "-S", "-R")
	in, _ = cmd.StdinPipe()
	cmd.Stdout = os.Stdout
	io.Copy(in, bytes.NewBufferString(output))
	in.Close()
	cmd.Run()
}

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
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get("https://github.com/mvrpl/Terminal-Cheat-Sheet/blob/master/cheat_sheets.db?raw=true")
	if err != nil {
		fmt.Println("Error updating database!")
		os.Exit(1)
	}
	db, err := os.Create(os.Getenv("HOME") + "/.cheat_sheets.db")
	check(err)
	defer db.Close()
	defer resp.Body.Close()
	io.Copy(db, resp.Body)
}

func help() {
	fmt.Println("version: " + version)
	use := []string{"chsht vim", "Example: Show cheat sheet for vim."}
	fmt.Println(strings.Join(use, "\n\t"))
	opt_update := []string{"chsht --update", "Update database, internet connection is required."}
	fmt.Println(strings.Join(opt_update, "\n\t"))
	opt_ls := []string{"chsht --list", "List sofwares with cheat sheets in database."}
	fmt.Println(strings.Join(opt_ls, "\n\t"))
}

func main() {
	if _, err := os.Stat(os.Getenv("HOME") + "/.cheat_sheets.db"); os.IsNotExist(err) {
		updateDB()
	}
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
		help()
		os.Exit(0)
	}

	if os.Args[1] == "--update" {
		updateDB()
		fmt.Println("Database updated!")
		os.Exit(0)
	}

	if os.Args[1] == "--version" {
		fmt.Println("version: " + version)
		os.Exit(0)
	}

	if os.Args[1] == "--help" {
		help()
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

	var strData string
	for key, value := range data {
		strData += "+" + strings.Repeat("-", len(key)+2) + "+\n"
		strData += "| \033[1m" + strings.ToUpper(key) + "\033[0m |\n"
		strData += "+" + strings.Repeat("-", len(key)+2) + "+\n"
		cmdSize := 0
		descSize := 0
		var cmdsANDdescs map[string]string
		cmdsANDdescs = make(map[string]string)
		for _, v := range value {
			for cmd, desc := range v {
				if len(cmd) > cmdSize {
					cmdSize = utf8.RuneCountInString(cmd)
				}
				if len(desc) > descSize {
					descSize = utf8.RuneCountInString(desc)
				}
				cmdsANDdescs[cmd] = desc
			}
		}
		strData += "+" + strings.Repeat("-", cmdSize+2) + "+" + strings.Repeat("-", descSize+2) + "+\n"
		for k, v := range cmdsANDdescs {
			strData += rightPad2Len("| "+k, " ", cmdSize+2) + "| "
			strData += rightPad2LenNorm(v, " ", descSize) + " |\n"
			strData += "+" + strings.Repeat("-", cmdSize+2) + "+" + strings.Repeat("-", descSize+2) + "+\n"
		}
	}
	OutLess(strData)
}
