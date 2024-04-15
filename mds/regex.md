#### STRING REPLACEMENT
|||
|-|-|
`\+`|Last matched string
`\&`|Entire matched string
`\n`|nth non-passive group
`\2`|"­xyz­" in /^(abc(xyz))$/
`\1`|"­xyz­" in /^(?:abc)(xyz)$/
`\``|Before matched string
`\'`|After matched string
#### POSIX
|||
|-|-|
`[:lower:]`|Lower case letters
`[:alpha:]`|All letters
`[:alnum:]`|Digits and letters
`[:digit:]`|Digits
`[:xdigit:]`|Hexade­cimal digits
`[:punct:]`|Punctu­ation
`[:upper:]`|Upper case letters
`[:blank:]`|Space and tab
`[:space:]`|Blank characters
`[:cntrl:]`|Control characters
`[:graph:]`|Printed characters
`[:print:]`|Printed characters and spaces
`[:word:]`|Digits, letters and underscore
#### ASSERTIONS
|||
|-|-|
`?=`|Lookahead assertion
`?!`|Negative lookahead
`?<=`|Lookbehind assertion
`?!= or ?<!`|Negative lookbehind
`?>`|Once-only Subexp­ression
`?()`|Condition [if then]
`?()|`|Condition [if then else]
`?#`|Comment
#### ANCHORS
|||
|-|-|
`^`|Start of string, or start of line in multi-line pattern.
`\A`|Start of string.
`$`|End of string, or end of line in multi-line pattern.
`\Z`|End of string.
`\b`|Word boundary.
`\B`|Not word boundary.
`\<`|Start of word.
`\>`|End of word.
#### QUANTIFIERS
|||
|-|-|
`?`|0 or 1 {3,5} 3, 4 or 5
`*`|0 or more {3} Exactly 3
`+`|1 or more {3,} 3 or more
#### ESCAPE SEQUENCES
|||
|-|-|
`\`|Escape following character
`\Q`|Begin literal sequence
`\E`|End literal sequence
#### SPECIAL CHARACTERS
|||
|-|-|
`\xxx`|Octal character xxx
`\xhh`|Hex character hh
`\n`|New line
`\r`|Carriage return
`\t`|Tab
`\v`|Vertical tab
`\f`|Form feed
#### GROUPS AND RANGES
|||
|-|-|
`(a|b)`|a or b
`(...)`|Group
`(?:...)`|Passive (non-capturing) group
`[^abc]`|Not (a or b or c)
`[a-q]`|Lower case letter from a to q
`\x`|Group/­sub­pattern number "­x"
`.`|Any character except new line (\n)
`[abc]`|Range (a or b or c)
`[A-Q]`|Upper case letter from A to Q
`[0-7]`|Digit from 0 to 7
#### PATTERN MODIFIERS
|||
|-|-|
`U *`|Ungreedy pattern
`g`|Global match
`i *`|Case-insensitive
`m *`|Multiple lines
`s *`|Treat string as single line
`x *`|Allow comments and whitespace in pattern
`e *`|Evaluate replac­ement
#### CHARACTER CLASSES
|||
|-|-|
`\d`|Digit
`\w`|Word
`\O`|Octal digit
`\s`|White space
`\S`|Not white space
`\D`|Not digit
`\W`|Not word
`\x`|Hexade­cimal digit
`\c`|Control character
