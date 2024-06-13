package main

import "fmt"

// LeetCode 1
// Merge nums1 and nums2 into a single array sorted in non-decreasing order, removing zeros.
func Merge(nums1 []int, m int, nums2 []int, n int) []int  {
    for _, v := range nums2 {
        nums1 = append(nums1, v)
    }
    fmt.Println(nums1)
	return nums1
}

func main() {
	Merge([]int{1,2,3,0,0,0}, 3, []int{2,5,6}, 3)
}