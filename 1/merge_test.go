package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMergeSuccess(t *testing.T) {
	checkSlice := func(t testing.TB, got, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v expected %v", got, expected)
		}
	}

	t.Run("merge two slices [1,2,3,0,0,0] and [2,5,6]" , func(t *testing.T) {
		slice1 := []int{1,2,3,0,0,0}
		m := 6
		slice2 := []int{2,5,6}
		n := 3

		got := Merge(slice1, m, slice2, n)
		expected := []int{1,2,2,3,5,6}
		checkSlice(t, got, expected)
	})
}

func ExampleMerge() {
	slice1 := []int{1,2,3,0,0,0}
	m := 6
	slice2 := []int{2,5,6}
	n := 3
	result := Merge(slice1, m, slice2, n)
	fmt.Println(result)
	// Output: [1,2,2,3,5,6]
}