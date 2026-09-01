package main

import (
	"fmt"
)

func ex2() {
	// Exercise 2
	// This question covers type conversions with strings and runes,
	// as well as the zero-based offset for slices.
	// You need to convert the `message` string to a `[]rune` using a type conversion.
	// Then you get the 4th rune in the rune slice using `runes[3]`.
	// Finally, you need to use a `string` type conversion to convert the rune back into a character.
	header("exercise 2")

	// message := "Hi 👩 and 👨"
	// runes := []rune(message)

	var message string = "Hi 👩 and 👨"
	var rm []rune = []rune(message)

	fmt.Println(rm, rm[3], string(rm[3]))
}

// Note
// Parentheses = conversion
// Type Conversion: []type(value)
// []rune("hello")     // convert string → slice of runes
// []byte("hello")     // convert string → slice of bytes
// int(3.14)           // convert float → int
// float64(42)         // convert int → float
// string([]rune{72})  // convert rune slice → string

// Note
// Curly braces = create new slice
// Literal Creation: []type{values}
// []string{"hello", "world"}
// []int{1, 2, 3}
// []rune{'H', 'e', 'l', 'l', 'o'}
// []byte{72, 101, 108}
