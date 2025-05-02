Standard Library Overview
=====================

The Animal language comes with a comprehensive standard library of built-in functions and capabilities that extend the core language. This page provides an overview of the various modules and functions available.

Core Components
-------------

The standard library is organized into several categories:

1. **Math Functions** - Mathematical operations and utilities
2. **List Functions** - Tools for working with and manipulating lists
3. **String Functions** - String manipulation and processing
4. **I/O Functions** - Reading from and writing to console and files
5. **Random Functions** - Random number generation and randomization tools

How to Use the Standard Library
-----------------------------

All standard library functions are available globally without requiring any import statements. Simply call them in your code:

.. code-block:: animal

   :: Using math function
   result -> max(5, 10)

   :: Using list function
   items -> [1, 2, 3]
   shuffled -> tumble(items)

   :: Using I/O function
   content -> fetch("data.txt")

Standard Library Reference
------------------------

Math Functions
~~~~~~~~~~~~

.. list-table::
   :header-rows: 1
   :widths: 25 75

   * - Function
     - Description
   * - ``max(a, b)``
     - Returns the larger of two numbers
   * - ``min(a, b)``
     - Returns the smaller of two numbers
   * - ``abs(x)``
     - Returns the absolute value of a number
   * - ``purr(num, base)``
     - Converts a number to string in a specified base
   * - ``scent(str, base)``
     - Converts a string to a number in a specified base

See :doc:`math-functions` for detailed documentation.

List Functions
~~~~~~~~~~~~

.. list-table::
   :header-rows: 1
   :widths: 25 75

   * - Function
     - Description
   * - ``paw(x, min, max)``
     - Clamps a number between min and max values
   * - ``nuzzle(a, b)``
     - Merges two lists or concatenates two strings
   * - ``burrow(n)``
     - Creates a list of n nil elements
   * - ``perch(list)``
     - Returns all permutations of a list
   * - ``lick(list)``
     - Flattens a nested list
   * - ``howl(list, item)``
     - Finds the index of an item in a list
   * - ``chase(x, n)``
     - Repeats element x n times into a list
   * - ``trace(list)``
     - Creates a running sum of list elements
   * - ``trail(list)``
     - Creates prefixes of a list
   * - ``pelt(value, times)``
     - Repeats a value as a string
   * - ``howlpack(list, item)``
     - Returns all indices where item appears
   * - ``nest(value, depth)``
     - Nests a value to the specified depth

See :doc:`list-functions` for detailed documentation.

String Functions
~~~~~~~~~~~~~

.. list-table::
   :header-rows: 1
   :widths: 25 75

   * - Function
     - Description
   * - ``purr``
     - String concatenation operator
   * - ``pelt(value, times)``
     - Repeats a value as a string
   * - ``nuzzle(str1, str2)``
     - Joins two strings (function form of purr)

See :doc:`string-functions` for detailed documentation.

I/O Functions
~~~~~~~~~~

.. list-table::
   :header-rows: 1
   :widths: 25 75

   * - Function
     - Description
   * - ``roar(values...)``
     - Prints values to the console
   * - ``listen``
     - Reads a line from the console
   * - ``fetch(filename)``
     - Reads a file and returns its contents
   * - ``drop(filename, content)``
     - Writes content to a file
   * - ``drop_append(filename, content)``
     - Appends content to a file
   * - ``sniff_file(filename)``
     - Checks if a file exists
   * - ``fetch_json(filename)``
     - Reads and parses a JSON file
   * - ``fetch_csv(filename, sep, header)``
     - Reads and parses a CSV file

See :doc:`io-functions` for detailed documentation.

Random Functions
~~~~~~~~~~~~~

.. list-table::
   :header-rows: 1
   :widths: 25 75

   * - Function
     - Description
   * - ``pounce(min, max)``
     - Generates a random integer in range
   * - ``stalk(list)``
     - Returns a random element from a list
   * - ``tumble(list)``
     - Returns a randomly shuffled list

These functions are included in the :doc:`math-functions` documentation.

Extending the Standard Library
----------------------------

The Animal language standard library is designed to be extensible. If you're implementing your own Animal interpreter or contributing to the project, you can add new standard library functions by:

1. Implementing the function in the appropriate Go file in the ``core/std/`` directory
2. Registering the function in ``core/std/register.go``
3. Adding appropriate tests and documentation

Example of a standard library function implementation:

.. code-block:: go

   // In core/std/example.go
   package std

   import (
       "fmt"
   )

   // AnimalExampleFunction is a standard library function
   func AnimalExampleFunction(args []interface{}) interface{} {
       if len(args) != 1 {
           return fmt.Errorf("example_function expects 1 argument")
       }

       // Function implementation
       // ...

       return result
   }

   // In core/std/register.go
   func RegisterStandardLibrary(symbolTable *core.SymbolTable) {
       // ...existing registrations...

       // Register the new function
       symbolTable.Set("example_function", AnimalExampleFunction)
   }

