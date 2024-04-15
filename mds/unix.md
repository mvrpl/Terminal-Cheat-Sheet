#### HISTORIC
|||
|-|-|
`history | awk '{a[$2]++}END{for(i in a){print a[i] " " i}}' | sort -rn | head`|Rank most used by user commands.
`echo "!!" > foo.sh`|Create a script of the last executed command.
`export HISTTIMEFORMAT="%F %T "`|Add timestamp to history.
`echo "discreet";history -d $(history 1)`|Execute command without keeping it in history.
#### SED
|||
|-|-|
`echo -e "FIELD;1;2;3\nFIELD;4;5;6" | sed 's/\(^[^;]\+\);[^;]\+;\(.*\)$/\1;X;\2/g'`|Change second value with delimited stdin.
`echo -e "teste1\nteste2\nteste3\nteste4\nteste5" | sed '/2/{n;n;s/teste/TESTE/}'`|Change content in N lines after match specified line, this example modifying teste to TESTE after 2 lines (because has two n;) after matched string "2".
`sed 's/.$//g' filename`|Convert dos files to unix.
`sed -i '1d' *.txt`|Remove first line in all files .txt.
`sed -i '/^\s*$/d' file.txt`|Remove blank lines or lines with only spaces.
`echo "this is a test" | sed 's/.*/\L&/; s/[a-z]*/\u&/g'`|Convert a string to "Title Case".
`seq 100 | sed 10q`|Print the first 10 lines of a file (emulates "head -10").
`echo -e "um\ndois\ntres\nquatro\ncinco" | sed -n '/\(tres\|quatro\)/p'`|Print only the lines that match a regular expression (emulates "grep").
`echo -e "um\ndois\ntres\nquatro\ncinco" | sed '/tres/,$!d'`|Delete all lines before a keyword with sed.
`seq 100 | sed -e :a -e '$q;N;11,$D;ba'`|Print the last 10 lines of a file (emulates "tail -10").
`sed -n '10,20p' teste.txt`|Show 10 to 20 line of the file.
`sed -i <file> -re '<start>,<end>d'`|Remove a range of lines from a file.
`sed -i '1 i FIRST LINE' <filename>`|Add a line at beginning of file.
`echo -e "um\ndois\ntres\nquatro" | sed -n '/tres/{x;p;};h'`|Get line before match.
`echo -e "um\ndois\ntres\nquatro" | sed -n '/tres/{n;p;};h'`|Get line after match.
`echo -e "um\ndois\ntres\nquatro\ncinco" | sed '/tres/,$d'`|Delete all lines including the keyword and after the keyword with sed.
`echo "1000" | sed ':a;s/\B[0-9]\{3\}\>/\.&/;ta'`|Add thousands separator in a number.
`echo "test1,test2,test3," | sed 's/,$//g'`|Remove last character.
`sed -i '$ a LAST LINE' <filename>`|Add at end of file.
`echo -e "um\ndois\ntres\nquatro\ncinco" | sed -n '/tres/{x;p;n;p;};h'`|Get lines after/before match.
`echo -e "um\ndois\ntres\nquatro\ncinco" | sed -n '/tres/{N;p;}'`|Get matched line and next line.
`echo -e "um\ndois\ntres\nquatro\ncinco" | sed -n '/\(tres\|quatro\)/!p'`|Print only the lines that do not match a regular expression (emulates "grep -v").
`sed -i 'y/áÁãÃâÂàÀéÉêÊíÍóÓõÕôÔúÚüÜçÇ/aAaAaAaAeEeEiIoOoOoOuUuUcC/' file.txt`|Replace accents characteres in file.
`echo -e "um\ndois\ntres\nquatro\ncinco" | sed '/tres/q'`|Delete all lines after a keyword with sed.
`echo 123 | sed -e :a -e 's/^.\{0,5\}$/0&/;ta'`|Zero pad a number.
#### CHARACTERS
|||
|-|-|
`echo -e "\x7c"`|Print hexadecimal character.
#### XARGS
|||
|-|-|
`ls *.txt | xargs -I '{}' -n1 -P4 python script.py {}`|Parallelize script with args passing filenames of return from ls, keep 4 processes until finished.
`find . -name "*.txt" | xargs rm -f`|Remove forced all .txt files in current directory.
`echo 1 2 3 | xargs -l bash -c 'echo this args $0 $1 $2'`|Multiple arguments in xargs.
#### SORT
|||
|-|-|
`echo -e "1,a,3\n2,f,4\n3,b,2" | sort -t ',' -k 1rn`|Invert sort by column 1 (number).
`echo -e "4\n4\n2\n1" | sort -u`|Sort and unique.
`echo -e "1,a,3\n2,f,4\n3,b,2" | sort -t ',' -k 1n`|Sort by column 1 (number).
#### DIRECTORIES
|||
|-|-|
`cd -`|Back to previous directory.
#### ARCHIVE
|||
|-|-|
`tar xvf archive_name.tar`|Extract existing tar the current directory.
`tar cvzf archive_name.tar.gz dirname/`|Compact folder to .tar.gz file
`gunzip -c FILE.txt.gz > FILE.txt`|Unzip file and maintain original.
`tar cf - <dir>|split -b<max_size>M - <name>.tar.`|Split a tarball into multiple parts.
`tar xvfj somefilename.tar.bz2`|Extract your tar.bz2 file.
`tar tvf archive_name.tar`|See tarball.
#### PROCESSES
|||
|-|-|
`pkill -u $USER -f script.sh`|Kill all processes that correspond to "script.sh" on current user.
`sudo cpulimit -p <pid> -l 50`|Limit the cpu usage of a process.
`pwdx <pid>`|Show current working directory of a process.
#### DATETIME
|||
|-|-|
`date -d@1493211932`|Convert seconds to human-readable format.
#### SSH
|||
|-|-|
`ssh -D 1080 -N user@server`|Create socks proxy via SSH.
#### FOLDERS
|||
|-|-|
`find . -type d -empty -delete`|Find and delete empty dirs, start in current working dir.
#### FILES
|||
|-|-|
`find / -name *.jpg -type f -print | xargs tar cvzf images.tar.gz`|Finds all .jpg files and compresses them.
`cmp --silent file1.txt file2.txt`|Check two file identical with ($ echo $?) after run command.
`cat my-long-file.txt | more`|It shows great content file.
`ls *.jpg | xargs -n1 -i cp {} /directory`|List all .jpg files and copy to specific directory.
`rename 'y/A-Z/a-z/' *`|Convert filenames in current directory to lowercase.
`split -b 19m <file> <nameForPart>`|Split File in parts.
`cat -n myfile.txt`|Shows content from file with lines number.
`rename 'y/ /_/' *`|Replace spaces in filenames with underscores.
`find . -type f -name "*.*" | awk -F'.' '{print $NF}' | sort | uniq -c | sort -g`|List all file extensions in a directory.
`touch --date "2010-01-05" <file>`|Change/Modify timestamp of file.
`fallocate -l 3.3G teste_arq.img`|Generate fake big file with bytes size of 3.3GB.
`rm !(*.foo|*.bar|*.baz)`|Delete files that are not the extensions in parentheses.
`find . -maxdepth 1 -size 0c -delete`|Remove all zero size files from current directory (not recursive).
`find /path_here -type l -! -exec test -e {} \; -print | xargs rm`|Alternative to delete broken symlinks.
`find -L /path/to/check -type l -delete`|Find broken symlinks and delete them.
#### SYSTEM INFORMATION
|||
|-|-|
`compgen -v | while read line; do echo $line=${!line};done`|Lists all variables and values including local ones.
`lsb_release -a`|Show S.O. release.
`getconf LONG_BIT`|Shows whether the system and 32 or 64 bits.
#### OTHERS
|||
|-|-|
`nc -v -l 9088 < teste.txt`|Share file in the port 9088 (sudo).
`if [ $(($RANDOM%2)) -eq 0 ];then echo -$(shuf -i 1-1000 -n 1);else shuf -i 1-1000 -n 1;fi`|Random positive/negative numbers.
`echo $[RANDOM%10+1]`|Random number between 1 and 10.
`echo -e "one\ntwo\nthree\nfour" | pr -Tn`|Show line number of output.
`echo -n "Marcos" | od -t x1 -c`|Display stdout in Hexadecimal.
`shuf -i 1-100 -n 1`|Number from 1 to 100 random.
`set -o vi`|Put readline into VI mode.
`lower() { echo ${@,,}; }`|Function to convert text to lowercase, usage: ($ lower STRING).
`echo -e "1\n2\n3\n4" | tac`|Reverse output lines [TAC].
`echo $(( ((RANDOM<<15)|RANDOM) % 5 + 1 ))`|Show random number between 1-5.
`? () { echo "$*" | bc -l; }`|Define a quick calculator function, to use: (bash$ ? 10+10).
#### GREP
|||
|-|-|
`grep -Hrn -i -F "TEXTO" . -I --exclude-dir .git`|Text search for files in the current folder, H (filename) r (recursive) n (line number) -i (ignore case) -F (search without using regex) -i (ignores binary files).
#### AWK
|||
|-|-|
`echo -e "1\n2\n3\n4\n5\n6\n7" | awk 'BEGIN { c=0 } { if( c == 3 ) { print l; l=$0; c=1 } else { if( l != "") { l=l" "$0 } else { l=$0 }; c++}} END { print l }'`|Batch output of the first command.
`echo "/home/f/i/le.txt" | awk '{sub(".*/", "", $0);print}'`|Basename in AWK.
` echo -e "1\n1\n2\n3" | awk '++seen[$0] == 2'`|Print duplicates.
`echo "test1,test2,test3" | awk '{count=split($0,list,",");for (i=1;i<=count;i++){printf (i == count) ? list[i] "\n" : list[i] "-"}}'`|Split and join in AWK.
`awk '{$1=$3=""}1' <file>`|Print all columns except the 1st and 3rd.
`echo -e "zero\num\ndois\ntres\nquatro\ncinco" | awk '/tres/{for(i=1;i<=x;)print a[i++];print}{for(i=1;i<x;i++)a[i]=a[i+1];a[x]=$0;}' x=2`|Get N lines before keyword with keyword.
`echo "123456" | awk 'length > 5'`|Print lines longer than 72 characters.
`echo "one=two three" | awk -F "=| " '{print $1, $3}'`|Using multiple field separators, the | is separator.
`echo "\"string,2\",2" | awk -vFPAT='([^,]*)|("[^"]+")' '{print $2}'`|Ignore the field delimiter inside double quotes.
`echo -e "zero\num\ndois\ntres\nquatro\ncinco" | awk '{a[++i]=$0;}/tres/{for(j=NR-2;j<=NR;j++)print a[j];}'`|Get N lines before keyword with keyword SHORT.
`echo -e "1  \n2  \n3  " | awk 'function rtrim(s) { sub(/[ \\t\\r\\n]+$/, "", s); return s }{print rtrim($0)}'`|Remove last spaces and tabs with function rtrim.
`echo -e "-|PI||PES|-" | awk '{gsub(/^[\x7C|\x2D]+/, "", $0);gsub(/[\x7C|\x2D]+$/, "", $0);print $0}'`|Trim (| and - in HEX) left and right in AWK.
`echo -e "Line1\nLine2\nLine3\nLine4\nLine5" | awk '{if($0 ~ /Line3/){ show[NR+2]++ }} show[NR]'`|Print second line after match.
`echo "/" | awk '{ system("ls "$0"") }'`|Run sys commands.
`echo "123 asd" | awk '{ if (match($0,/[0-9]+/,m)) print m[0] }'`|Get regex matches.
`echo "Your home folder size" | awk '{cmd = "du -sh $HOME | awk \x27{print $1}\x27";cmd | getline out;close(cmd)};{print $0 " = " out }'`|Run system command inside AWK.
`echo -e "12\nstart\n345\nstop\n678" | awk '/start/, /stop/'`|Print all lines between start/stop pairs
`echo "1 2 3 4" | awk '{ for (i = NF; i > 0; --i) print $i }'`|Print fields in reverse order.
`awk '!x[$0]++' <file>`|Remove duplicate entries in a file without sorting.
`echo 1000000000000 | awk '{xin=$1;if(xin==0){print "0 B";}else{x=(xin<0?-xin:xin);s=(xin<0?-1:1);split("B KB MB GB TB PB",type);for(i=5;y < 1;i--){y=x/(2^(10*i));}print y*s " " type[i+2];};}'`|Convert bytes to human.
`echo "hello world" | awk '{ print NF }'`|Count words.
`echo "test1test2test3" | awk '{ width=5; line = $0; do { print substr(line, 0, width); line = substr(line, width+1); } while (length(line) > 0); }'`|Folding output.
`echo -e "zero\num\ndois\ntres\nquatro\ncinco" | awk '{a[++i]=$0;}/tres/{for(j=NR-1;j<=NR;j++)print a[j-1];}'`|Get N lines before keyword without keyword SHORT.
`echo -e "Line 1\nLine 2\nLine 3\nLine 4" | awk '{printf "%-5s", NR;print $0}'`|Print number line with spaces of content.
`awk -F, 'FILENAME == ARGV[1] { one[FNR]=$1","$2","$3 };FILENAME == ARGV[2] { two[FNR]=$0 };END {for (i=1; i<=length(one); i++) {print one[i] "," two[i]}}' file1.csv file2.csv`|Join the first 3 columns of file 1 with all columns of file 2.
`echo -e "1\n2\n3\n4" | awk '{a[i++]=$0} END {for (j=i-1; j>=0;) print a[j--] }`|Reverse output lines.
`echo -e "1\n2\n3\n4\nPATTERN\n6\n7\n8" | awk 'c&&c--; /^PATTERN/{c=2}'`|Skip line matched pattern and print the next two lines.
`echo -e "1\n2\n3\n4\nPATTERN\n6\n7\n8" | awk '/^PATTERN/{c=2} c&&c--'`|Print two lines with the pattern line.
`echo -e "1\n5\n3\n4\n5\n2" | awk '!seen[$0]++'`|Unique without sort.
`echo -e "  1\n  2\n  3" | awk 'function ltrim(s) { sub(/^[ \\t\\r\\n]+/, "", s); return s }{print ltrim($0)}'`|Remove initial spaces and tabs with function ltrim.
`echo -e "  1  \n  2  \n  3  " | awk 'function ltrim(s) { sub(/^[ \\t\\r\\n]+/, "", s); return s };function rtrim(s) { sub(/[ \\t\\r\\n]+$/, "", s); return s };function trim(s)  { return rtrim(ltrim(s)); }{print trim($0)}'`|Remove initial/last spaces and tabs with function trim.
`echo -e "1,2,3\n1,2,3" | awk -F, 'function randomstr(num){cmd = "cat /dev/urandom | tr -cd 'A-Z0-9' | head -c "num"";cmd | getline randoms;close(cmd);return randoms}{print $1 "|" $2 "|" $3 "|" randomstr(10) }'`|Function to generate letters and random numbers with fixed size.
`echo "1 23 4" | awk '{print length($2)}'`|Print length of string in 2nd column.
`echo -e "zero\num\ndois\ntres\nquatro\ncinco" | awk '/tres/{for(i=1;i<=x;)print a[i++]}{for(i=1;i<x;i++)a[i]=a[i+1];a[x]=$0;}' x=2`|Get N lines before keyword without keyword.
#### FIND
|||
|-|-|
`find . -name '*.csv' -o -name '*.txt'`|Find files .csv or .txt.
`find . -group sudo`|Find all files that belong to a particular group.
`find /home/bob -cmin -60`|Find files modified within the last 1 hour.
`find / -type f -size +500M`|Find all file larger than 500M.
`find . -name '*.sh' ! -name '*.py'`|Find .sh and not .py files.
`find ./dir1 ./dir2 -type f -name "abc*"`|Find files in dir1 and dir2.
`find / -atime 50`|Find all files that were accessed in the last 50 days.
`find / -mtime 50`|To find all the files which are modified 50 days back.
`find / -amin -60`|To find all the files which are accessed in last 1 hour.
`find / -mmin -60`|To find all the files which are modified in last 1 hour.
`find / -type f -size +500M -size -1G`|Find all files larger than 500M and less than 1GB.
`find . -user <UNIX USER>`|Find files belonging to particular user.
`find . -type f ! -perm 0777`|Find files with not permission 777.
`find / -mtime +50 –mtime -100`|Find all files that were modified between 50 to 100 days ago.
`find . -type f -newermt "2017-04-01" ! -newermt "2017-04-26"`|Find files in a date range.
`find . -type f -perm 0664`|Find files with the permission 0664.
#### SYSTEM
|||
|-|-|
`color()(set -o pipefail;"$@" 2>&1>&3|sed $'s,.*,\e[31m&\e[m,'>&2)3>&1`|Function to turns red the stderr output, to use ($ color eco 123).
`taskset -c 0 <your_command>`|Start a command on only one CPU core.
`echo $PATH | tr \: \\n`|Show directories in the PATH, one per line.
`cat /etc/*release`|Display which distro is installed.
#### PERL
|||
|-|-|
`perl -e 'use Term::ANSIColor;$|++; while (1) { print color("green")," " x (rand(35) + 1), int(rand(2)), color("reset") }'`|Binary digits Matrix effect.
`echo -e "1\n2\n3\n4" | perl -e 'print reverse <>'`|Reverse output lines.
#### CURL
|||
|-|-|
`curl -sI http://imgs.xkcd.com/comics/i_know_youre_listening.png | grep Last-Modified`|Check the last modified date of a file over HTTP.
