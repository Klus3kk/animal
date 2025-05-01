Control Flow
===========

Animal provides several keywords for controlling the flow of program execution, including conditionals and loops.

Conditional Statements
--------------------

growl-sniff-wag (if-else if-else)
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

The Animal language uses ``growl``, ``sniff``, and ``wag`` for conditional execution:

.. code-block:: animal

   growl condition {
       // Executed if condition is true
   } sniff another_condition {
       // Executed if the first condition is false and another_condition is true
   } wag {
       // Executed if all previous conditions are false
   }

Examples:

.. code-block:: animal

   score -> 85

   growl score >= 90 {
       roar "Excellent!"
   } sniff score >= 70 {
       roar "Good job!"
   } sniff score >= 50 {
       roar "Passed"
   } wag {
       roar "Try again"
   }

You can use just ``growl`` without ``sniff`` or ``wag``:

.. code-block:: animal

   growl user_logged_in {
       roar "Welcome back!"
   }

Or ``growl`` with ``wag`` but no ``sniff``:

.. code-block:: animal

   growl has_access {
       roar "Access granted"
   } wag {
       roar "Access denied"
   }

mimic (switch)
~~~~~~~~~~~~

For matching a value against multiple cases, use ``mimic``:

.. code-block:: animal

   day -> "Monday"

   mimic day {
       "Monday" -> roar "Start of work week"
       "Friday" -> roar "End of work week"
       "Saturday" -> roar "Weekend!"
       "Sunday" -> roar "Weekend!"
       _ -> roar "Mid-week"
   }

The underscore ``_`` acts as a default case, which executes when no other case matches.

Looping Constructs
----------------

leap (for loop)
~~~~~~~~~~~~~

The ``leap`` keyword creates counting loops:

.. code-block:: animal

   leap i from 0 to 5 {
       roar i  :: Prints 0, 1, 2, 3, 4
   }

The loop variable automatically increments by 1 each iteration. The loop executes from the start value (inclusive) to the end value (exclusive).

pounce (while loop)
~~~~~~~~~~~~~~~~

The ``pounce`` keyword creates loops that continue while a condition is true:

.. code-block:: animal

   count -> 0
   pounce count < 3 {
       roar "Count is:", count
       count -> count meow 1
   }

Loop Control
----------

whimper (break)
~~~~~~~~~~~~~

Use ``whimper`` to exit a loop early:

.. code-block:: animal

   leap i from 0 to 10 {
       growl i == 5 {
           whimper  :: Exit the loop when i equals 5
       }
       roar i  :: Only prints 0, 1, 2, 3, 4
   }

hiss (continue)
~~~~~~~~~~~~

Use ``hiss`` to skip the rest of the current iteration and continue with the next:

.. code-block:: animal

   leap i from 0 to 5 {
       growl i == 2 {
           hiss  :: Skip the rest of this iteration
       }
       roar i  :: Prints 0, 1, 3, 4 (skips 2)
   }

Error Handling
------------

try-catch
~~~~~~~~

Animal uses symbolic syntax for try-catch blocks:

.. code-block:: animal

   *[
       :: Code that might throw an error
       *{ "something wrong" }*  :: Explicitly throw an error
   ]*
   *(
       roar "Caught error:", _error  :: _error contains the error message
   )*

The ``*[ ... ]*`` syntax defines a try block, and the ``*( ... )*`` syntax defines the catch block.
Inside a catch block, the special variable ``_error`` contains the error message.

Explicitly throwing errors:

.. code-block:: animal

   howl divide(a, b) {
       growl b == 0 {
           *{ "Division by zero" }*  :: Throw an error
       }
       a drone b sniffback
   }

Best Practices
------------

1. Use descriptive conditions in ``growl`` statements to make code more readable
2. Keep loop bodies simple and focused on a single task