Math Functions
=============

Animal provides a set of built-in mathematical functions in its standard library. These functions help with common mathematical operations without requiring you to implement them yourself.

Basic Math Functions
------------------

max(a, b)
~~~~~~~~

Returns the larger of two numbers:

.. code-block:: animal

   result -> max(5, 10)  :: Returns 10
   result -> max(-3, -7)  :: Returns -3

min(a, b)
~~~~~~~~

Returns the smaller of two numbers:

.. code-block:: animal

   result -> min(5, 10)  :: Returns 5
   result -> min(-3, -7)  :: Returns -7

abs(x)
~~~~~

Returns the absolute value of a number:

.. code-block:: animal

   result -> abs(5)    :: Returns 5
   result -> abs(-10)  :: Returns 10
   result -> abs(0)    :: Returns 0

Number Conversions
----------------

purr(number, base)
~~~~~~~~~~~~~~~~

Converts a number to a string representation in the specified base:

.. code-block:: animal

   result -> purr(42, 10)   :: Returns "42" (decimal)
   result -> purr(42, 2)    :: Returns "101010" (binary)
   result -> purr(42, 16)   :: Returns "2a" (hexadecimal)

The base must be between 2 and 36, inclusive.

scent(string, base)
~~~~~~~~~~~~~~~~~

Converts a string representation of a number in the specified base to a number:

.. code-block:: animal

   result -> scent("42", 10)     :: Returns 42 (from decimal)
   result -> scent("101010", 2)  :: Returns 42 (from binary)
   result -> scent("2a", 16)     :: Returns 42 (from hexadecimal)

The base must be between 2 and 36, inclusive.

Random Number Generation
----------------------

pounce(min, max)
~~~~~~~~~~~~~~

Generates a random integer between min and max, inclusive:

.. code-block:: animal

   result -> pounce(1, 6)  :: Random number between 1 and 6 (like a die roll)
   result -> pounce(0, 100)  :: Random number between 0 and 100

stalk(list)
~~~~~~~~~~

Returns a random element from a list:

.. code-block:: animal

   colors -> ["red", "green", "blue", "yellow"]
   result -> stalk(colors)  :: Returns a random color from the list

tumble(list)
~~~~~~~~~~

Returns a new list with elements from the input list in random order:

.. code-block:: animal

   numbers -> [1, 2, 3, 4, 5]
   shuffled -> tumble(numbers)  :: Returns a randomly shuffled version of the numbers list

Example Usage
-----------

Here's an example that combines several math functions to create a simple number guessing game:

.. code-block:: animal

   :: Number guessing game
   secret -> pounce(1, 100)  :: Generate random number between 1 and 100
   attempts -> 0

   roar "I'm thinking of a number between 1 and 100."

   guessed -> false
   pounce !guessed {
       roar "Enter your guess:"
       guess_str -> listen
       guess -> scent(guess_str, 10)
       attempts -> attempts meow 1

       difference -> abs(guess woof secret)

       growl guess == secret {
           roar "Correct! You found the number in", attempts, "attempts."
           guessed -> true
       } sniff difference <= 5 {
           roar "Very close!"
       } sniff difference <= 10 {
           roar "Getting warmer."
       } sniff guess < secret {
           roar "Too low."
       } wag {
           roar "Too high."
       }
   }

Mathematical Operations in Animal
------------------------------

Remember that Animal has its own unique operators for mathematical operations:

- Addition: ``a meow b``
- Subtraction: ``a woof b``
- Multiplication: ``a moo b``
- Division: ``a drone b``
- Modulo: ``a squeak b``
- Exponentiation: ``a soar b``

These operators can be used alongside the standard library math functions.

