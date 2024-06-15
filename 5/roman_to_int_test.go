package main

import (
	"testing"
	"reflect"
)

func checkConversion(t *testing.T, got int, expected int) {
	t.Helper()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v expected %v", got, expected)
	}
}

func TestRomanToIntSuccess(t *testing.T) {

	t.Run("LXXII to 72", func(t *testing.T) {
		got := romanToInt("LXXII")
		expected := 72
		checkConversion(t, got, expected)
	})
	t.Run("MCMXCIX to 1999 - subtraction cases", func(t *testing.T) {
		got := romanToInt("MCMXCIX")
		expected := 1999
		checkConversion(t, got, expected)
	})
	t.Run("DCCXXXIII to 733 - two and three in a row", func(t *testing.T) {
		got := romanToInt("DCCXXXIII")
		expected := 733
		checkConversion(t, got, expected)
	})
}