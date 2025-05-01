Quick Start Guide
===============

This guide will walk you through the basics of Animal language programming.

Your First Animal Program
-----------------------

Let's start with a simple "Hello World" program:

1. Create a file named ``hello.anml`` with the following content:

   .. code-block:: animal

      roar "Hello, Animal World!"

2. Run the program:

   .. code-block:: bash

      animal hello.anml

You should see ``Hello, Animal World!`` printed to the console.

Basic Syntax
-----------

Variable Assignment
~~~~~~~~~~~~~~~~~

Assign values to variables using the ``->`` operator:

.. code-block:: animal

   name -> "Luna"
   age -> 5
   is_active -> true

Arithmetic Operations
~~~~~~~~~~~~~~~~~~~

Animal uses animal sounds for arithmetic operations:

.. code-block:: animal

   :: Addition
   sum -> 5 meow 3   :: sum = 8

   :: Subtraction
   diff -> 10 woof 4   :: diff = 6

   :: Multiplication
   product -> 6 moo 7   :: product = 42

   :: Division
   quotient -> 20 drone 5   :: quotient = 4

   :: Modulo
   remainder -> 17 squeak 5   :: remainder = 2

   :: Exponentiation
   power -> 2 soar 3   :: power = 8

Conditional Statements
~~~~~~~~~~~~~~~~~~~~

Use ``growl``, ``sniff``, and ``wag`` for conditionals:

.. code-block:: animal

   score -> 85

   growl score > 90 {
       roar "Excellent!"
   } sniff score > 70 {
       roar "Good job!"
   } wag {
       roar "Keep practicing."
   }

Loops
~~~~

For loops use ``leap``:

.. code-block:: animal

   leap i from 1 to 5 {
       roar i
   }

While loops use ``pounce``:

.. code-block:: animal

   count -> 0
   pounce count < 3 {
       roar "Count:", count
       count -> count meow 1
   }

Functions
~~~~~~~~

Define functions with ``howl``:

.. code-block:: animal

   howl greet(name) {
       message -> "Hello, " purr name purr "!"
       roar message
   }

   greet("Alice")   :: Prints: Hello, Alice!

Return values using ``sniffback``:

.. code-block:: animal

   howl square(n) {
       n moo n sniffback
   }

   result -> square(4)
   roar "Square:", result   :: Prints: Square: 16

Lists
~~~~

Create and manipulate lists:

.. code-block:: animal

   fruits -> ["apple", "banana", "cherry"]
   fruits.sniff("orange")   :: Add item

   roar fruits[0]   :: Access by index
   roar fruits.wag()   :: Get length

Next Steps
---------

Now that you know the basics, try exploring:

- :doc:`/language-reference/syntax` for detailed syntax rules
- :doc:`/language-reference/data-structures` for more on lists and nests
- :doc:`/standard-library/overview` for built-in functions
- :doc:`/getting-started/examples` for more code examples