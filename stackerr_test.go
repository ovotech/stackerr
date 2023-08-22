package stackerr_test

import (
	"errors"
	"log"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/ovotech/stackerr"
)

var errTest = errors.New("test")

func TestHandle(t *testing.T) {
	t.Parallel()

	err := Handle(errTest, "test", "some", "arguments")

	assert.Equal(t, "github.com/ovotech/stackerr_test.TestHandle.test(\"some\", \"arguments\"): test", err.Error())
}

func myStackErrHandleFunction(input string) error {
	// Instantiate a new error helper at the top of your functions.
	// Include any relevant data here that you'd like all errors to inherit.
	stackErr := NewStackErr(input)

	if _, err := strconv.Atoi(input); err != nil {
		// Instead of returning a raw error, you now ask your helper to handle it for you.
		// After the error is a location, which can be used to pinpoint the exact location of this error.
		return stackErr.Handle(err, "atoi")
	}

	return nil
}

func ExampleNewStackErr() {
	err := myStackErrHandleFunction("test")
	if err != nil {
		// Panic! mymodule.myStackErrHandleFunction(\"test\").atoi: error.
		log.Panic(err)
	}
}

func myHandleFunction(input string) error {
	if _, err := strconv.Atoi(input); err != nil {
		// If you'd rather not include a helper, you can just use the raw Handle function with some arguments.
		// Using the helper cuts down repetition though.
		return Handle(err, "atoi", input)
	}

	return nil
}

func ExampleHandle() {
	// You can just use Handle directly if you don't want to use the whole helper.
	err := myHandleFunction("test")
	if err != nil {
		// Panic! mymodule.myStackErrHandleFunction.atoi(\"test\"): error.
		log.Panic(err)
	}
}

func ExampleStackErr_Handle() {
	err := myStackErrHandleFunction("test")
	if err != nil {
		// Panic! mymodule.myStackErrHandleFunction(\"test\").atoi: error.
		log.Panic(err)
	}
}

func ExampleStackErr_Copy() {
	// Create & configure your top level error helper.
	originalStackErr := NewStackErr("my", "original", "helper")
	originalStackErr.Separator = "'"
	originalStackErr.Punctuation = ": "

	// If you want to persist your changes, just make a new copy for subsequent helpers.
	// It doesn't copy the arguments, only the Separator and Punctuation.
	_ = originalStackErr.Copy("some", "new", "file")
}
