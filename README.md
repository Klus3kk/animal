# Animal

The goal of this project is to create a new programming language with unique features and applications. The basic functions, operators, etc., are correlated to animal characteristics and sounds.

## Building

To build the Animal interpreter:

```bash
go build -o animal animal.go shell.go
```

This will create a binary called `animal`.

To run a `.anml` file:

```bash
./animal path/to/yourfile.anml
```

To launch the interactive REPL:

```bash
./animal
```

## Features

- Arithmetic operations using animal sounds (`meow`, `woof`, `moo`, etc.)
- Conditional statements: `growl`, `sniff`, `wag`
- Looping constructs: `leap`, `pounce`
- Printing output with `roar`
- Variable assignment via `->`
- Function definitions and calls via `howl`
- `nest` structures (custom object definitions)
- Lists with built-in methods

## Technologies

- Go language
- Custom Lexer
- Parser with AST generation
- Interpreter engine

## Instruction Manual

### Operators

| Symbol | Animal Word | Meaning        |
|--------|-------------|----------------|
| `*`    | `moo`       | Multiplication |
| `+`    | `meow`      | Addition       |
| `-`    | `woof`      | Subtraction    |
| `/`    | `drone`     | Division       |
| `%`    | `squeak`    | Modulo         |
| `^`    | `soar`      | Exponentiation |

### Brackets

- `()` → Round (grouping, function calls)
- `[]` → Square (lists)
- `{}` → Curly (blocks)

## Functions and Output

### `roar` — Print to console

```animal
roar "Hello", 2 meow 3
```

### `howl` — Define and call functions

```animal
howl greet(name) {
    roar "Hi", name
}

greet("Lucas")
```
## Full Documentation

You can explore the complete documentation in the [`docs/`](docs/) folder:

- [Introduction](docs/intro.md)
- [Syntax Guide](docs/syntax.md)
- [Language Features](docs/features.md)
- [Code Examples](docs/examples.md)
- [Standard Library](docs/stdlib.md)
