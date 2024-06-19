package main

import (
	"fmt"
	"testing"
)

func TestGetMiddle(t *testing.T) {

	checkEquality := func(t testing.TB, got, expected string) {
		t.Helper()
		if got != expected {
			t.Errorf("got %v but expected %v", got, expected)
		}
	}

	t.Run("test even letter", func(t *testing.T){
		got := getMiddle("find")
		expected := "in"
	
		checkEquality(t, got, expected)
	})
	t.Run("test odd letter", func(t *testing.T){
		got := getMiddle("apple")
		expected := "p"
	
		checkEquality(t, got, expected)
	})
	t.Run("test two letter word", func(t *testing.T){
		got := getMiddle("IF")
		expected := "IF"
	
		checkEquality(t, got, expected)
	})
	t.Run("test empty string", func(t *testing.T){
		got := getMiddle("")
		expected := ""
	
		checkEquality(t, got, expected)
	})
}

func ExamplegetMiddle() {
	got := getMiddle("find")
	fmt.Println(got)
	// Output: in
}