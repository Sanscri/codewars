package main

import "fmt"

/*
Maximum Multiple
https://www.codewars.com/kata/5aba780a6a176b029800041c
Given a Divisor and a Bound, Find the largest integer N, Such That,
Conditions :
N is divisible by divisor
N is less than or equal to bound
N is greater than 0.
*/
func MaxMultiple(d, b int) int {
	var n = b
	for n <= b {
		if n % d != 0 {
			n -= 1
		} else{
			break
		}
	}
	return n
}
func main() {
	fmt.Println(MaxMultiple(2, 7)) //6
}