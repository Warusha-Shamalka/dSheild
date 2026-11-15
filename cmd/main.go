package main

// Go exports package-level names and methods whose first character is an uppercase
// Unicode letter. That is why another package can use os.Open, os.Args, fmt.Println,
// and parser.TokenizerDockerRules. Lowercase names such as parser.readLine are private
// to their package. This rule also applies to types, variables, constants, and fields.
// These are package APIs, not language built-ins. Built-ins such as len and append
// are lowercase and need no import. Use Go documentation or editor autocomplete to
// check the exact spelling and signature of a function; do not guess its casing.
import (
	"fmt"
	"os"
	"github.com/Warusha-Shamalka/dSheild/internals/parser"
)

func main(){
	// os.Args is a []string. Index 0 identifies the executable; index 1 is the
	// first user-supplied argument. len is a built-in that returns an int.
	// Check the length before indexing so missing arguments cannot cause a panic.
	if len(os.Args) != 2{
		// fmt.Println writes its arguments to standard output (os.Stdout), with
		// spaces between arguments and a newline at the end.
		fmt.Println("Provide file name as the first argument")
		// os.Exit(code int) terminates the entire process immediately. Zero means
		// success; nonzero means failure. Here, 2 is our chosen usage-error status.
		// Deferred calls do not run when os.Exit terminates the process.
		os.Exit(2)
	}

	fileName := os.Args[1]
     
	fmt.Println("Docker Parser Started - main.go")
	parser.TokenizerDockerRules(fileName)
}
