package stackerr_test

import (
	"errors"
	"log"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ovotech/stackerr"
)

var errTest = errors.New("test")

func TestHandle(t *testing.T) {
	t.Parallel()

	err := stackerr.Handle(errTest, "test", "some", "arguments")

	assert.Equal(t, "github.com/ovotech/stackerr_test.TestHandle.test(\"some\", \"arguments\"): test", err.Error())
}

func myFunction(input string) error {
	if _, err := strconv.Atoi(input); err != nil {
		// If you'd rather not include a helper, you can just use the raw Handle function with some arguments.
		// Using the helper cuts down repetition though.
		return stackerr.Handle(err, "atoi", input)
	}

	return nil
}

func ExampleHandle() {
	// You can just use Handle directly if you don't want to use the whole helper.
	err := myFunction("test")
	if err != nil {
		log.Panic(err)
	}

	// panic: main.myFunction.atoi("test"): strconv.Atoi: parsing "foo": invalid syntax
}
