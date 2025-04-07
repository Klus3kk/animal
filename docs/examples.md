# 🧪 Animal Language - Code Examples

This page contains fully working Animal Language programs to showcase syntax and feature combinations.

---

## 🖨️ Hello World
```anml
roar "Hello, Animal World!"
```

---

## 🔁 Loop from 1 to 5
```anml
leap i from 1 to 5 {
    roar "Step:", i
}
```

---

## 🔢 FizzBuzz
```anml
leap i from 1 to 15 {
    growl i squeak 3 == 0 and i squeak 5 == 0 {
        roar "FizzBuzz"
    } sniff i squeak 3 == 0 {
        roar "Fizz"
    } sniff i squeak 5 == 0 {
        roar "Buzz"
    } wag {
        roar i
    }
}
```

---

## 🐺 Function Example
```anml
howl greet(name) {
    roar "Hello", name
}

greet("Fox")
```

---

## 🐚 Nest Example
```anml
nest Cat {
    name
    howl meow() {
        roar this.name, "says Meow!"
    }
}

c -> Cat()
c.name -> "Whiskers"
c.meow()
```

---

## 📃 List Manipulation
```anml
nums -> [1, 2, 3]
nums.sniff(4)
nums.snarl()
roar nums
```

---

## 🧠 Factorial Function
```anml
howl factorial(n) {
    growl n <= 1 {
        return 1
    } wag {
        return n moo factorial(n woof 1)
    }
}

roar factorial(5)  # → 120
```

More programs coming soon!

