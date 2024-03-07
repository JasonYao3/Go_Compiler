package repl

import (
	"bufio"
	"fmt"
	"go_compiler/compiler"

	// "go_compiler/evaluator"
	"go_compiler/lexer"
	"go_compiler/object"
	"go_compiler/parser"
	"go_compiler/vm"
	"io"
)

/*
the REPL reads input, sends it to the interpreter for evaluation, prints the result/output of the
interpreter and starts again. Read, Evaluate, Print, Loop.
*/

const PROMPT = ">> "

// TODO: Initialize the lexer with an io.Reader and the filename.
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	// env := object.NewEnvironment()
	// macroEnv := object.NewEnvironment()

	constants := []object.Object{}
	globals := make([]object.Object, vm.GlobalSize)
	symbolTable := compiler.NewSymbolTable()
	for i, v := range object.Builtins {
		symbolTable.DefineBuiltin(i, v.Name)
	}

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		// evaluator.DefineMacros(program, macroEnv)
		// expanded := evaluator.ExpandedMacros(program, macroEnv)
		// evaluated := evaluator.Eval(expanded, env)
		// if evaluated != nil {
		// 	io.WriteString(out, evaluated.Inspect())
		// 	io.WriteString(out, "\n")
		// }

		comp := compiler.NewWithState(symbolTable, constants)
		err := comp.Compile(program)
		if err != nil {
			fmt.Fprintf(out, "Compilation failed:\n %s\n", err)
			continue
		}

		code := comp.Bytecode()
		constants = code.Constants

		machine := vm.NewWithGlobalsStore(code, globals)
		err = machine.Run()
		if err != nil {
			fmt.Fprintf(out, "Executing bytecode failed:\n %s\n", err)
			continue
		}

		lastPopped := machine.LastPoppedStackElem()
		io.WriteString(out, lastPopped.Inspect())
		io.WriteString(out, "\n")

	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
