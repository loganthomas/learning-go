package main

import "fmt"

func ex2() {
	// Exercise 2
	// To solve this, make `value` an untyped constant that's assigned an integer literal value.
	// You can then assign it to both an `int` and a `float64`.
	// Since the value of `value` is an integer literal (which has a default type of `int`),
	// you do not need to specify the type for `i` and can use the short declaration format.
	// However, to make `f` a `float64`, you need to specify its type.
	header("exercise 2")

	const value = 42

	// var i int = value
	// var i = value
	i := value
	var f float64 = value
	fmt.Println(i, f)
}
