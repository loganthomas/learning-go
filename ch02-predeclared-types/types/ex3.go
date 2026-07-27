package main

import (
	"fmt"
	// "math"
)

func ex3() {
	// Exercise 3
	// The maximum values for each type is specified in the table in Chapter 2.
	// If you are industrious, you can also find the constants declared in the standard library in the `math` package.
	// fmt.Println(math.MaxUint8)  // 255
	// fmt.Println(math.MaxInt32)  // 2147483647
	// There is no math.MaxUint64 constant because it would overflow int on some architectures. Instead you can use:
	// var maxUint64 uint64 = ^uint64(0)  // 18446744073709551615
	// Adding 1 to each variable causes an overflow, not an error. You get the minimum value for each one.
	header("exercise 3")

	// var b byte = 255
	// var smallI int32 = 2147483647
	// var bigI uint64 = 18446744073709551615
	var (
		b      byte   = 255 // 8 bits in 1 byte; same as uint8 (2^8 possible values; 0 to 255)
		smallI int32  = 2147483647
		bigI   uint64 = 18446744073709551615
	)
	fmt.Println(b, smallI, bigI)

	// b = b + 1
	// smallI = smallI + 1
	// bigI = bigI + 1
	b += 1
	smallI += 1
	bigI += 1
	fmt.Println(b, smallI, bigI)

	// Appendix
	// fmt.Println(math.MaxUint8)        // 255
	// fmt.Println(math.MaxInt32)        // 2147483647
	// var maxUint64 uint64 = ^uint64(0) // 18446744073709551615
	// fmt.Println(maxUint64)
}
