# Animal

The goal of this project is to create a new programming language with unique features and applications. The basic functions, operators etc. will be correlated to animals characteristics and sounds.  
You can test the language by writing:

```bash
go build
```

And after that:

```bash
./animal
```

## Features

* Arithmetic operations with animal sounds (e.g., `meow`, `woof`, `moo`)
* Conditional statements (`growl`, `sniff`, `wag`)
* Looping constructs (`leap`, `pounce`)
* Printing output (`roar`)
* Variable assignments (`x -> 5`)
* Function definitions and calls (`howl`)

## Technologies

* Go language,
* Lexer,
* Parser,
* Tokens.

## Instruction

**Operators:**

* `"*"` - moo
* `"+"` - meow
* `"-"` - woof
* `"/"` - drone
* `"%"` - squeak
* `"^"` - soar

**Brackets:**

* `()` - round
* `[]` - square
* `{}` - curly

**Functions:**

* `"roar"` — print to console
  ```animal
  roar "Hello", 2 meow 3
  ```

* `"howl"` — define and call functions
  ```animal
  howl greet(name) {
      roar "Hi", name
  }

  greet("Lucas")
  ```
