# Compiler in Go

### Introduction

---
The Go Compiler is a project aimed at translating Go source code into an intermediate bytecode representation that can be executed by a custom-built virtual machine. This project demystifies the process of writing compilers and virtual machines, making it accessible for developers to understand and implement these complex systems.

Currently extending the programming language designed in [_Writing An Compiler In Go_](https://compilerbook.com) by [Thorsten Ball](https://github.com/mrnugget).

---

### Features

- **Lexical Analysis**: Tokenizes the source code into meaningful symbols.
- **Parsing**: Converts token sequences into an abstract syntax tree (AST).
- **Semantic Analysis**: Ensures the AST adheres to the rules of the Go language.
- **Code Generation**: Translates the AST into bytecode for the virtual machine.
- **Error Handling**: Provides detailed error messages for lexical, syntactic, and semantic errors.
- **Optimization**: Performs basic optimizations on the bytecode for better performance.

---

### Built With

- Go 1.13+

---

### It Supports

- **Let Statements**: Define identifiers using the `let` keyword.
- **If Expressions**: Conditional expressions using `if` and `else`.
- **Function Literals**: Define functions with parameters and block statements.
- **Integer Literals**: Support for integer operations.
- **Boolean Literals**: Support for boolean values and expressions.
- **Grouped Expressions**: Use parentheses to group expressions.
- **Array Literals**: Create and manipulate arrays.
- **Hash Literals**: Create and access hash maps.
