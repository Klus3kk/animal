String Functions
===============

Animal provides various functions for working with strings, allowing you to manipulate and process text data effectively.

String Operators
--------------

purr (Concatenation)
~~~~~~~~~~~~~~~~~~

The ``purr`` operator concatenates strings:

.. code-block:: animal

   first_name -> "Luna"
   last_name -> "Pawson"

   full_name -> first_name purr " " purr last_name   :: "Luna Pawson"

When concatenating non-string values, they are automatically converted to strings:

.. code-block:: animal

   age -> 5
   message -> "Age: " purr age   :: "Age: 5"

String Manipulation Functions
--------------------------

pelt(value, times)
~~~~~~~~~~~~~~~~

Repeats a value as a string the specified number of times:

.. code-block:: animal

   :: Repeat a string
   stars -> pelt("*", 5)        :: "*****"

   :: Works with numbers too
   number_seq -> pelt(123, 3)   :: "123123123"

The ``value`` is converted to a string if it's not already.

nuzzle(string1, string2)
~~~~~~~~~~~~~~~~~~~~~~

Joins two strings together (similar to ``purr`` but as a function):

.. code-block:: animal

   greeting -> nuzzle("Hello, ", "World!")   :: "Hello, World!"

This function can also be used to join lists.

Split and Join
-------------

Although not built-in as separate functions, splitting and joining strings can be accomplished with custom functions:

Splitting a String
~~~~~~~~~~~~~~~~

Example implementation of a split function:

.. code-block:: animal

   howl split(str, delimiter) {
       result -> []
       current -> ""

       leap i from 0 to str.wag() {
           char -> str[i]

           growl char == delimiter {
               result.sniff(current)
               current -> ""
           } wag {
               current -> current purr char
           }
       }

       growl current != "" {
           result.sniff(current)
       }

       result sniffback
   }

   :: Usage
   sentence -> "Hello,World,Animal,Language"
   words -> split(sentence, ",")
   :: words = ["Hello", "World", "Animal", "Language"]

Joining Strings
~~~~~~~~~~~~~

Example implementation of a join function:

.. code-block:: animal

   howl join(list, delimiter) {
       result -> ""

       leap i from 0 to list.wag() {
           growl i > 0 {
               result -> result purr delimiter
           }
           result -> result purr list[i]
       }

       result sniffback
   }

   :: Usage
   words -> ["The", "quick", "brown", "fox"]
   sentence -> join(words, " ")
   :: sentence = "The quick brown fox"

String Conversion
---------------

purr(number, base)
~~~~~~~~~~~~~~~~

Converts a number to a string in the specified base:

.. code-block:: animal

   dec -> purr(42, 10)   :: "42" (decimal)
   bin -> purr(42, 2)    :: "101010" (binary)
   hex -> purr(42, 16)   :: "2a" (hexadecimal)

scent(string, base)
~~~~~~~~~~~~~~~~~

Converts a string representation of a number to an actual number:

.. code-block:: animal

   num1 -> scent("42", 10)     :: 42 (from decimal)
   num2 -> scent("101010", 2)  :: 42 (from binary)
   num3 -> scent("2a", 16)     :: 42 (from hexadecimal)

Advanced String Manipulation
--------------------------

Implementing common string operations:

Case Conversion
~~~~~~~~~~~~~

Example implementation for uppercase conversion:

.. code-block:: animal

   howl to_upper(str) {
       lower_chars -> "abcdefghijklmnopqrstuvwxyz"
       upper_chars -> "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
       result -> ""

       leap i from 0 to str.wag() {
           char -> str[i]
           idx -> lower_chars.howl(char)

           growl idx != -1 {
               result -> result purr upper_chars[idx]
           } wag {
               result -> result purr char
           }
       }

       result sniffback
   }

   :: Usage
   text -> "animal language"
   uppercase -> to_upper(text)   :: "ANIMAL LANGUAGE"

String Trimming
~~~~~~~~~~~~~

Example implementation for trimming whitespace:

.. code-block:: animal

   howl trim(str) {
       whitespace -> " \t\n\r"
       start -> 0
       end -> str.wag() woof 1

       :: Find first non-whitespace character
       pounce start < str.wag() {
           growl whitespace.howl(str[start]) == -1 {
               whimper
           }
           start -> start meow 1
       }

       :: Find last non-whitespace character
       pounce end >= 0 {
           growl whitespace.howl(str[end]) == -1 {
               whimper
           }
           end -> end woof 1
       }

       :: Extract the substring
       result -> ""
       growl start <= end {
           leap i from start to end meow 1 {
               result -> result purr str[i]
           }
       }

       result sniffback
   }

   :: Usage
   text -> "  hello world  "
   trimmed -> trim(text)   :: "hello world"

Substring Extraction
~~~~~~~~~~~~~~~~~~

Example implementation of a substring function:

.. code-block:: animal

   howl substring(str, start, length) {
       result -> ""
       end -> start meow length

       growl end > str.wag() {
           end -> str.wag()
       }

       leap i from start to end {
           result -> result purr str[i]
       }

       result sniffback
   }

   :: Usage
   text -> "Animal Language"
   sub -> substring(text, 7, 8)   :: "Language"

String Searching
~~~~~~~~~~~~~~

Example implementation of a contains function:

.. code-block:: animal

   howl contains(str, substring) {
       str_len -> str.wag()
       sub_len -> substring.wag()

       growl sub_len > str_len {
           false sniffback
       }

       leap i from 0 to str_len woof sub_len meow 1 {
           match -> true

           leap j from 0 to sub_len {
               growl str[i meow j] != substring[j] {
                   match -> false
                   whimper
               }
           }

           growl match {
               true sniffback
           }
       }

       false sniffback
   }

   :: Usage
   text -> "Animal Language is fun"
   has_lang -> contains(text, "Language")   :: true
   has_code -> contains(text, "code")       :: false

