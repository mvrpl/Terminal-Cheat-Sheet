#### GENERAL
|||
|-|-|
`disable 'TABELA'`|Inactive table.
`drop 'TABELA'`|Remove table first disable it.
`truncate 'TABELA'`|Clean table records.
`get 'TABELA', 'ROW_KEY'`|Get record.
`create_namespace 'meu_ns'`|Create namespace.
`delete 'tabelax', 'row_key', 'familiacoluna:coluna'`|Delete record.
`list`|Tables list.
`scan 'TABELA'`|Shows all table records.
`create 'meu_ns:tabelax', 'familia'`|Create table namespace.
`drop_namespace 'meu_ns'`|Erases namespace.
`put 'tabelax', 'row_key', 'familiacoluna:coluna', 'valor'`|Insert / update record.
#### SCAN FILTERS
|||
|-|-|
`scan 'TABELA', {FILTER => "ValueFilter(=,'regexstring:[0-9]+')"}`|Filter rows if contains numbers in value.
`scan 'TABELA', {FILTER => "QualifierFilter(!=, 'binary:columnX')"}`|Filter column if not name columnX.
`scan 'TABELA', {FILTER => "ColumnCountGetFilter(1)"}`|Get one column per rowkey.
`scan 'TABELA', {FILTER => "(PrefixFilter('55') AND ValueFilter(=,'regexstring:[0-9]+')"}`|Two Filters with AND or OR.
`scan 'TABELA', { TIMERANGE => [1522870156000,1522870454000] }`|Get lines between timestamps.
`scan 'TABELA', {FILTER =>"(PrefixFilter('55'))"}`|Scan rowkeys starts with.
`scan 'TABELA', {FILTER => "ColumnPrefixFilter('columnx')"}`|Filter starts with column name.
`scan 'TABELA', {FILTER => "ValueFilter(!=,'binaryprefix:STRING')"}`|Filter rows if value not starts with.
