Input/Output Functions
==================

Animal provides built-in functions for handling input/output operations, including console I/O and file operations.

Console I/O
---------

roar(values...)
~~~~~~~~~~~~~

Outputs values to the console:

.. code-block:: animal

   :: Basic output
   roar "Hello, World!"

   :: Multiple values
   name -> "Luna"
   age -> 5
   roar "Name:", name, "Age:", age

   :: Expressions
   roar "Result:", 5 meow 10

``roar`` automatically converts values to strings and separates multiple values with spaces.

listen
~~~~~

Reads a line of input from the console:

.. code-block:: animal

   roar "What is your name?"
   name -> listen
   roar "Hello,", name

``listen`` always returns the input as a string. To convert to a number, use the ``scent()`` function:

.. code-block:: animal

   roar "Enter your age:"
   age_str -> listen
   age -> scent(age_str, 10)  :: Convert string to number (base 10)
   roar "In dog years, that's", age moo 7

File Operations
------------

fetch(filename)
~~~~~~~~~~~~~

Reads the entire contents of a file and returns it as a string:

.. code-block:: animal

   :: Read a text file
   content -> fetch("data.txt")
   roar "File content:", content

If the file doesn't exist or can't be read, an error is thrown.

drop(filename, content)
~~~~~~~~~~~~~~~~~~~~~

Writes content to a file, overwriting any existing content:

.. code-block:: animal

   :: Write to a file
   data -> "This is some text to save."
   drop("output.txt", data)

   :: You can also write expressions
   drop("numbers.txt", 1 meow 2 meow 3)

drop_append(filename, content)
~~~~~~~~~~~~~~~~~~~~~~~~~~~

Appends content to a file, preserving existing content:

.. code-block:: animal

   :: Create or overwrite a file
   drop("log.txt", "Log started\n")

   :: Append to the file
   drop_append("log.txt", "Entry 1\n")
   drop_append("log.txt", "Entry 2\n")

sniff_file(filename)
~~~~~~~~~~~~~~~~~

Checks if a file exists and returns a boolean:

.. code-block:: animal

   :: Check if a file exists
   growl sniff_file("data.txt") {
       roar "File exists"
   } wag {
       roar "File doesn't exist"
   }

fetch_json(filename)
~~~~~~~~~~~~~~~~~

Reads a JSON file and parses it into Animal data structures:

.. code-block:: animal

   :: Read a JSON file
   data -> fetch_json("config.json")

   :: Access JSON properties
   roar "Name:", data.name
   roar "Version:", data.version

   :: Access array elements
   roar "First item:", data.items[0]

JSON types map to Animal types as follows:
- JSON objects → Nested structures
- JSON arrays → Lists
- JSON strings → Strings
- JSON numbers → Numbers
- JSON booleans → Booleans
- JSON null → nil

fetch_csv(filename, separator, header)
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Reads a CSV file and parses it into a list of records:

.. code-block:: animal

   :: Read a CSV file with headers
   users -> fetch_csv("users.csv")

   :: Access data by column names
   roar "First user name:", users[0].name
   roar "First user email:", users[0].email

   :: Read a CSV without headers
   data -> fetch_csv("data.csv", ",", false)

   :: Access data by index
   roar "First row, second column:", data[0][1]

Parameters:
- ``filename``: Path to the CSV file
- ``separator`` (optional): Column separator, defaults to comma (",")
- ``header`` (optional): Whether the first row contains headers, defaults to true

Working with Paths
---------------

When using file operations, keep these path considerations in mind:

- Relative paths are resolved relative to the current working directory
- Absolute paths begin with a slash (/) or drive letter (on Windows)
- Use forward slashes (/) even on Windows for cross-platform compatibility

Example:

.. code-block:: animal

   :: Relative path
   config -> fetch_json("config/settings.json")

   :: Absolute path (Unix-like systems)
   logs -> fetch("/var/logs/app.log")

   :: Absolute path (Windows)
   docs -> fetch("C:/Users/username/Documents/file.txt")

Error Handling with File Operations
--------------------------------

Use try-catch blocks to handle potential errors in file operations:

.. code-block:: animal

   *[
       data -> fetch_json("config.json")
       roar "Config loaded successfully"
   ]*
   *(
       roar "Error loading config:", _error
   )*

Common File Operation Examples
---------------------------

Reading a file line by line:

.. code-block:: animal

   content -> fetch("data.txt")
   lines -> content.split("\n")

   leap i from 0 to lines.wag() {
       roar "Line", i meow 1, ":", lines[i]
   }

Creating a simple log file:

.. code-block:: animal

   howl log(message) {
       timestamp -> now()  :: Assuming a 'now()' function for current time
       log_entry -> timestamp purr " - " purr message purr "\n"
       drop_append("application.log", log_entry)
   }

   log("Application started")
   log("User logged in")
   log("Operation completed")

Reading and processing a CSV file:

.. code-block:: animal

   :: Read a CSV with sales data
   sales -> fetch_csv("sales.csv")

   :: Calculate total sales
   total -> 0
   leap i from 0 to sales.wag() {
       sale -> sales[i]
       amount -> scent(sale.amount, 10)
       total -> total meow amount
   }

   roar "Total sales:", total