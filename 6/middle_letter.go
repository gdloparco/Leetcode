// Your job is to return the middle letter of a word. If the word's length is odd, return the middle letter.
// If the word's length is even, return the middle 2 letters.

// Acceptance Criteria
// get_middle("test") # => "es"
// get_middle("testing") # => "t"
// get_middle("middle") # => "dd"
// get_middle("A") # => "A"
// get_middle("of") # => "of"

package main


func getMiddle(s string) string {
	middle := len(s)/2
	odd := len(s) % 2 == 1

		if odd && len(s) != 0 {
			return string(s[middle])
		} else if !odd && len(s) != 0 {
			return string(s[middle -1]) + string(s[middle])
		} else {
			return ""
		}
}