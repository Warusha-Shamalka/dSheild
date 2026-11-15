package parser

import (
	"bufio"
	"fmt"
	"os"
) 
type Instruction struct{
	InstructionId string
	LineNumber    int
	Lexeme        string
}

type File struct{
	FileType 	  string
	FileName 	  string
	Instructions  []Instruction
}

// error is a predefined Go interface type with one method: Error() string.
// A value implementing that method can describe a failure. An error value of nil
// conventionally means success. Functions return errors explicitly; callers check
// them instead of using try/catch. This function returns an error to its caller.
func readLine(path string) error{
	// os.Open(path string) returns (*os.File, error). It opens an existing file
	// for reading. *os.File is a pointer to a file handle, not the file's contents.
	// The handle implements io.Reader, so a buffered scanner can read from it.
	// On success err is nil; on failure the file result is nil and err explains
	// why opening failed. The := declaration infers both variable types.
	file , err := os.Open(path)
	// Return the opening error before trying to read or close an invalid handle.
	if err != nil{
		return err
	}
	// defer registers a function call to run when this surrounding function exits,
	// including an early return or normal panic unwinding. Multiple deferred calls
	// run in reverse registration order. The function and arguments are evaluated
	// now, but the call happens later. Here, Close releases the open file handle.
	// Register this only after opening succeeds. os.Exit bypasses deferred calls.
	defer file.Close()

	// bufio.NewScanner accepts an io.Reader and returns a *bufio.Scanner.
	// It creates a buffered token reader; by default a token is one physical line.
	// It does not parse Dockerfile instructions or join continued lines for us.
	// The default maximum token buffer is 64 KiB, so sufficiently long lines can
	// cause an error. Scanner.Buffer can configure a larger limit before scanning.
	scanner := bufio.NewScanner(file)

	// Scan() returns bool. Each true result means the scanner has obtained another
	// token (a line with the default splitting rule). The loop body handles that line.
	// Empty lines are valid tokens. A final nonempty line without a newline is also
	// returned. Scan returns false at end-of-file or on an error; false alone does
	// not tell us which occurred. Scan may read ahead into its internal buffer.
	for scanner.Scan(){
		// Text() returns the current token as a string. With the default line
		// splitting rule, the trailing LF or CRLF line ending is removed.
		// Text does not advance the scanner; the next Scan call does that.
		fmt.Println(scanner.Text())
	}

	// Err() returns the scanning error, or nil after normal end-of-file. It does
	// not report normal EOF as an error. Returning it distinguishes a completed
	// read from a failed one. The deferred Close runs before control returns to
	// our caller, including when a scanning error is returned.
	return scanner.Err()
}


func TokenizerDockerRules(fileName string){
	fmt.Println("This is a Tokernizer")		
	err := readLine(fileName)
	if err != nil{
		// os.Stderr is an exported variable of type *os.File, initially connected
		// to standard error (file descriptor 2). It is a separate output stream
		// for diagnostics, usually shown in the terminal but independently redirectable.
		// fmt.Fprint writes values using their default formats to an io.Writer;
		// here that writer is os.Stderr. Fprint does not append a newline.
		// fmt.Fprintln(writer, values...) instead puts spaces between values and
		// appends a newline. Both return (bytesWritten int, err error).
		// fmt.Fprintf is the variant that accepts a format string such as "%s".
		fmt.Fprint(os.Stderr, "Failed to read the file ",err)
		// Stop the process with our chosen error status. os.Exit does not execute
		// defers. In this flow readLine has already returned, so its deferred file
		// close has already run before we get here.
		os.Exit(2)
	}
}
