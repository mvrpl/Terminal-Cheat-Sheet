#### HTTP
|||
|-|-|
`http.Get("http://mvrpl.com.br")`|Get code source url.
#### SYSTEM
|||
|-|-|
`os.Open("arquivo.txt")`|Open file pointer.
`os.Getenv("HOME")`|Get environment variable value.
`os.Create("arquivo.txt")`|Create new file.
`if _, err := os.Stat("cheat_sheets.db"); os.IsNotExist(err) {fmt.Println("arquivo nao existe")}`|Check if file exists.
