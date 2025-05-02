Operators
=========

Animal language uses a unique set of operators, many of which are named after animal sounds or behaviors.

Arithmetic Operators
------------------

The core arithmetic operations in Animal use animal-themed keywords:

.. list-table::
   :header-rows: 1
   :widths: 15 15 40 30

   * - Animal Operator
     - Traditional
     - Description
     - Example
   * - ``meow``
     - ``+``
     - Addition
     - ``5 meow 3`` → ``8``
   * - ``woof``
     - ``-``
     - Subtraction
     - ``10 woof 4`` → ``6``
   * - ``moo``
     - ``*``
     - Multiplication
     - ``6 moo 7`` → ``42``
   * - ``drone``
     - ``/``
     - Division
     - ``20 drone 5`` → ``4``
   * - ``squeak``
     - ``%``
     - Modulo (remainder)
     - ``17 squeak 5`` → ``2``
   * - ``soar``
     - ``^``
     - Exponentiation
     - ``2 soar 3`` → ``8``

These operators work with both integer and floating-point numbers:

.. code-block:: animal

   float_result -> 10.5 meow 2.3   :: 12.8
   integer_div -> 7 drone 2        :: 3.5 (division always returns float)
   mixed -> 5 moo 1.5              :: 7.5

Unary Operators
-------------

Animal supports unary plus and minus:

.. code-block:: animal

   positive -> +5   :: Same as 5
   negative -> -10  :: Negation

String Operations
---------------

For string operations, Animal provides:

.. list-table::
   :header-rows: 1
   :widths: 15 40 45

   * - Operator
     - Description
     - Example
   * - ``purr``
     - Concatenation
     - ``"hello" purr " world"`` → ``"hello world"``

String and numeric values can be concatenated with automatic conversion:

.. code-block:: animal

   greeting -> "Age: " purr 25   :: "Age: 25"

Assignment Operators
-----------------

Animal uses an arrow for assignment:

.. list-table::
   :header-rows: 1
   :widths: 15 85

   * - Operator
     - Description
   * - ``->``
     - Assigns the value on the right to the variable on the left

Example:

.. code-block:: animal

   name -> "Luna"
   age -> 5

   :: Multiple assignments
   a -> b -> c -> 10   :: All three variables get value 10

Compound assignment operators (like ``+=``, ``-=``) are not supported in the current version.

Comparison Operators
-----------------

Animal uses standard comparison operators:

.. list-table::
   :header-rows: 1
   :widths: 15 40 45

   * - Operator
     - Description
     - Example
   * - ``==``
     - Equal to
     - ``5 == 5`` → ``true``
   * - ``!=``
     - Not equal to
     - ``5 != 3`` → ``true``
   * - ``>``
     - Greater than
     - ``10 > 5`` → ``true``
   * - ``<``
     - Less than
     - ``3 < 7`` → ``true``
   * - ``>=``
     - Greater than or equal to
     - ``5 >= 5`` → ``true``
   * - ``<=``
     - Less than or equal to
     - ``4 <= 6`` → ``true``

Comparison operators work with numbers and strings (lexicographical comparison):

.. code-block:: animal

   :: Number comparison
   temperature -> 28
   is_hot -> temperature > 25   :: true

   :: String comparison
   name -> "Alice"
   alphabetical -> name < "Bob"  :: true (A comes before B)

Logical Operators
--------------

For combining boolean expressions:

.. list-table::
   :header-rows: 1
   :widths: 15 40 45

   * - Operator
     - Description
     - Example
   * - ``and``
     - Logical AND
     - ``true and false`` → ``false``
   * - ``or``
     - Logical OR
     - ``true or false`` → ``true``

These can be combined for complex conditions:

.. code-block:: animal

   temp -> 28
   is_sunny -> true

   go_swimming -> temp > 25 and is_sunny   :: true

Operator Precedence
-----------------

Operators in Animal follow this precedence order (highest to lowest):

1. Parentheses ``()``
2. Unary operators ``+``, ``-``
3. Exponentiation ``soar``
4. Multiplication, Division, Modulo ``moo``, ``drone``, ``squeak``
5. Addition, Subtraction ``meow``, ``woof``
6. String concatenation ``purr``
7. Comparison operators ``>``, ``<``, ``>=``, ``<=``, ``==``, ``!=``
8. Logical operators ``and``, ``or``
9. Assignment ``->``

You can use parentheses to override default precedence:

.. code-block:: animal

   :: Default precedence (multiplication before addition)
   result1 -> 2 meow 3 moo 4   :: 2 + (3 * 4) = 14

   :: With parentheses
   result2 -> (2 meow 3) moo 4  :: (2 + 3) * 4 = 20


