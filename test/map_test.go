package test

import (
	// "fmt"
	"testing"
)

func TestTest(t *testing.T) {

	var m map[int]string
	m = map[int]string

	m[1] = "OK"
	t.Log(m[1])
}
