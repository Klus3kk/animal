Modules
=======

Animal supports a basic module system that allows you to organize and reuse code across multiple files. This page explains how to work with modules in Animal.

Importing Code
------------

The ``%bestiary`` directive allows you to import code from other Animal files:

.. code-block:: animal

   %bestiary "path/to/your/file.anml"

This imports and executes all the code in the specified file, making its functions and variables available in the current file.

Export Control with Shelter
-------------------------

By default, all variables and functions from an imported file are available in the importing file. To control what gets exported, use the ``!shelter`` directive:

.. code-block:: animal

   :: In math_utils.anml
   !shelter -> ["add", "multiply"]  :: Only export these functions

   howl add(a, b) {
       a meow b sniffback
   }

   howl multiply(a, b) {
       a moo b sniffback
   }

   howl internal_secret(x) {  :: This won't be exported
       x meow 1000 sniffback
   }

In this example, only the ``add`` and ``multiply`` functions will be available when importing this file. The ``internal_secret`` function remains private to the module.

Importing Example
--------------

.. code-block:: animal

   :: In main.anml
   %bestiary "math_utils.anml"

   :: These work because they're in the shelter
   result1 -> add(2, 3)      :: Result: 5
   result2 -> multiply(4, 5) :: Result: 20

   :: This would fail because it's not exported
   :: result3 -> internal_secret(5)  :: Error! not visible

Module Path Resolution
-------------------

When importing modules:

1. Relative paths are resolved relative to the current file
2. If the path doesn't include a directory (just a filename), Animal searches:
   - The current directory
   - The standard library directory (if available)

Examples:

.. code-block:: animal

   :: Import from same directory
   %bestiary "helper.anml"

   :: Import from a subdirectory
   %bestiary "utils/math.anml"

   :: Import from parent directory
   %bestiary "../common/shared.anml"

Avoiding Circular Imports
----------------------

Animal has basic protection against circular imports. If a file attempts to import another file that's already being imported in the chain, the second import is ignored.

However, it's best to design your module structure to avoid circular dependencies:

.. code-block:: text

   Good structure (hierarchical):

   main.anml
   ├── utils.anml
   └── features.anml
       └── sub_feature.anml

   Problematic structure (circular):

   a.anml → imports b.anml
   b.anml → imports a.anml

Best Practices for Modules
------------------------

1. **Group Related Functionality**

   Place related functions and variables in the same module:

   .. code-block:: animal

      :: math.anml - Math utilities
      !shelter -> ["add", "subtract", "multiply", "divide"]

      howl add(a, b) { a meow b sniffback }
      howl subtract(a, b) { a woof b sniffback }
      howl multiply(a, b) { a moo b sniffback }
      howl divide(a, b) { a drone b sniffback }

2. **Explicitly Control Exports**

   Always use ``!shelter`` to explicitly declare what your module exports:

   .. code-block:: animal

      !shelter -> ["public_function", "public_variable"]

3. **Use Descriptive Module Names**

   Choose clear, descriptive names for module files:

   - ``string_utils.anml`` for string manipulation functions
   - ``data_processing.anml`` for data processing functions
   - ``ui_components.anml`` for user interface components

4. **Document Module Interfaces**

   Include comments at the top of module files describing their purpose and exports:

   .. code-block:: animal

      :: =============================================
      :: list_utils.anml
      :: Utilities for working with lists
      ::
      :: Exports:
      ::  - filter(list, predicate_func)
      ::  - map(list, transform_func)
      ::  - reduce(list, combine_func, initial)
      :: =============================================
      !shelter -> ["filter", "map", "reduce"]

5. **Minimize Side Effects**

   Module imports are executed when imported, so minimize side effects:

   .. code-block:: animal

      :: Not ideal - has side effects on import
      roar "Module imported!"  :: Prints when imported

      :: Better - initialization function that can be called when needed
      !shelter -> ["initialize"]

      howl initialize() {
          roar "Module initialized!"
      }

Advanced Module Patterns
---------------------

**Namespace Modules**

Create namespaces by returning objects from modules:

.. code-block:: animal

   :: math.anml
   !shelter -> ["Math"]

   Math -> {
       "PI": 3.14159,
       "E": 2.71828
   }

   :: main.anml
   %bestiary "math.anml"
   roar Math.PI  :: Access through namespace

**Configuration Modules**

Create configuration files that can be imported:

.. code-block:: animal

   :: config.anml
   !shelter -> ["CONFIG"]

   CONFIG -> {
       "VERSION": "1.0.0",
       "DEBUG": true,
       "API_ENDPOINT": "https://api.example.com"
   }

   :: main.anml
   %bestiary "config.anml"

   growl CONFIG.DEBUG {
       roar "Debug mode is enabled"
   }

Limitations
---------

The current module system in Animal has some limitations:

1. No support for circular dependencies
2. No lazy loading - all code is executed on import
3. Limited namespace management
4. No built-in package manager

Future Enhancements
-----------------

Future versions of Animal may introduce:

- Better dependency management
- Runtime module loading
- Proper namespaces
- Package management system