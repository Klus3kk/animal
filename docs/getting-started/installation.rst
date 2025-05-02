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

   Alternatively, you can use the Go install command:

   .. code-block:: bash

      go install ./cmd/animal@latest

3. Verify the installation

   .. code-block:: bash

      ./animal --version

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

Platform-Specific Notes
---------------------

Windows
^^^^^^^

- The executable will be named ``animal.exe``
- You may need to run the command prompt as Administrator when building

Linux
^^^^^

- You may need to add execute permissions: ``chmod +x animal``
- Consider placing the binary in ``/usr/local/bin`` for system-wide installation

macOS
^^^^^

- You may need to bypass Gatekeeper the first time you run Animal
- Homebrew installation may be available in the future