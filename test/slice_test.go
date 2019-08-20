package test

import (
	"testing"
)

func TestSlice(t *testing.T) {

	var s0 []int // 不定长

	append(s0, 1)

	t.Log(len(s0), cap(s0))
}
