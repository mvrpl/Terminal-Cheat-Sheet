#### IO
|||
|-|-|
`io:format(" I am ~s~n", ["asda"]).`|Output on the console.
#### SYSTEM
|||
|-|-|
`erlang:system_info(version).`|System version.
`code:lib_dir().`|Get directory libs.
`os:getenv("HOME").`|Obtain environment variable.
#### FILES
|||
|-|-|
`file:read_file("TESTE.txt").`|Read file contents.
`file:write_file("TESTE.txt", "Conteudo", [append]).`|Record content in files.
#### WEB
|||
|-|-|
`inets:start(), {ok, {{_, _, _}, _, Body}} = httpc:request(get, {"http://mvrpl.com.br/", []}, [], []).`|Read content url.
#### STRINGS
|||
|-|-|
`binary_to_list(VAR).`|Bytes for string.
`lists:concat(["CON", "CAT"]).`|Concatenate strings.
#### DATA STRUCTURE
|||
|-|-|
`maps:get(<<"KEY">>, Map).`|Get key value maps.
#### EXCEPTIONS
|||
|-|-|
`try funcX() catch _:_ -> io:fwrite("~p", ["ERROR"]) end.`|Trycatch.
