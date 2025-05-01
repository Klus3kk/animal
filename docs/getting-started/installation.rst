Installation
============

This guide will help you set up the Animal language interpreter on your system.

Requirements
-----------

- Go 1.18 or higher
- Git

Installing from Source
---------------------

1. Clone the repository

   .. code-block:: bash

      git clone https://github.com/your-username/animal.git
      cd animal

2. Build the interpreter

   .. code-block:: bash

      go build -o animal.exe ./cmd/animal

3. Verify the installation

   .. code-block:: bash

      ./animal.exe --version

The command should display the version of Animal that you've just installed.

Adding Animal to Your PATH
-------------------------

To use Animal from any directory, add it to your system PATH:

On Linux/macOS:

.. code-block:: bash

   # Add this line to your .bashrc or .zshrc
   export PATH=$PATH:/path/to/animal/directory

On Windows:

1. Right-click on 'This PC' and select 'Properties'
2. Click on 'Advanced system settings'
3. Click on 'Environment Variables'
4. Under 'System variables', select 'Path' and click 'Edit'
5. Click 'New' and add the path to your Animal directory
6. Click 'OK' to close all dialogs

Using the Animal REPL
-------------------

You can start an interactive REPL (Read-Eval-Print Loop) by running Animal without arguments:

.. code-block:: bash

   animal

This will open an interactive session where you can type Animal code and see the results immediately.

Installing VS Code Extension
--------------------------

For a better development experience, you can install the Animal VS Code extension:

1. Open VS Code
2. Go to Extensions (Ctrl+Shift+X)
3. Search for "Animal Language"
4. Click "Install"

The extension provides syntax highlighting, code snippets, and integration with the Animal interpreter.

Troubleshooting
-------------

If you encounter any issues during installation:

1. Make sure you have the correct Go version: ``go version``
2. Check that your ``GOPATH`` is set correctly: ``go env GOPATH``
3. Ensure you have build permissions in your directory

For further assistance, please file an issue on our GitHub repository.