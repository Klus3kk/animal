# 🧬 Animal Language - Features

This document outlines the expressive features of the Animal Language beyond basic syntax — including **lists**, **custom structures**, and **function mechanics**.

---

## 📃 Lists
Animal supports lists with a full toolbox of methods:
```anml
mylist -> [1, 2, 3]
mylist.sniff(4)      # append 4
mylist.wag()         # get length (→ 4)
mylist.snarl()       # reverse the list
mylist.prowl()       # shuffle items
mylist.howl(2)       # remove item at index 2
```

### 🐍 Advanced List Methods
```anml
nested -> [[1, 2], [3]]
nested.lick()        # flatten → [1, 2, 3]

nums -> [1, 4, 5, 7]
nums.howl_at(4)      # filter all >= 4 → [4, 5, 7]

nums.nest(2)         # chunk into groups → [[1, 4], [5, 7]]
```

---

## 🐚 Nests (Custom Structures)
A `nest` is like a class:
```anml
nest Dog {
    name
    howl speak() {
        roar this.name, "says Woof!"
    }
}

d -> Dog()
d.name -> "Rex"
d.speak()
```

### 🧠 Behavior
- Fields are declared directly.
- Methods are defined with `howl` inside the nest.
- You can access fields with `this.` inside methods.

---

## 🐺 Function Mechanics
Functions are declared with `howl`:
```anml
howl add(a, b) {
    return a meow b
}
roar add(2, 3)  # → 5
```

- Functions can return values.
- They support argument passing.
- Nest methods can be called with dot syntax.

---

## 🧠 Context & Scoping
- Global variables are shared across the program.
- Functions and methods create their own local scope.
- Nest instances store their fields in their own object-specific context.

---

Up next: check out [examples.md](examples.md) for complete Animal programs!

