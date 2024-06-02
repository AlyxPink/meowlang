# 🐱 Welcome to MeowLang 🐾

The purrfect educational programming language for curious developers. 🐈

**MeowLang** is a fun, cat-themed programming language implemented in Go.

It is designed to be simple, easy to understand, and fun to use. The goal of this project is to learn about language design and implementation, and to have fun along the way. 😸

# 📚 Purposes

This project is for educational purposes, to learn about language design and implementation, and to have fun along the way. 😸

## 📖 Backstory

Few days before my first commit, I had literally zero knowledge on how to make this happen. I did not know what "AST", "Tokens", "Lexer", or any of that meant. I thought it was impossible for me to make a programming language.

**I thought I was not smart enough to do it. But I was wrong.**

And I swear that if I was able to do it, you can do it too. 🩷 If you need some help to understand how this works, feel free to send me an e-mail or open a GitHub Discussion. I will do my best to help you.

## 🐾 Features

- 🐱 `meow` Declare variables
- 🐾 `purr` Print/output
- 😼 `hiss` Conditional/if statements
- 🐈‍⬛ `scratch` Loops (while loops)
- 🐾 `paw` Define functions
- 🐾 `claw` Return values from functions
- 💤 `nap` Sleep for a specified duration

## ✅ Basic Features Checklist

- [x] **Lexer**: Convert input into tokens
- [x] **Parser**: Convert tokens into an Abstract Syntax Tree (AST)
- [x] **Interpreter**: Execute the AST
- [x] **Variable Declaration**: Implement `lick` for declaring variables
- [x] **Print/Output**: Implement `purr` for printing/outputting values
- [ ] **Conditionals**: Implement `hiss-growl` for if-else statements
- [ ] **Loops**: Implement `scratch` for while loops
- [x] **Function Definitions**: Implement `meow` for defining functions
- [x] **Function Calls**: Implement `meow double(a) { claw 2 * a }; purr double(10);` for calling functions
- [x] **Return Statement**: Implement `claw` for returning values from functions
- [ ] **Sleep Function**: Implement `nap` for sleeping
- [x] **Comments**: Implement `//`, `/*` and `*/` for comments

## 🏗️ Project Structure

- `cmd/meowlang/main.go`: Entry point of the application.
- `lexer/lexer.go`: Lexer implementation.
- `parser/parser.go`: Parser implementation.
- `ast/ast.go`: AST node definitions.
- `interpreter/interpreter.go`: Interpreter implementation.
- `token/token.go`: Token definitions.
- `util/util.go`: Utility functions.

## 🔨 How to Build

To build the project, run:

```sh
go build -o meowlang ./cmd/meowlang
```

## 🚀 How to Run

To run a MeowLang program, use the following command:

```sh
./meowlang <filename>
```

## 📜 Example Code

Here's a sneak peek at what a MeowLang program might look like:

**ℹ️ Note**: Not all features are implemented yet. This is just a planned syntax.

```meowlang
// This is a simple MeowLang program

/* Not everything is implemented yet
   Consider this as the planned syntax */

// Variable declaration
lick a = 5
lick b = 10
lick result = a + b

// Print statement
purr "Meow!"
purr 125
purr "Meow" + " " + ":3"
purr result

// Function definition
meow add(p, q) {
    claw p + q
}

// Conditional
hiss (a < b) {
    purr "a is less than b"
} growl {
    purr "a is not less than b"
}

// Loop
scratch (a < b) {
    purr a
    a = a + 1
    nap(1) // Sleep for 1 unit of time
}

// Function call
lick result = add(10, 5)
purr "Result of addition: " + result
```

## 🛠️ Contributing

I don't think it's worth extending the language to make it a full-fledged programming language. The goal of this project is to learn about language design and implementation, and to have fun along the way. 😸

Howver if you want to contribute, you can add some features that are not implemented yet. You can also improve the code quality, add tests, or improve the documentation.

Bonus points if you add more cat-themed features 🐾

## 📜 License

This project is licensed under the AGPL-3.0 License. You are free to use, modify, and distribute this project, provided that any derivative works also comply with the AGPL-3.0 License. For more details, see the [LICENSE](LICENSE) file.
