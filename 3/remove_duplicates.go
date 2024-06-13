package main

import "fmt"

// LeetCode 3
// Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place
// such that each unique element appears only once. The relative order of the elements should
// be kept the same. Then return the number of unique elements in nums.

// Consider the number of unique elements of nums to be k, to get accepted, you need to do the
// following things:

// Change the array nums such that the first k elements of nums contain the unique elements
// in the order they were present in nums initially. The remaining elements of nums are not
// important as well as the size of nums.
// Return k.

func RemoveDuplicates(slice []int) (int, []int) {

    if len(slice) == 0 {
        return 0, slice
    }
    
    newIndex := 1
    for i := 1; i < len(slice); i++ {
        if slice[i] != slice[i-1] {
            slice[newIndex] = slice[i]
            newIndex++
        }
    }
    
    return newIndex, slice[:newIndex]
}

func main() {
	length, slice := RemoveDuplicates([]int{1,1,1,1,1,1,1})
    fmt.Println(length, slice)
    fmt.Println("Expecting 1 and [1] above")

    length2, slice2 := RemoveDuplicates([]int{1,2,2,3,5,5,5})
    fmt.Println(length2, slice2)
    fmt.Println("Expecting 4 and [1 2 3 5] above")
}