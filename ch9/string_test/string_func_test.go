package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "A, B, C"
	parts := strings.Split(s, ",")
	for _, parts := range parts {
		t.Log(parts)
	}

	t.Log(strings.Join(parts, "-"))

}

func TestStringConvert(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str" + s)

	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}

}
