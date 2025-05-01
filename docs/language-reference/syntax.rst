Syntax Reference
===============

This page details the complete syntax rules for the Animal language.

Code Structure
-------------

Animal programs consist of statements, separated implicitly by newlines or explicitly by semicolons (``;``).

Comments
-------

Comments in Animal start with ``::`` and continue to the end of the line:

.. code-block:: animal

   :: This is a comment
   x -> 5  :: This is an inline comment

Variables
--------

Variable Naming
~~~~~~~~~~~~~

Variable names:

- Must start with a letter or underscore
- Can contain letters, numbers, and underscores
- Are case-sensitive
- Cannot be a reserved keyword

.. code-block:: animal

   valid_name -> 10
   _also_valid -> 20
   ThisIsValid -> 30

   :: Invalid examples:
   :: 1name -> 40        (starts with number)
   :: name-with-dash -> 50 (contains dash)
   :: growl -> 60        (reserved keyword)

Variable Assignment
~~~~~~~~~~~~~~~~

The assignment operator is ``->``:

.. code-block:: animal

   name -> "Luna"
   age -> 5

   :: Multiple assignments
   a -> b -> c -> 10  :: All variables get value 10

Variable Type Annotations
~~~~~~~~~~~~~~~~~~~~~~

Optional type annotations use a colon after the variable name:

.. code-block:: animal

   count: int -> 5
   name: string -> "Buddy"
   is_active: bool -> true

Data Types
---------

Animal supports the following primitive data types:

Integers
~~~~~~~

Whole numbers:

.. code-block:: animal

   a -> 42
   b -> -7

Floats
~~~~~

Numbers with decimal points:

.. code-block:: animal

   pi -> 3.14159
   half -> 0.5

Booleans
~~~~~~~

``true`` or ``false`` values:

.. code-block:: animal

   is_valid -> true
   has_error -> false

Strings
~~~~~~

Text enclosed in single or double quotes:

.. code-block:: animal

   name -> "Alice"
   message -> 'Hello, world!'

Lists
~~~~

Ordered collections of items, enclosed in square brackets:

.. code-block:: animal

   numbers -> [1, 2, 3, 4]
   mixed -> ["apple", 5, true]

Operators
--------

Arithmetic Operators
~~~~~~~~~~~~~~~~~

.. list-table::
   :header-rows: 1
   :widths: 20 20 60

   * - Animal Operator
     - Equivalent
     - Description
   * - ``meow``
     - ``+``
     - Addition
   * - ``woof``
     - ``-``
     - Subtraction
   * - ``moo``
     - ``*``
     - Multiplication
   * - ``drone``
     - ``/``
     - Division
   * - ``squeak``
     - ``%``
     - Modulo (remainder)
   * - ``soar``
     - ``^``
     - Exponentiation
   * - ``purr``
     - ``+`` (for strings)
     - String concatenation

Comparison Operators
~~~~~~~~~~~~~~~~~

.. list-table::
   :header-rows: 1
   :widths: 25 75

   * - Operator
     - Description
   * - ``==``
     - Equal to
   * - ``!=``
     - Not equal to
   * - ``>``
     - Greater than
   * - ``<``
     - Less than
   * - ``>=``
     - Greater than or equal to
   * - ``<=``
     - Less than or equal to

Logical Operators
~~~~~~~~~~~~~~

.. list-table::
   :header-rows: 1
   :widths: 25 75

   * - Operator
     - Description
   * - ``and``
     - Logical AND
   * - ``or``
     - Logical OR
   * - ``not``
     - Logical NOT

Precedence
~~~~~~~~

Operators follow this precedence order (highest to lowest):

1. Parentheses ``()``
2. Exponentiation ``soar``
3. Unary plus/minus ``+``, ``-``
4. Multiplication, Division, Modulo ``moo``, ``drone``, ``squeak``
5. Addition, Subtraction ``meow``, ``woof``
6. Comparison operators ``>``, ``<``, ``>=``, ``<=``, ``==``, ``!=``
7. Logical operators ``and``, ``or``
8. Assignment ``->``

Parentheses can be used to override the default precedence:

.. code-block:: animal

   :: Default precedence: multiplication before addition
   result -> 2 meow 3 moo 4   :: 2 + (3 * 4) = 14

   :: Override with parentheses
   result -> (2 meow 3) moo 4   :: (2 + 3) * 4 = 20

Reserved Keywords
---------------

The following are reserved keywords in Animal and cannot be used as variable names:

.. code-block:: animal

   :: Control flow
   growl, sniff, wag, pounce, leap

   :: Functions/structures
   howl, nest

   :: Input/output
   roar, listen

   :: Return
   sniffback

   :: File operations
   fetch, drop, drop_append, sniff_file, fetch_json, fetch_csv

   :: Control
   whimper, hiss

   :: Other
   mimic, _

Identifiers that match these keywords cannot be used as variable or function names.