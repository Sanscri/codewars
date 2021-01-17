package main

import "fmt"
/*
Reversed strings
https://www.codewars.com/kata/5168bb5dfe9a00b126000018
Complete the solution so that it reverses the string passed into it.
'world'  =>  'dlrow'
*/
func Solution(word string) string {
	runes := []rune(word)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
func main() {
	fmt.Println(Solution("world"))
}
