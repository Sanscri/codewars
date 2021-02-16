package main

import (
	"fmt"
	"strconv"
)

/*
Error correction #1 - Hamming Code
https://www.codewars.com/kata/5ef9ca8b76be6d001d5e1c3e

Background information
The Hamming Code is used to correct errors, so-called bit flips, in data transmissions.
Later in the description follows a detailed explanation of how it works.
In this Kata we will implement the Hamming Code with bit length 3; this has some advantages and disadvantages:

[ + ] It's simple to implement
[ + ] Compared to other versions of hamming code, we can correct more mistakes
[ - ] The size of the input triples
Task 1: Encode function
Implement the encode function, using the following steps:

convert every letter of the text to its ASCII value;
convert ASCII values to 8-bit binary;
triple every bit;
concatenate the result;
For example:

input: "hey"
--> 104, 101, 121                  // ASCII values
--> 01101000, 01100101, 01111001   // binary
--> 000111111000111000000000 000111111000000111000111 000111111111111000000111  // tripled
--> "000111111000111000000000000111111000000111000111000111111111111000000111"  // concatenated
Task 2: Decode function:
Check if any errors happened and correct them. Errors will be only bit flips, and not a loss of bits:

111 --> 101 : this can and will happen
111 --> 11 : this cannot happen
Note: the length of the input string is also always divisible by 24 so that you can convert it to an ASCII value.

Steps:

Split the input into groups of three characters;
Check if an error occurred: replace each group with the character that occurs most often, e.g. 010 --> 0, 110 --> 1, etc;
Take each group of 8 characters and convert that binary number;
Convert the binary values to decimal (ASCII);
Convert the ASCII values to characters and concatenate the result
For example:

input: "100111111000111001000010000111111000000111001111000111110110111000010111"
--> 100, 111, 111, 000, 111, 001, ...  // triples
-->  0,   1,   1,   0,   1,   0,  ...  // corrected bits
--> 01101000, 01100101, 01111001       // bytes
--> 104, 101, 121                      // ASCII values
--> "hey"
*/

func Encode(text string) string {
	var asciiValues = []byte(text)
	var bits = ""
	for i := 0; i < len(asciiValues); i++ {
		for j := 7; j >= 0; j-- {
			mask := byte(1 << uint8(j))
			var currentValue = (asciiValues[i] & mask) >> uint8(j)
			if currentValue == byte(0) {
				bits += "000"
			} else {
				bits += "111"
			}
		}
	}
	return bits
}

func Decode(bits string) string {
	var text = ""
	var vector []byte
	for i := 0; i < len(bits); i += 3 {
		var tmpValue = bits[i:i+3]
		var sum = 0
		for j := 0; j < 3; j++ {
			value, _ := strconv.Atoi(string(tmpValue[j]))
			sum += value
		}
		if sum < 2 {
			vector = append(vector, 0)
		} else {
			vector = append(vector, 1)
		}
	}
	var bytes []byte
	for i := 0; i < len(vector); i += 8 {
		var bitsArray = vector[i:i+8]
		var value byte = 0
		for j := 7; j >= 0; j-- {
			value = value | (bitsArray[7 - j] << uint8(j))
		}
		bytes = append(bytes, value)
	}
	text = string(bytes)
	return text
}

func main() {
	var tmp = Encode("hey")
	fmt.Println(tmp)
	var text = Decode("100111111000111001000010000111111000000111001111000111110110111000010111")
	fmt.Println(text)
}
