package main

// LeetCode 4
// Given an array nums of size n, return the majority element.

// The majority element is the element that appears more than ⌊n / 2⌋ times. 
// You may assume that the majority element always exists in the array.

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numsMap := make(map[int]int)

	for _, v := range nums {
		numsMap[v] ++
	}

	majorityNumber := nums[0]
	maxCount := 0

	for key, value := range numsMap {
		if value > maxCount {
			majorityNumber = key
			maxCount = value
		}
	}

	return majorityNumber
}

// Boyer-Moore Voting Algorithm solution

// func majorityElement(nums []int) int {
//     var majNum, count int
//     for _, num := range nums {
//         if count == 0 {
//             majNum = num
//         }
//         if num == majNum {
//             count++
//         } else {
//             count--
//         }
//     }
//     return majNum
// }