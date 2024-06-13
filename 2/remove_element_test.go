package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRemoveElementSuccess(t *testing.T) {
	checkSlice := func(t testing.TB, length, expectedLength int, result, expectedResult []int) {
		t.Helper()
		if !reflect.DeepEqual(length, expectedLength) || !reflect.DeepEqual(result, expectedResult) {
			t.Errorf("got %v expected %v", length, expectedLength)
			t.Errorf("got %v expected %v", result, expectedResult)
		}
	}

	t.Run("remove occurrences of number 3 from slice [1,2,3,3,5] and return both new length of slice and updated slice" , func(t *testing.T) {
		slice := []int{1,2,3,3,7,1}
		val := 3

		length, result := RemoveElement(slice, val)
		expectedLength := 4
		expectedResult := []int{1,2,7,1}
		checkSlice(t, length, expectedLength, result, expectedResult)
	})
	t.Run("return same slice if number given is not present in the original slice" , func(t *testing.T) {
		slice := []int{1,2,3}
		val := 5

		length, result := RemoveElement(slice, val)
		expectedLength := 3
		expectedResult := []int{1,2,3}
		checkSlice(t, length, expectedLength, result, expectedResult)
	})
}

func ExampleRemoveElement() {
	slice := []int{1,2,3,3,7,1}
	val := 7

	length, result := RemoveElement(slice, val)
	fmt.Println(result, length)
	// Output: [1 2 3 3 1] 5
}