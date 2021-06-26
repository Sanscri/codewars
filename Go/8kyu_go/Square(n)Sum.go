package main

import "fmt"

/*
Square(n) Sum
https://www.codewars.com/kata/515e271a311df0350d00000f
Complete the square sum function so that it squares each number passed into it and then sums the results together.
For example, for [1, 2, 2] it should return 9 because 1^2 + 2^2 + 2^2 = 9.
*/
func SquareSum(numbers []int) int {
	var answer = 0
	for _, value := range numbers{
		answer += value * value
	}
	return answer
}

func main() {
	fmt.Println(SquareSum([]int{1,2})) //5
	fmt.Println(SquareSum([]int{0,3,4,5})) //50
}
