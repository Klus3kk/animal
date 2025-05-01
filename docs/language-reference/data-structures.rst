Data Structures
==============

Animal provides several built-in data structures for organizing and manipulating data efficiently.

Lists
----

Lists are ordered collections of items that can be of any type.

Creating Lists
~~~~~~~~~~~~

.. code-block:: animal

   :: Empty list
   empty_list -> []

   :: List with items
   numbers -> [1, 2, 3, 4, 5]

   :: Mixed-type list
   mixed -> ["apple", 42, true, [1, 2]]

Accessing List Elements
~~~~~~~~~~~~~~~~~~~~~

Use zero-based indexing with square brackets:

.. code-block:: animal

   fruits -> ["apple", "banana", "cherry"]

   first -> fruits[0]  :: "apple"
   second -> fruits[1]  :: "banana"

   :: Attempting to access an index out of bounds will cause an error
   :: error_item -> fruits[99]  :: Error: index out of bounds

Modifying Lists
~~~~~~~~~~~~~

Lists can be modified after creation:

.. code-block:: animal

   colors -> ["red", "green", "blue"]

   :: Add an item
   colors.sniff("yellow")  :: colors becomes ["red", "green", "blue", "yellow"]

   :: Modify an element
   colors[1] -> "purple"  :: colors becomes ["red", "purple", "blue", "yellow"]

List Methods
~~~~~~~~~~

Animal provides several built-in methods for lists:

.. code-block:: animal

   nums -> [3, 1, 4, 1, 5]

   :: Get length
   len -> nums.wag()  :: 5

   :: Reverse in place
   nums.snarl()  :: nums becomes [5, 1, 4, 1, 3]

   :: Find index
   idx -> nums.howl(4)  :: 2 (index of value 4)

   :: Random shuffle
   nums.prowl()  :: Randomly rearranges the elements

   :: See the list functions documentation for more

Nests (Custom Object Structures)
------------------------------

The ``nest`` keyword lets you define custom data structures similar to classes or objects.

Defining a Nest
~~~~~~~~~~~~~

.. code-block:: animal

   nest Dog {
       name
       breed
       age

       howl bark() {
           roar this.name, "says: Woof!"
       }

       howl get_human_age() {
           this.age moo 7 sniffback
       }
   }

Creating Nest Instances
~~~~~~~~~~~~~~~~~~~~~

.. code-block:: animal

   :: Create a new Dog instance
   my_dog -> Dog()

   :: Set properties
   my_dog.name -> "Rex"
   my_dog.breed -> "Shepherd"
   my_dog.age -> 3

   :: Call methods
   my_dog.bark()  :: Prints "Rex says: Woof!"
   human_age -> my_dog.get_human_age()  :: 21

The ``this`` Keyword
~~~~~~~~~~~~~~~~~

Within a nest method, use ``this`` to refer to the current instance:

.. code-block:: animal

   nest Counter {
       value

       howl increment() {
           this.value -> this.value meow 1
       }

       howl get() {
           this.value sniffback
       }
   }

Nest Limitations
~~~~~~~~~~~~~~

In the current version of Animal:

- Nests do not support inheritance
- All properties are public
- Constructor methods are not supported (initialize properties after creation)
- Properties must be primitive types, lists, or other nest instances

Dictionaries (Key-Value Pairs)
----------------------------

Animal does not have a built-in dictionary type, but you can simulate one using nests:

.. code-block:: animal

   nest Dictionary {
       keys
       values

       howl init() {
           this.keys -> []
           this.values -> []
       }

       howl set(key, value) {
           idx -> this.keys.howl(key)

           growl idx == -1 {
               :: Key doesn't exist, add it
               this.keys.sniff(key)
               this.values.sniff(value)
           } wag {
               :: Key exists, update value
               this.values[idx] -> value
           }
       }

       howl get(key) {
           idx -> this.keys.howl(key)

           growl idx == -1 {
               nil sniffback
           } wag {
               this.values[idx] sniffback
           }
       }
   }

   :: Usage
   dict -> Dictionary()
   dict.init()
   dict.set("name", "Luna")
   dict.set("age", 5)
   roar dict.get("name")  :: Prints "Luna"

Data Structure Patterns
---------------------

Implementing a Stack
~~~~~~~~~~~~~~~~~~

.. code-block:: animal

   nest Stack {
       items

       howl init() {
           this.items -> []
       }

       howl push(item) {
           this.items.sniff(item)
       }

       howl pop() {
           growl this.items.wag() == 0 {
               nil sniffback
           }

           idx -> this.items.wag() woof 1
           item -> this.items[idx]
           :: TODO: Implement list removal
           item sniffback
       }

       howl peek() {
           growl this.items.wag() == 0 {
               nil sniffback
           }

           this.items[this.items.wag() woof 1] sniffback
       }

       howl size() {
           this.items.wag() sniffback
       }
   }