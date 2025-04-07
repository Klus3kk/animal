# 🧰 Animal Language - Standard Library

This document lists the built-in methods and behaviors for **lists**, **objects**, and other core elements of the Animal Language.

---

## 📃 List Methods
All lists support the following methods:

### `.sniff(value)`
Appends a value to the list.
```anml
mylist.sniff(42)
```

### `.howl(index)`
Removes the element at the given index.
```anml
mylist.howl(1)
```

### `.wag()`
Returns the length of the list.
```anml
len -> mylist.wag()
```

### `.snarl()`
Reverses the list in-place.
```anml
mylist.snarl()
```

### `.prowl()`
Shuffles the list randomly.
```anml
mylist.prowl()
```

### `.lick()`
Flattens a nested list.
```anml
[[1, 2], [3]].lick()  # → [1, 2, 3]
```

### `.howl_at(threshold)`
Filters the list, keeping only values >= threshold.
```anml
nums.howl_at(4)  # [4, 5, 6]
```

### `.nest(size)`
Chunks the list into sublists of given size.
```anml
[1, 2, 3, 4].nest(2)  # [[1, 2], [3, 4]]
```

---

## 🐚 Object Methods
No global object methods exist yet — all are defined in `nest` blocks using `howl()`.

Example:
```anml
nest Dog {
    name
    howl speak() {
        roar this.name, "barks."
    }
}
```

---

## 🗣 Built-In Functions
### `roar` — print to output
```anml
roar "Hello", 123
```

### `listen` — read a line from input
```anml
name -> listen
```

---

Coming soon:
- Math helpers (random, min, max)
- String manipulation
- Date/time
- File I/O