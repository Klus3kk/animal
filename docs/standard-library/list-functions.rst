List Functions
=============

The Animal language provides a rich set of built-in functions for working with lists. This page documents all list-related functions in the standard library.

List Methods
-----------

Lists in Animal support several methods using dot notation:

sniff(value)
~~~~~~~~~~

Appends a value to the end of the list:

.. code-block:: animal

   list -> [1, 2, 3]
   list.sniff(4)  :: Result: [1, 2, 3, 4]

wag()
~~~~

Returns the length of the list:

.. code-block:: animal

   list -> [1, 2, 3, 4]
   length -> list.wag()  :: Result: 4

howl(item)
~~~~~~~~~

Finds the index of an item in the list, returns -1 if not found:

.. code-block:: animal

   list -> [10, 20, 30]
   idx -> list.howl(20)  :: Result: 1
   not_found -> list.howl(50)  :: Result: -1

snarl()
~~~~~~

Reverses the list in place:

.. code-block:: animal

   list -> [1, 2, 3]
   list.snarl()  :: Result: [3, 2, 1]

prowl()
~~~~~~

Randomly shuffles the list in place:

.. code-block:: animal

   list -> [1, 2, 3, 4, 5]
   list.prowl()  :: Result: Elements in random order

lick()
~~~~~

Flattens a nested list into a single-level list:

.. code-block:: animal

   nested -> [[1, 2], [3], [4, 5]]
   flat -> nested.lick()  :: Result: [1, 2, 3, 4, 5]

nest(size)
~~~~~~~~~

Chunks the list into sublists of the given size:

.. code-block:: animal

   list -> [1, 2, 3, 4, 5, 6]
   chunks -> list.nest(2)  :: Result: [[1, 2], [3, 4], [5, 6]]

howl_at(threshold)
~~~~~~~~~~~~~~~~

Filters the list, keeping only values greater than or equal to the threshold:

.. code-block:: animal

   nums -> [1, 5, 3, 8, 4]
   filtered -> nums.howl_at(4)  :: Result: [5, 8, 4]

Global List Functions
-------------------

The following functions can be called on lists without using dot notation:

paw(x, min, max)
~~~~~~~~~~~~~~

Clamps a number between minimum and maximum values:

.. code-block:: animal

   result -> paw(15, 1, 10)  :: Result: 10 (value clamped to max)
   result -> paw(-5, 0, 100)  :: Result: 0 (value clamped to min)
   result -> paw(7, 0, 10)  :: Result: 7 (value within range)

nuzzle(a, b)
~~~~~~~~~~~

Merges two lists or concatenates two strings:

.. code-block:: animal

   list1 -> [1, 2]
   list2 -> [3, 4]
   merged -> nuzzle(list1, list2)  :: Result: [1, 2, 3, 4]

   str1 -> "hello"
   str2 -> "world"
   joined -> nuzzle(str1, str2)  :: Result: "helloworld"

burrow(n)
~~~~~~~~

Creates a list of n nil elements:

.. code-block:: animal

   empty -> burrow(3)  :: Result: [nil, nil, nil]

perch(list)
~~~~~~~~~~

Returns all permutations of the list:

.. code-block:: animal

   items -> [1, 2, 3]
   perms -> perch(items)  :: Result: [[1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], [3,2,1]]

chase(element, times)
~~~~~~~~~~~~~~~~~~

Creates a list with the element repeated the specified number of times:

.. code-block:: animal

   repeated -> chase("a", 3)  :: Result: ["a", "a", "a"]

trace(numbers)
~~~~~~~~~~~~

Creates a running sum of the list elements:

.. code-block:: animal

   nums -> [1, 2, 3, 4]
   sums -> trace(nums)  :: Result: [1, 3, 6, 10]

trail(list)
~~~~~~~~~~

Creates prefixes of the list:

.. code-block:: animal

   items -> [1, 2, 3]
   prefixes -> trail(items)  :: Result: [[1], [1, 2], [1, 2, 3]]

pelt(value, times)
~~~~~~~~~~~~~~~~

Repeats a value as a string a specified number of times:

.. code-block:: animal

   result -> pelt(123, 3)  :: Result: "123123123"

howlpack(list, item)
~~~~~~~~~~~~~~~~~~

Returns all indices where item appears in the list:

.. code-block:: animal

   items -> [1, 2, 3, 2, 4, 2]
   indices -> howlpack(items, 2)  :: Result: [1, 3, 5]

nest(value, depth)
~~~~~~~~~~~~~~~~

Nests a value to the specified depth:

.. code-block:: animal

   result -> nest(42, 3)  :: Result: [[[42]]]

Best Practices
------------

- Use ``list.wag()`` to check if a list is empty before accessing elements
- Use ``burrow()`` to pre-allocate lists of known size
- Use ``howl()`` to check if an item exists in a list before trying to access it
- Use ``lick()`` to simplify processing of nested data