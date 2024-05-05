package main

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"unicode"
	"unicode/utf8"

	_ "github.com/mattn/go-sqlite3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

const (
	version = "0.3.0"
)

type Command struct {
	Command string `bson:"command,omitempty"`
	About   string `bson:"about,omitempty"`
	Session string `bson:"session,omitempty"`
}

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
	os_name := runtime.GOOS
	switch os_name {
	case "windows":
		homedir, err := os.UserHomeDir()
		check(err)
		bin_less := fmt.Sprintf("%s\\scoop\\shims\\less.EXE", homedir)
		cmd = exec.Command(bin_less, "-S", "-R")
	case "darwin", "linux":
		os.Setenv("LESSCHARSET", "utf-8")
		cmd = exec.Command("/usr/bin/less", "-S", "-R")
	}
	in, _ = cmd.StdinPipe()
	cmd.Stdout = os.Stdout
	go func(w io.WriteCloser, data []byte) {
		w.Write(data)
		w.Close()
	}(in, []byte(output))
	err := cmd.Run()
	check(err)
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

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func updateDB() {
	home, _ := os.UserHomeDir()

	db, err := sql.Open("sqlite3", home+string(os.PathSeparator)+".cheat_sheets.db")
	check(err)
	defer db.Close()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	mongoURI := getEnv("MONGODB_CS_URI", "mongodb+srv://public:yDPLVycT6GymvAye@terminalcheatsheet.fbozh8z.mongodb.net/?retryWrites=true&w=majority&appName=TerminalCheatSheet")
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	check(err)
	defer func() {
		err = client.Disconnect(context.TODO())
		check(err)
	}()
	mdb := client.Database("TerminalCheatSheet")
	result, err := mdb.ListCollectionNames(context.TODO(), bson.D{})
	check(err)

	var sql string

	for _, coll := range result {
		sql = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			command TEXT PRIMARY KEY,
			about TEXT NOT NULL,
			session TEXT NOT NULL
		 );`, coll)

		_, errsql := db.Exec(sql)
		check(errsql)

		coll := mdb.Collection(coll)

		cursor, err := coll.Find(context.TODO(), bson.D{})
		check(err)

		var results []Command
		if err = cursor.All(context.TODO(), &results); err != nil {
			panic(err)
		}

		for _, result := range results {
			stmt, err := db.Prepare(fmt.Sprintf("INSERT OR REPLACE INTO %s (command, about, session) values (?, ?, ?);", coll.Name()))
			check(err)
			_, err = stmt.Exec(result.Command, result.About, result.Session)
			check(err)
		}
	}
}

func help() {
	fmt.Println("version: " + version)
	use := []string{"chsht vim", "Example: Show cheat sheet for vim."}
	fmt.Println(strings.Join(use, "\n\t"))
	save_md := []string{"chsht vim --save", "Flag --save to create markdown file (<software_name>.md) in current directory."}
	fmt.Println(strings.Join(save_md, "\n\t"))
	opt_update := []string{"chsht --update", "Update database, internet connection is required."}
	fmt.Println(strings.Join(opt_update, "\n\t"))
	opt_ls := []string{"chsht --list", "List sofwares with cheat sheets in database."}
	fmt.Println(strings.Join(opt_ls, "\n\t"))
}

func main() {
	home, _ := os.UserHomeDir()
	if _, err := os.Stat(home + string(os.PathSeparator) + ".cheat_sheets.db"); os.IsNotExist(err) {
		updateDB()
	}
	db, err := sql.Open("sqlite3", home+string(os.PathSeparator)+".cheat_sheets.db")
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

	if len(os.Args) == 3 && os.Args[2] == "--save" {
		var strDataMd string
		for key, value := range data {
			strDataMd += "#### " + strings.ToUpper(key) + "\n"
			strDataMd += "|||\n|-|-|\n"
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
			for k, v := range cmdsANDdescs {
				strDataMd += fmt.Sprintf("`%s`|%s\n", k, v)
			}
		}

		fileName := fmt.Sprintf("%s.md", os.Args[1])
		if err := os.WriteFile(fileName, []byte(strDataMd), 0755); err != nil {
			fmt.Println("Send <program> --save")
			panic(err)
		}
	} else {
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
}
