package main

import (
	"fmt"
)

func ex1() {
	// Exercise 1
	//The key here is that you need to use a type conversion to convert `i` to a `float64`
	//so it can be assigned. Go doesn't have automatic type promotion.
	// Notice in our code that you don't need to specify the type for `i` or `f`;
	// it can be inferred from the right hand side. The default type for an integer literal is `int`,
	// so `i`'s type can be inferred.
	// The `float64` type conversion on the right hand side of `f`'s declaration specifies `f`'s type.
	header("exercise 1")

	// var i int = 20
	// var i = 20
	i := 20
	fmt.Println(i)

	var f = float64(i)
	fmt.Println(i, f)

}
