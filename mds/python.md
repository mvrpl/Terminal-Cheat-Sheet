#### METHODS STRINGS
|||
|-|-|
`'STRING'.capitalize()`|Changes string to first upper rest to lower.
`'STRING'.center(10)`|Centralizes string space of 10 characters
`'   '.isspace()`|Checks if the string contains only spaces
`'STRING/STRING2'.split('/')`|Separates in list based on character
`'STRING'.isupper()`|Validates the string this all in upper.
#### DATA TYPES
|||
|-|-|
`bool`|Boolean.
`dict`|Dictionary.
`list`|Sequence changeable.
`tuple`|Sequence immutable.
`float`|Floating-point number.
`long`|Full greater than 32-bits.
`int`|Full 32-bit.
`str`|Sequence of characters.
`complex`|Number complex.
#### FORMAT DATES
|||
|-|-|
`%m`|Months (01 to 12).
`%w`|The weekday number (0-6).
`%Y`|Year four digits.
`%b`|Abbreviated month (jan).
`%c`|Date and time.
`%H`|24 hours (leading zeros) (00-23).
`%j`|Day of the year (001-366).
`%M`|Minute (00-59).
`%f`|Milliseconds.
`%A`|Day (monday).
`%d`|Day (leading zeros) (01-31).
`%x`|Date.
`%X`|Hour.
`%y`|Two year digits.
`%a`|Day abbreviation (mon).
`%B`|Name of the month (january).
`%I`|12 noon (leading zeros) (01 to 12).
`%S`|Seconds (00-59).
`%Z`|Timezone.
`%p`|Am or pm.
`%U`|Week number (00-53).
`%%`|Literally%.
#### UTILS
|||
|-|-|
`print ''.join(['\\x'+char.encode('hex') for char in 'Hello World\x0a'])`|Convert unicode characters in HEX.
