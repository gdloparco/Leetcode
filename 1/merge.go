package main

import (
	"sort"
)

// LeetCode 1
// Merge nums1 and nums2 into a single array sorted in non-decreasing order.

// The final sorted array should not be returned by the function, but instead be stored inside the array nums1. 
// To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should
// be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

func Merge(nums1 []int, m int, nums2 []int, n int) []int  {
    if n != 0 {
        for i, v := range nums2 {
            nums1[m+i] = v
    }
    sort.Ints(nums1)
    }
    return nums1
}

func main() {
	Merge([]int{1,2,3,0,0,0}, 3, []int{2,5,6}, 3)
}