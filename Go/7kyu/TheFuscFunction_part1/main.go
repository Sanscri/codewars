package main
/*
The fusc function -- Part 1

https://www.codewars.com/kata/570409d3d80ec699af001bf9
The fusc function is defined recursively as follows:

1. fusc(0) = 0
2. fusc(1) = 1
3. fusc(2 * n) = fusc(n)
4. fusc(2 * n + 1) = fusc(n) + fusc(n + 1)
The 4 rules above are sufficient to determine the value of fusc for any non-negative input n.
For example, let's say you want to compute fusc(10).

fusc(10) = fusc(5), by rule 3.
fusc(5) = fusc(2) + fusc(3), by rule 4.
fusc(2) = fusc(1), by rule 3.
fusc(1) = 1, by rule 2.
fusc(3) = fusc(1) + fusc(2) by rule 4.
fusc(1) and fusc(2) have already been computed are both equal to 1.
Putting these results together fusc(10) = fusc(5) = fusc(2) + fusc(3) = 1 + 2 = 3

Your job is to produce the code for the fusc function.
In this kata, your function will be tested with small values of n,
so you should not need to be concerned about stack overflow or timeouts.

Hint: Use recursion.
*/

import "fmt"
func Fusc(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n % 2 == 0 {
		return Fusc(n / 2)
	} else {
		return Fusc(n / 2) + Fusc(n / 2 + 1)
	}
}
func main() {
	fmt.Println(Fusc(0)) // 0
	fmt.Println(Fusc(1)) // 1
	fmt.Println(Fusc(10)) // 3
}
