package main

import "fmt"

/*
Power of 3
https://www.codewars.com/kata/57be674b93687de78c0001d9/go
Given a positive integer N, return the largest integer k such that 3^k < N.

For example,

LargestPower(3) // returns 0
LargestPower(4) // returns 1
You may assume that the input to your function is always a positive integer.
*/
func LargestPower(n uint64) int {
	var k = 0
	var num = uint64(1)
	for num < n {
		k += 1
		num *= uint64(3)
	}
	k -= 1
	return k // return the largest integer k such that 3^k < n
}

func main() {
	fmt.Println(LargestPower(3)) // returns 0
	fmt.Println(LargestPower(4))  // returns 1
}
