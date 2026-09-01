package main

import (
	"fmt"
)

func ex1() {
	// Exercise 1
	// This problem tests creating a slice literal and calculating slice offsets.
	// The first value in a slice expression is the zero-based starting position.
	// It's optional if you are starting at the beginning of the slice.
	// The second value is one greater than the last position you want to include.
	// The second value is optional if you want to go all the way to the end of the slice.
	header("exercise 1")

	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	s1 := greetings[:2]
	s2 := greetings[1:4]
	s3 := greetings[3:]
	fmt.Println(greetings)
	fmt.Println(s1, s2, s3)
}
