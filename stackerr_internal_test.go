package stackerr

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var errTest = errors.New("test")

func TestNewStackErr(t *testing.T) {
	t.Parallel()

	stackErr := NewStackErr("some", "arguments")

	assert.ElementsMatch(t,
		[]string{"some", "arguments"},
		stackErr.arguments,
		"Check arguments are being passed to error arguments",
	)
	assert.Equal(t, "\"", stackErr.Punctuation, "Default Punctuation should be '\"'")
	assert.Equal(t, ": ", stackErr.Separator, "Default Separator should be \": \"")
}

func TestStackErr_Handle(t *testing.T) {
	t.Parallel()

	stackErr := &StackErr{
		Punctuation: "\"",
		Separator:   ": ",
	}

	err := stackErr.Handle(errTest, "test", "some", "arguments")

	assert.ErrorContains(t, err, "github.com/ovotech/stackerr.TestStackErr_Handle.test(\"some\", \"arguments\"): test")
}

func TestStackErr_Copy(t *testing.T) {
	t.Parallel()

	originalStackErr := &StackErr{
		Punctuation: "foo",
		Separator:   "bar",
		arguments:   []string{"some", "arguments"},
	}

	copyOfStackErr := originalStackErr.Copy("fizz", "buzz")

	assert.Equal(t, "foo", copyOfStackErr.Punctuation)
	assert.Equal(t, "bar", copyOfStackErr.Separator)
	assert.ElementsMatch(t, []string{"fizz", "buzz"}, copyOfStackErr.arguments)
}

func TestStackErr_Punctuation(t *testing.T) {
	t.Parallel()

	stackErr := &StackErr{
		Punctuation: "foo",
	}

	err := stackErr.Handle(errTest, "test")

	assert.Equal(t, "github.com/ovotech/stackerr.TestStackErr_Punctuation.testtest", err.Error())
}

func TestStackErr_Separator(t *testing.T) {
	t.Parallel()

	stackErr := &StackErr{
		Separator: "foo",
	}

	err := stackErr.Handle(errTest, "test")

	assert.Equal(t, "github.com/ovotech/stackerr.TestStackErr_Separator.testfootest", err.Error())
}
