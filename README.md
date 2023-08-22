# StackErr

Make your errors beautiful, because writing them is boring enough.

* Automatically wrap your errors with the package and calling function name.
* Keep your code DRY by declaring the things you want in your errors once.
* Or keep it simple, and just plainly wrap your errors with `Handle(err, "your", "args")`.
* Full control over what is/isn't included in your errors. Include as much detail as you need to help with debugging.
* Fully interoperable with other error handlers.
* Customisable to suit your needs.

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

	. "github.com/ovotech/stackerr"
)

func myFunction(input string) error {
	// Create a new StackErr to with the input parameter as we want to know that in all of our errors.
	stackErr := NewStackErr(input)

	if _, err := strconv.Atoi(input); err != nil {
		// Handle the error with the mandatory function name "atoi" to tell us where the error came from, and some 
		// optional context which would help with debugging.
		return stackErr.Handle(err, "atoi", "some", "context")
	}

	return nil
}

func main() {
	if err := myFunction("foo"); err != nil {
		log.Panic(err)
	}

	// panic: main.myFunction("foo").atoi("some", "context"): strconv.Atoi: parsing "foo": invalid syntax
}
```

<details>
<summary>Basic usage without instantiating</summary>
You can use StackErr without instantiating it by using the Handle function by itself.

```go
package main

import (
	"log"
	"strconv"

	. "github.com/ovotech/stackerr"
)

func myFunction(input string) error {
	if _, err := strconv.Atoi(input); err != nil {
		return Handle(err, "atoi", input)
	}

	return nil
}

func main() {
	if err := myFunction("foo"); err != nil {
		log.Panic(err)
	}

	// panic: main.myFunction.atoi("foo"): strconv.Atoi: parsing "foo": invalid syntax
}
```

</details>


<details>
<summary>Customising StackErr</summary>
You can customise StackErr by instantiating a new struct, and then using Copy in functions you wish to use it in.

```go
package main

import (
	"log"
	"strconv"

	. "github.com/ovotech/stackerr"
)

var myStackErr = StackErr{
	Separator:   " -> ",
	Punctuation: "\'",
}

func myFunction(input string) error {
	stackErr := myStackErr.Copy(input)

	if _, err := strconv.Atoi(input); err != nil {
		return stackErr.Handle(err, "atoi", "hello", "world")
	}

	return nil
}

func main() {
	if err := myFunction("foo"); err != nil {
		log.Panic(err)
	}

	// Note the punctuation and separator:
	// panic: main.myFunction('foo').atoi('hello', 'world') -> strconv.Atoi: parsing "foo": invalid syntax
}
```

</details>

### Use with `wrapcheck`

If you use [golangci-lint](https://github.com/golangci/golangci-lint) and the wrapcheck linter, you should add this
package to your config to prevent errors:

```yaml
linters-settings:
  wrapcheck:
    ignorePackageGlobs:
      - github.com/ovotech/stackerr
```
