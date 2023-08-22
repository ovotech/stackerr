package stackerr

import (
	"fmt"
	"runtime"
	"strings"
)

const (
	separator   = ": "
	punctuation = "\""
)

// argsToString takes a slice of args and punctuation, and returns a string.
// If no args are passed, it will return an empty string.
// If 1+ args are passed, it will wrap them in parentheses and each arg in the punctuation e.g. ("foo", "bar").
func argsToString(args []string, punctuation string) string {
	argString := ""
	if len(args) > 0 {
		argString = fmt.Sprintf(
			"(%s%s%s)",
			punctuation,
			strings.Join(args, fmt.Sprintf("%s, %s", punctuation, punctuation)),
			punctuation,
		)
	}

	return argString
}

// handle is the internal mechanism for wrapping errors with extra contextual information automatically.
func handle(
	err error,
	loc string,
	sharedArgs []string,
	invokedArgs []string,
	punctuation string,
	separator string,
) error {
	functionName := "unknown"

	// Get the name of the function invoking this function.
	if pc, _, _, ok := runtime.Caller(2); ok { //nolint:gomnd
		details := runtime.FuncForPC(pc)
		functionName = details.Name()
	}

	// Convert both the invokedArgs and handle sharedArgs into strings.
	sharedArgsString := argsToString(sharedArgs, punctuation)
	invokedArgsString := argsToString(invokedArgs, punctuation)

	return fmt.Errorf(
		"%s%s.%s%s%s%w",
		functionName,
		invokedArgsString,
		loc,
		sharedArgsString,
		separator,
		err,
	)
}

// Handle an error, providing the module, function name and optional arguments.
// Wraps the error e.g. "mymodule.myFunction.functionName("arg1", "arg2") -> error.
func Handle(err error, functionName string, args ...string) error {
	return handle(err, functionName, args, nil, punctuation, separator)
}

type StackErr struct {
	// Arguments defined when the StackErr was instantiated.
	arguments []string

	Separator   string // Separator used between errors e.g. " -> "
	Punctuation string // Punctuation surrounding arguments.
}

// Handle an error.
func (s *StackErr) Handle(err error, functionName string, args ...string) error {
	return handle(err, functionName, args, s.arguments, s.Punctuation, s.Separator)
}

// Copy the existing StackErr's separator and punctuation, and create a new StackErr.
func (s *StackErr) Copy(args ...string) *StackErr {
	// Copy the StackErr struct's Separator and Punctuation, and return a new one.
	return &StackErr{
		arguments:   args,
		Separator:   s.Separator,
		Punctuation: s.Punctuation,
	}
}

// NewStackErr takes optional arguments, and returns a new StackErr with the default settings.
func NewStackErr(args ...string) *StackErr {
	return &StackErr{
		arguments:   args,
		Separator:   separator,
		Punctuation: punctuation,
	}
}
