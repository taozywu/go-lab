package test

import (
	"testing"
)

func TestTest1(t *testing.T) {

	var arr [2]int
	t.Log(arr)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [3]int{1, 2, 3}

	for i := 0; i < 3; i++ {
		t.Log(arr3[i])
	}

	for idx, val := range arr3 {
		t.Log(idx, val)
	}
}

func TestArraySection(t *testing.T) {
	arr3 := [3]int{1, 2, 3}

	arr1 := arr3[:]
	t.Log(arr1)
}
