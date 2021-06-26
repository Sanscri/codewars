package main

import "fmt"

/*
Array element parity
https://www.codewars.com/kata/5a092d9e46d843b9db000064
In this Kata, you will be given an array of integers whose elements have both a negative and a positive value,
except for one integer that is either only negative or only positive.
Your task will be to find that integer.

Examples:

[1, -1, 2, -2, 3] => 3

3 has no matching negative appearance
*/
func FindIndex(arr []int, value int) int {
	for i, n := range arr {
		if value == n {
			return i
		}
	}
	return -1
}
func Solve(arr []int) int {
	for _, value := range arr {
		if FindIndex(arr, -value) == -1 {
			return value
		}
	}
	return 0
}

func main() {
	fmt.Println(Solve([]int{1,-1,2,-2,3})) // 3
	fmt.Println(Solve([]int {-3, 1, 2, 3, -1, -4, -2})) // -4
}
