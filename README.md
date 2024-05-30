# 🐱 MeowLang

Welcome to **MeowLang**! 🐾 The purrfect programming language for cat lovers and curious developers. 🐈

**MeowLang** is a fun, cat-themed programming language implemented in Go. This project is for educational purposes, to learn about language design and implementation, and to have a little fun along the way. 😸

## Features 🐾

- 🐱 **meow**: Declare variables
- 🐾 **purr**: Print/output
- 😼 **hiss**: Conditional/if statements
- 🐈‍⬛ **scratch**: Loops (while loops)
- 🐾 **paw**: Define functions
- 🐾 **claw**: Return values from functions
- 💤 **nap**: Sleep for a specified duration

## Project Structure 🏗️

- `cmd/meowlang/main.go`: Entry point of the application.
- `lexer/lexer.go`: Lexer implementation.
- `parser/parser.go`: Parser implementation.
- `ast/ast.go`: AST node definitions.
- `interpreter/interpreter.go`: Interpreter implementation.
- `token/token.go`: Token definitions.
- `util/util.go`: Utility functions.

## How to Build 🔨

To build the project, run:

```sh
go build -o meowlang ./cmd/meowlang
```

## How to Run 🚀

To run a MeowLang program, use the following command:

```sh
./meowlang <filename>
```

## Basic Features Checklist ✅

- [x] **Lexer**: Convert input into tokens
- [x] **Parser**: Convert tokens into an Abstract Syntax Tree (AST)
- [x] **Interpreter**: Execute the AST
- [x] **Variable Declaration**: Implement `lick` for declaring variables
- [ ] **Print/Output**: Implement `purr` for printing/outputting values
- [ ] **Conditionals**: Implement `hiss` for if-else statements
- [ ] **Loops**: Implement `scratch` for while loops
- [ ] **Function Definitions**: Implement `meow` for defining functions
- [ ] **Return Statement**: Implement `claw` for returning values from functions
- [ ] **Sleep Function**: Implement `nap` for sleeping

## Example Code 📜

Here's a sneak peek at what a MeowLang program might look like:

```meowlang
# This is a simple MeowLang program

# Variable declaration
lick a = 5
lick b = 10

# Function definition
meow add(p, q) {
    claw p + q
}

# Conditional
hiss (a < b) {
    purr "a is less than b"
} else {
    purr "a is not less than b"
}

# Loop
scratch (a < b) {
    purr a
    a = a + 1
    nap(1)  # Sleep for 1 unit of time
}

# Function call
lick result = add(a, b)
purr "Result of addition: " + result
```

## Contributing 🛠️

Feel free to contribute to MeowLang! Whether it's fixing bugs, adding new features, or just suggesting ideas, every bit of help is appreciated. Just remember to keep it pawsitive! 😸

## License 📜

This project is licensed under the AGPL-3.0 License. You are free to use, modify, and distribute this project, provided that any derivative works also comply with the AGPL-3.0 License. For more details, see the [LICENSE](LICENSE) file.
