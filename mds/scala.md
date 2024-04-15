#### SPARK
|||
|-|-|
`val confAcc = sc.collectionAccumulator[String]("Conf");sc.parallelize(Array("1", "2", "3")).foreach(c => confAcc.add(c));println(confAcc.value)`|Get distributed data from spark executors to driver.
`val conf = sc.broadcast(Map(1 -> "ok", 2 -> "nok"));sc.parallelize(1 to 10).foreach(x => println(conf.value.getOrElse(x, "NoConf")))`|Input distributed data into spark executors from driver.
#### PACKAGES
|||
|-|-|
`import scala.collection._`|Character wildcard to import everything from a package.
#### STRINGS
|||
|-|-|
`"scala LANG".split("\\s+")`|Split a string based on a regular expression.
`"hello"(0)`|Accessing a Character in a String.
`"test".filter(_ != 't')`|Filter remove characters in string.
`"test".getBytes.foreach(println)`|String to sequence of bytes.
`"scala".drop(2)`|Remove two first letters in string.
`"SCALA".toLowerCase`|Lowercase string.
`"scala".toUpperCase`|Uppercase string.
`"scala language".split(" ")`|Make array of string by split.
`for (c <- "hello") println(c)`|String as a sequence of characters in a for loop.
`"scala".capitalize`|Upper first letter in string.
`"SCALA".replaceAll("A", "X")`|Replace all occurrences of character A in string.
`"scala".take(2)`|Get two first letters in string.
`"eggs, milk, butter, Coco Puffs".split(",").map(_.trim)`|Make array of string and remove spaces.
#### FORMAT SPECIFIERS
|||
|-|-|
`%x`|Hexadecimal number (base 16)
`%%`|Print a "percent" character
`%c`|Character
`%d`|Decimal number (integer, base 10)
`%f`|Floating-point number
`%o`|Octal number (base 8)
`%u`|Unsigned decimal (integer) number
`\%`|Print a "percent" character
`%e`|Exponential floating-point number
`%i`|Integer (base 10)
`%s`|A string of characters
#### DATA TYPE
|||
|-|-|
`val a = 1d`|Double short datatype.
`val a = 1f`|Float short datatype.
`val a = 1000L`|Long short datatype.
`val a = 0: Byte`|Byte datatype.
`val a = 0: Int`|Int datatype.
#### DATA STRUCTURES
|||
|-|-|
`var (x,y,z) = (1,2,3)`|Attribution dysfunctional: unwrapping a tuple through the pattern matching.
`var xs = List(1,2,3)`|List (immutable).
#### FORMAT
|||
|-|-|
`val favLanguage = "scala";println(s"my favorite language is $favLanguage")`|String interpolation with each variable name preceded by a $ character.
`val year = 2017;println(s"Next year is: ${year + 1}")`|Add +1 to the variable(int) inside the string.
`val value = 200.5;println(f"I have $value%.2f pounds.")`|Print variable(int) with two decimal places.
`case class Languages(name: String);val favorite = Languages("Scala");println(s"I love ${favorite.name}")`|Curly braces when printing object fields.
`val name = "Scala";val year = 2001;println("%s born in %d".format(name, year))`|Format string with variables.
#### REGEX
|||
|-|-|
`"[0-9]+".r.findFirstIn("Text").getOrElse("No match")`|Find number in text and print, if not not find, print "No match".
`"123 Main Street".replaceAll("[0-9]", "x")`|Replace all numbers occurrences.
`"[0-9]".r.replaceAllIn("123 Main Street", "x")`|Replace all numbers occurrences.
`"123".replaceFirst("[0-9]", "x")`|Replace first number match.
`"[0-9]".r.replaceFirstIn("123", "x")`|Replace first number match.
`val pattern = "([0-9]+) ([A-Za-z]+)".r ; val pattern(count, fruit) = "100 Bananas"`|Groups with separated variables.
`"[0-9]+".r.findAllIn("350 5th Avenue, New York, New York, 10118.").toArray`|Find all numbers in text and make array.
`"[0-9]+".r.findFirstIn("350 5th Avenue, New York, New York, 10118.")`|Find first number in text.
#### UTILITIES
|||
|-|-|
`for (i <- 1 to 5) yield r.nextInt(100)`|Make vector with 5 random numbers.
`val r = 1 to 10 by 2 toArray`|Make array 1 to 10 stepping by 2
`val x = 1 to 10 toList`|Make list 1 to 10.
`import java.util.{Currency, Locale};val br = Currency.getInstance(new Locale("br", "BR"))`|Set currency to specific country.
`val names = Map("fname" -> "Robert", "lname" -> "Goren");for ((k,v) <- names) println(s"key: $k, value: $v")`|Make map and iteract.
`implicit class StringImprovements(val s: String) {def asBoolean = s match { case "0" | "zero" | "" | " " => false;case _ => true}};"1".asBoolean`|Make StringImprovements asBoolean.
`val r = scala.util.Random;r.nextFloat`|Random float.
`val x = (1 to 10).toList`|Make list 1 to 10.
`List("smallsize", "bigggggsize", "s").reduceLeft((x,y) => if (x.length > y.length) x else y)`|Obtain bigger string by length.
`val arr = Array(("one", 1),("two", 2),("three", 3),("four", 4));var mapa = scala.collection.mutable.Map.empty[String, Int];arr.foldLeft(mapa){(acc, act) => acc += act}`|Fold left array.
`implicit class StringImprovements(val s: String) {def hideAll: String = s.replaceAll(".", "*")};"PASS".hideAll`|Make StringImprovements with implict class.
`val r = scala.util.Random;r.nextInt(100)`|Random number 0 to 100.
