# StackErr

Make your errors beautiful, because writing them is boring enough.

* Automatically wrap your errors with the package and calling function name.
* Keep it simple, and just plainly wrap your errors with `Handle(err, "myFunction", "some", "args")`.
* Full control over what is/isn't included in your errors. Include as much detail as you need to help with debugging.
* Performant. StackErr is only invoked when handling errors.
* Fully interoperable with other error handlers.

## Installation

```shell
go get github.com/ovotech/stackerr
```

## Usage

```go
package main

import (
	"log"
	"strconv"

	"github.com/ovotech/stackerr"
)

func myFunction(input string) error {
	if _, err := strconv.Atoi(input); err != nil {
		return stackerr.Handle(err, "atoi", input, "some", "added", "details")
	}

	return nil
}

func main() {
	if err := myFunction("foo"); err != nil {
		log.Panic(err)
	}

	// panic: main.myFunction.atoi("foo", "some", "added", "details"): strconv.Atoi: parsing "foo": invalid syntax
}
```

### Use with `wrapcheck`

If you use [golangci-lint](https://github.com/golangci/golangci-lint) and the wrapcheck linter, you should add this
package to your config to prevent errors:

```yaml
linters-settings:
  wrapcheck:
    ignorePackageGlobs:
      - github.com/ovotech/stackerr
```
