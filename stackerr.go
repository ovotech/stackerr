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

// Handle an error, providing the package & function name, locator, and optional arguments.
// Wraps the error e.g. "mypackage.myFunction.locator("arg1", "arg2") -> error.
func Handle(err error, locator string, args ...string) error {
	functionName := "unknown"

	// Get the name of the function invoking this function.
	if pc, _, _, ok := runtime.Caller(1); ok {
		details := runtime.FuncForPC(pc)
		functionName = details.Name()
	}

	argString := ""
	if len(args) > 0 {
		argString = fmt.Sprintf(
			"(%s%s%s)",
			punctuation,
			strings.Join(args, fmt.Sprintf("%s, %s", punctuation, punctuation)),
			punctuation,
		)
	}

	return fmt.Errorf(
		"%s.%s%s%s%w",
		functionName,
		locator,
		argString,
		separator,
		err,
	)
}
