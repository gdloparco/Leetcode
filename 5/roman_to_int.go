package main

// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000

// Roman numerals are usually written largest to smallest from left to right. H

// I can be placed before V (5) and X (10) to make 4 and 9.
// X can be placed before L (50) and C (100) to make 40 and 90.
// C can be placed before D (500) and M (1000) to make 400 and 900.

// Given a roman numeral, convert it to an integer.

func romanToInt(romanNum string) int {
	convTable := map[rune]int{'I':1, 'V':5, 'X':10, 'L':50, 'C':100, 'D':500, 'M':1000}
	result := 0

	for i, v := range romanNum {
		value := convTable[v]
		if i < len(romanNum)-1 && value < convTable[rune(romanNum[i+1])] {
			result -= value
		} else {
			result += value
		}
	}
	return result
}

// BEFORE REFACTORING

// func romanToInt(romanNum string) int {
// 	convTable := map[rune]int{'I':1, 'V':5, 'X':10, 'L':50, 'C':100, 'D':500, 'M':1000}
// 	integers := []int{}
// 	result := 0

// 	for _, v := range romanNum {
// 		integers = append(integers, convTable[v])
// 	}
// 	fmt.Println(integers)
// 	for i, v := range integers {
// 		if i == 0 {
// 			if v == integers[i+1] || v > integers[i+1]{
// 				result += v
// 				fmt.Println(result)
// 			} else if v < integers[i+1]{
// 				result += (integers[i+1] - v)
// 				fmt.Println(result)
// 			} 
// 		} else if i+1 < len(integers) {
// 			if v > integers[i-1] {
// 				continue
// 			} else if v == integers[i-1] {
// 				result += v
// 			} else if v == integers[i+1]{
// 				result += v
// 				fmt.Println(result)
// 			} else if v < integers[i+1]{
// 				result += (integers[i+1] - v)
// 				fmt.Println(result)
// 			} 
// 		} else {
// 			if v > integers[i-1] {
// 				continue
// 			} else {
// 				result += v
// 				fmt.Println(result)
// 			} 
// 		}
// 	}
// 	return result
// }