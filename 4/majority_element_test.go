package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMajorityElementSuccess(t *testing.T) {
	checkSlice := func(t testing.TB, got, expected int) {
		t.Helper()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v, expected %v", got, expected)
		}
	}
	t.Run("given [1 1 2 3 4 4 4 5] expecting 4", func(t *testing.T) {
		testerSlice := []int{1,1,2,3,4,4,4,5}
		got := majorityElement(testerSlice)
		expected := 4
		checkSlice(t, got, expected)
	})
	t.Run("given [0] expecting 0", func(t *testing.T) {
		testerSlice := []int{0}
		got := majorityElement(testerSlice)
		expected := 0
		checkSlice(t, got, expected)
	})
	t.Run("given empty slice expecting 0", func(t *testing.T) {
		testerSlice := []int{}
		got := majorityElement(testerSlice)
		expected := 0
		checkSlice(t, got, expected)
	})
}

func ExamplemajorityElement() {
	testerSlice := []int{1,1,2,3,3,7,7,9,11,11,11}
	result := majorityElement(testerSlice)
	fmt.Println(result)
	// Output: 11
}