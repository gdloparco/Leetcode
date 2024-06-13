package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRemoveDuplicatesSuccess(t *testing.T) {
	checkSlice := func(t testing.TB, length, expectedLength int, slice, expectedSlice []int) {
		t.Helper()
		if !reflect.DeepEqual(length, expectedLength) || !reflect.DeepEqual(slice, expectedSlice) {
			t.Errorf("got %v, expected %v", length, expectedLength)
			t.Errorf("got %v, expected %v", slice, expectedSlice)
		}
	}
	t.Run("given [1 2 3 3 4 5 5] expecting 5 and [1 2 3 4 5]", func(t *testing.T) {
		testerSlice := []int{1,1,2,3,4,4,5}
		length, slice := RemoveDuplicates(testerSlice)
		expectedLength := 5
		expectedSlice := []int{1, 2, 3, 4, 5}
		checkSlice(t, length, expectedLength, slice, expectedSlice)
	})
	t.Run("given [0] expecting 1 and [0]", func(t *testing.T) {
		testerSlice := []int{0}
		length, slice := RemoveDuplicates(testerSlice)
		expectedLength := 1
		expectedSlice := []int{0}
		checkSlice(t, length, expectedLength, slice, expectedSlice)
	})
	t.Run("given empty slice expecting 0 and []", func(t *testing.T) {
		testerSlice := []int{}
		length, slice := RemoveDuplicates(testerSlice)
		expectedLength := 0
		expectedSlice := []int{}
		checkSlice(t, length, expectedLength, slice, expectedSlice)
	})
}

func ExampleRemoveDuplicates() {
	testerSlice := []int{1,1,2,3,3,7,7,9,11,11,11}
	length, slice := RemoveDuplicates(testerSlice)
	fmt.Println(length, slice)
	// Output: 6 [1 2 3 7 9 11]
}