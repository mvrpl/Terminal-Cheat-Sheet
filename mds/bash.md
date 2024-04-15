#### EXPANSION
|||
|-|-|
`test="MARCOS";echo ${test1:-text}`|If varname exists and isn't null, return its value; otherwise return 'text'.
`test="TEST";echo ${test1:=text};echo $test1`|If varname exists and isn't null, return its value; otherwise set it 'test' and then return its value.
`test="marcos";echo ${test^^}`|String variable into uppercase (Required: Bash 4.x).
`test="MARCOS";echo ${test,,}`|String variable into lowercase (Required: Bash 4.x).
`test="My name is MARCOS";echo ${test/MARCOS/JONN}`|Find and replace string in variable. Only the first match is replaced.
`VERSION='2.3.3';echo "${VERSION//./,}"`|Replace all character matched.
`test="MARCOS";echo ${test%COS}`|Return (MAR) the rest before pattern.
`test="MARCOS vs MARCOS";echo ${test//MARCOS/JONN}`|Find and replace string in variable. All matches are replaced.
`echo ${test1:+test}`|If varname exists and isn't null, return 'test'; otherwise return null.
`test="MARCOS";echo ${test#MAR}`|Return (COS) the rest after pattern.
`echo ${test2:?error};echo $?`|If varname exists and isn't null, return its value; otherwise print varname, followed by message and abort the current command or script.
`test="LINUX";echo ${test:1:3}`|Slice variable with offset and length (${varname:offset:length}).
`test="marcos";echo ${#test}`|Length of the value of variable string.
`VERSION='2.3.3';echo "${VERSION//.}"`|Remove all character matched.
#### RETURN CODE
|||
|-|-|
`ech 1 | wc -l;echo -e "Return codes: \n  command1" ${PIPESTATUS[0]} "\n  command2" ${PIPESTATUS[1]}`|Pipes commands with list of returns codes.
#### ARRAYS
|||
|-|-|
`declare -A map;map[key]="val";map[key2]="val2";for k in "${!map[@]}";do echo "key:" $k;echo "value:" ${map[$k]};done`|Associative arrays in bash.
`arr=(one two three);if [[ " ${arr[*]} " == *" one "* ]]; then echo 'exists!';fi`|Check item exists in array.
#### DATE
|||
|-|-|
`data="20180701";dataFim="20180820";until [[ $data > $dataFim ]]; do echo "$data" ; data=$(date +'%Y%m%d' -d "$data + 1 day");done`|Print all days between two dates.
`echo {1..10} | xargs -I{} -d ' ' date --date={}' days ago' +"%Y-%m-%d"`|Get dates for 10 days ago.
