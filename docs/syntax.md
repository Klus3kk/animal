# 🐾 Animal Language - Syntax Guide

This guide walks you through the core syntax rules of the **Animal Language**, from assigning variables to calling functions, all using animal-inspired grammar.

---

## 📦 Variable Assignment
Use `->` to assign values:
```anml
x -> 5
name -> "Luna"
```

---

## 🧮 Arithmetic Operators
| Operator | Meaning        | Example             |
|----------|----------------|---------------------|
| `meow`   | Addition (`+`) | `5 meow 3` → `8`     |
| `woof`   | Subtraction (`-`)| `5 woof 2` → `3`    |
| `moo`    | Multiplication (`*`) | `3 moo 4` → `12` |
| `drone`  | Division (`/`) | `8 drone 2` → `4`    |
| `squeak` | Modulo (`%`)   | `7 squeak 3` → `1`   |
| `soar`   | Exponentiation (`^`) | `2 soar 3` → `8` |
| `purr`   | Concatenation (strings) | `"hi" purr " there"` → `"hi there"` |

---

## 🗣 Print & Input
```anml
roar "Hello, world!"
roar x, y, "value"

name -> listen
```

---

## 🧠 Conditionals
```anml
growl x > 5 {
    roar "Big!"
} sniff x == 5 {
    roar "Exactly five."
} wag {
    roar "Small."
}
```

---

## 🔁 Loops
### `leap` → For loop
```anml
leap i from 0 to 5 {
    roar i
}
```

### `pounce` → While loop
```anml
x -> 0
pounce x < 3 {
    roar x
    x -> x meow 1
}
```

---

## 🐺 Functions
Define using `howl`:
```anml
howl square(n) {
    return n moo n
}

roar square(5)  # 25
```

---

## 🐕 Dot Notation
Access fields and methods on `nest` objects:
```anml
d.name -> "Bolt"
d.speak()
```

---

Next up: see [features.md](features.md) for all language features and behaviors!

