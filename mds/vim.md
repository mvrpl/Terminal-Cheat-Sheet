#### EXTRAS
|||
|-|-|
`:%!python -c "import sys;entrada=sys.stdin.read().strip().split('\n');print '\n'.join([linha for linha in entrada])"`|Change lines with python.
`:%s/^/\=printf('%-2d', line('.'))`|Add the line number at the beginning of the line.
#### SEARCH
|||
|-|-|
`/\<\d\d\d\d\>`|Look exactly 4 digits.
`/fred\|joe`|Search 'fred' or 'joe'.
`/jo[ha]n`|Search 'john' or 'joan'.
`*`|Word hightlight the cursor.
`?word`|Search 'word' from the bottom up.
`/word`|Search 'word' from top to bottom.
`/^\n\{3}`|Find three empty lines.
#### ABAS
|||
|-|-|
`:tabnew`|Creates new tab.
`:tabdo %s/foo/bar/g`|Run a command on all tabs.
`:tablast`|Go to last tab.
`:tabfirst`|Goes to first tab.
`gt`|Next tab.
#### MARKINGS
|||
|-|-|
`'k`|Move cursor to mark 'k'.
`mk`|Mark current position as 'k'.
`:marks`|Shows all markings.
#### INDENT
|||
|-|-|
`[VISUAL SELECTION] shift+<`|Indent back.
`[VISUAL SELECTION] shift+>`|Indents.
`:%retab!4`|Changes tabulation for tab size 4.
#### DIFFERENCES
|||
|-|-|
`vim -d arquivo1.txt arquivo2.txt`|Shows differences between files.
`vim -d <(ls /home) <(ls /tmp)`|Differences of two folders.
`]c`|Next difference.
`[c`|Previous difference.
`do`|Get difference.
`dp`|Puts difference.
`:diffupdate`|Scan files again looking for differences.
#### CCP
|||
|-|-|
`yiw`|Copy word at the cursor.
`viw`|Visual copy word at the cursor.
#### EDITION
|||
|-|-|
`vim +X texto.txt`|Password file.
`:%d`|Clean content.
`:g!/Dave\|John/d`|Delete line does not contain or dave john
`:%!column -t`|Turns text into columns.
`r`|Replaces single character.
`J`|Joint line below the current one.
`cc`|Change (replace) the entire line.
`cw`|Change (replace) to the end of the word.
#### BASIC
|||
|-|-|
`:qa!`|Force output of vim with all open tabs.
`:e filename`|Open file for editing.
`:w!`|Write file and exit vim.
`:q`|Out of vim.
`:w`|Saved file.
`:wqa`|Save and exit flaps
`:qa`|Out of vim with all open tabs.
#### SUBSTITUTIONS
|||
|-|-|
`:s/Bill/Steve/g`|Replace bill by steve on the current line.
`:%s/Bill/Steve/g`|Replace bill by steve throughout the file.
`:g/string/d`|Delete all lines containing 'string'.
`:v/string/d`|Delete all lines that do not contain 'string'.
`:%s/\([a-z]\)/\U\1\e/g`|Changes to lowercase capital letters.
`:%s/\v(\S+)/\=expand(submatch(1))/g`|Substitutes ~ to / home / user.
`:%s!\~!\= expand($HOME)!g`|Substitutes ~ to / home / user.
`:%s/\(STRING\)/\=system("awk '{printf tolower($0)}' <<<".submatch(1))/g`|Replace STRING to string, using AWK tolower() function.
`:s/Bill/Steve/`|Replace the first occurrence of bill by steve on the current line.
`:2,35s/old/new/g`|Replace all occurrences between the lines 2:35
`:%s/\<./\u&/g`|Capitalize word.
`:%s/DATE/\=system("echo -n $(date +\"%Y%m%d\")")/g`|Replace DATE match with YYYYMMDD from system.
`:%s/STRING/\=system('echo -n "'.submatch(0).'+1"')`|Replace STRING to STRING+1 with system.
`:%s/\r//g`|Delete \ r files crlf (^ m).
`:%s#<[^>]\+>##g`|Delete html tags, but keep the text.
`:%s/ *$//g`|Delete all blank spaces
`:%s/old/new/gw`|Replace all occurrences with confirmation.
`:s/Bill/\02/g`|Bill seeks and uses the search value to replace and add 2 on the front.
`:%s/^\(.*\)\n\1$/\1/`|Delete rows that appear twice.
`:5,$s/old/new/g`|Replace all occurrences of 5 to eof line.
`:%s/\([A-Z]\)/\L\1\e/g`|Turns capital letters to lowercase.
`:%s/\d\+/\=(submatch(0)+1)/g`|Sum found numbers.
#### TEXT BOX
|||
|-|-|
`VU`|Change entire line to maiscula.
`gu`|Change selection to lowercase.
`gU`|Change selection to uppercase.
`:%s/\<./\u&/g`|Defines first letter of each word in uppercase.
`:%s/.*/\u&`|Sets first letter of each line to uppercase letters.
`:%s/\<./\l&/g`|Defines first letter of each word in miniscula.
`:%s/.*/\l&`|Sets first letter of each line to lowercase letters.
`g~`|Switch box selection.
`Vu`|Change entire line to tiny.
#### READ / WRITE FILES
|||
|-|-|
`:1,10 w >> texto.txt`|Adds lines 1 to 10 to text.txt.
`:23r texto.txt`|Insert the contents of text.txt file on line 23.
`:r texto.txt`|Insert content of text.txt file in the current.
`:1,10 w texto.txt`|Saved lines 1 to 10 in text.txt.
