package util

import (
	"testing"
)

func TestLineNumbers(t *testing.T) {
	Equal(t, "foo", "foo", "msg!")
}

func TestNotEqual(t *testing.T) {
	NotEqual(t, "foo", "bar", "msg!")
}
