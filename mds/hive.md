#### GENERAL
|||
|-|-|
`hive> ! /bin/echo "Hello World";`|Run unix command in the hive.
`hive> ! pwd;`|Shows current directory.
`hive> source /home/queries.hql;`|Execute hql script.
`SHOW COLUMNS FROM tabelax;`|Show table columns.
`LOAD DATA LOCAL INPATH '/home/datafile.txt' INTO TABLE <table name>;`|Load unix file to hive tables.
`LOAD DATA INPATH '/tmp/datafile.txt' INTO TABLE <table name>;`|Load hdfs file to hive tables.
#### DATA TYPE
|||
|-|-|
`TIMESTAMP`|[date and time] UTC, format: yyyy-mm-dd hh: mm: ss.fffffffff.
`DATE`|[date and time] Format: yyyy-mm-dd.
`FLOAT`|[numeric] 4-byte floating-point number of single precision.
`DOUBLE`|[numeric] 8-byte number of floating point double precision.
`BOOLEAN`|[types miscellaneous].
`ARRAY`|[complex types] e.g. array<string>.
`UNIONTYPE`|[complex types] e.g. uniontype<int, double, array<string>.
`TINYINT`|[numeric] Integer 1-byte with sign, -128 up to 127.
`CHAR`|[characters] Text with fixed length. The maximum length and set at 255.
`MAP`|[complex types] e.g. map<string,int>.
`INT`|[numeric] Entire 4-byte with sign, until -2147483648 2147483647.
`BIGINT`|[numeric] Entire 8-byte with sign. of -9,223,372,036,854,775,808 until 9,223,372,036,854,775,807.
`DECIMAL`|[numeric] Accuracy of 38 digits. user define precision and scale.
`STRING`|[characters] Text.
`VARCHAR`|[characters] Text with length specified between 1 and 65355.
`BINARY`|[types miscellaneous].
`STRUCT`|[complex types] e.g. struct<field1:string,field2:string,field3:bigint>.
`SMALLINT`|[numeric] Integer 2-byte with sign, until -32,768 32,767.
