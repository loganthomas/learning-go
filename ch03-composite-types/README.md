# Chapter 3: Predeclared Types and Declarations

> [!IMPORTANT]
> My personal reminders:
> - Arrays are rarely used directly in Go. 
> - Go considers the _size_ of the array to be part of the _type_ of the array.
>   This makes an array that's declared to be `[3]int` a
>   different type from an array that's declared to be `[4]int`.
>   This also means you cannot use a variable to specify the size of an array, because type must
>   be resolved at compile time, not at runtime.
>   You can't use a type conversion to directly convert arrays of different sizes to identical types.
>   Because you can't convert arrays of different sizes into each other,
>   you can't write a function that works with arrays of any size and
>   you can't assign arrays of different sizes to the same variable.
>   Because of these restrictions, **don't use arrays unless you know the exact length you need ahead of time.**
>   Most of the time, when you want a data structure that holds a sequence of values,
>   a slice is what you should use.
> - Go only has one dimensional arrays, but you can simulated multi-dimensional arrays.
>   Some languages have true matrix support; **Go isn't one of them**.
> - You cannot read or write past the end of an array o ruse a negative index.
>   If you do this with a constant of literal index, it is a compile-time error.
>   An out-of-bounds read or write with a variable index compiles but fails at rune time with a panic.
> - When defining an array and no values are specified, the values will take on the zero value for the declared type.
    When defining a slice and no values are specified (no length by defn), the values will take on the zero value for the slice
    which is `nil`.
> - In Go, `nil` is an identifier that represents the lack of a value for some types. `nil` has no type, so it can be
>   assigned or compared against values of different types. A `nil` slice contains nothing.
> - A slice is not comparable. It is a compile-time error to use `==` to see if two slices are identical
>   or `!=` to see if they are different. The only thing you can compare a slice with using `==` is `nil`.
>   Use `slices.Equal` or `slices.EqualFunc` instead.
> - `append` does not modify the original slice, it returns a new one. **Go is pass-by-value** (give a copy not the original).
    **Python is pass-by-assignment** or pass-by-object-reference (reference to the same object, not a copy).
    Said differently, Go is a _call-by-value_ language. Every time you pass a parameter to a function,
    Go makes a copy of the value that's passed in. Passing a slice to the `append` function actually passes _a copy_ of the slice
    to the function.
> - Just as the built-in `len` function returns the current length of a slice, the built-in `cap` function returns the current
  capacity of a slice. It is used far less frequently than `len`. Most of the time, `cap` is used to check if a slice is large
  enough to hold new data, or if a call to `make` is needed to crate a new slice.
> - While it's nice that slices grow automatically, it's far more efficient to size them once.
>   If you know how many things you plan to put into a slices, create it with the correct initial capacity.
>   You do that with the `make` function.
> - When you take a slice from a slice, you are _not_ making a copy of the data.
>   Instead, you now have two variables that are sharing memory.
>   Changes to an element in a slice affect all slices that share that element.




## Arrays -- Too Rigid to Use Directly
- All elments in the array must be of the type that's specified.
- There are a few different declaration styles:
  ```go
  // Specify the size and the type of the array
  // This creates an array of three ints
  // Since no values are specified,
  // x[0], x[1], and x[2] are initialized as the zero value for int (0)
  var x [3]int
  ```
  ```go
  // If you have initial values of the array,
  // you can specify them with an array literal
  var x [3]int{10, 20, 30}
  ```
  ```go
  // If you have a sparse array (most elements are set to their zero value) 
  // you can specify only the indices with nonzero values in the array literal 
  // This creates the array
  //       [1, 0, 0, 0, 0, 4, 6, 0, 0, 0, 100, 15]
  // index: 0, 1, 2, 3, 4, 5, 6, 7, 8, 9,  10, 11
  var x [12]int{1, 5: 4, 6, 10: 100, 15}
  ```
- When using an array literal to initialize an array, you can replace the number
  that specifies the number of elements in the array with `...`
  ```go
  var x = [...]int{10, 20, 30}
  ```
- You can use the `==` and `!=` to compare two arrays.
  Arrays are equal if they are the same length and contain the same values:
  ```go
  var x [...]int{1, 2, 3}
  var y [3]int{1, 2, 3}
  fmt.Println(x==y) // true
  ```
- Go only has one-dimensional arrays, but you can simulate multidimensional arrays.
  This declares `x` to be an array of length `2` whose type is an array of `ints`
  of length `3`. This sounds pedantic, but some languages have true matrix support;
  **Go isn't one of them**.
  ```go
  var x [2][3]int
  fmt.Println(x) // [[0 0 0] [0 0 0]]
  ```
- Like most languages, arrays in Go are read and written using bracket syntax:
  ```go
	var x = [3]int{1, 2, 3}
	fmt.Println(x) // [1 2 3]
	x[1] = 10
	fmt.Println(x[1], x) // 10 [1 10 3]
  ```
- The reason arrays in Go are rarely used explicitly: Go considers the _size_ of the array
  to be part of the _type_ of the array. This makes an array that's declared to be `[3]int` a
  different type from an array that's declared to be `[4]int`.
- This also means you cannot use a variable to specify the size of an array, because type must
  be resolved at compile time, not at runtime.
- You can't use a type conversion to directly convert arrays of different sizes to identical types.
  Because you can't convert arrays of different sizes into each other,
  you can't write a function that works with arrays of any size and
  you can't assign arrays of different sizes to the same variable.
- Because of these restrictions, don't use arrays unless you know the exact length you need ahead of time.
  For example, some of the cryptographic functions in the standard library return arrays because the size
  of checksums are defined as part of the algorithm. This is the exception, not the rule.

## Slices
- Most of the time, when you want a data structure that holds a sequence of values,
  a slice is what you should use. What makes slices so useful is that you can grow
  slices as needed (the length of the slice is _not_ part of its type).
- Working with slices looks a lot like working with arrays, but subtle differences exist.
  You don't specify the size of the slice when you declare it:
  ```go
  // Careful, [...] makes an array but [] makes a slice
  var x []int{10, 20, 30}
  ```
- Same rules as before apply:
  ```go

  // [1, 0, 0, 0, 0, 4, 6, 0, 0, 0, 100, 15]
  var x = []int{1, 5: 4, 6, 10: 100, 15}

  // Simulate multidimensional slices and make a slice of slices
  var x = [][]int

  var x = []int{1, 2, 3}
	fmt.Println(x) // [1 2 3]
	x[1] = 10
	fmt.Println(x[1], x) // 10 [1 10 3]
  ```
- More differences can be seen when you look at declaring slices without using a literal.
  This declares a slice of `int`s. Since no value is assigned, `x` gets the zero value for a slice, which is `nil`.
  In Go, `nil` is an identifier that represents the lack of a value for some types. `nil` has no type, so it can be
  assigned or compared against values of different types. A `nil` slice contains nothing.
  ```go
  var x[]int
  fmt.Println(x) // []
  ```
- A slice is not comparable. It is a compile-time error to use `==` to see if two slices are identical
  or `!=` to see if they are different. The only thing you can compare a slice with using `==` is `nil`:
  ```go
  fmt.Println(x == nil) // true
  ```
- Use the `slices` package in the standard library to compare slices.
  - `slices.Equal` takes two slices and returns `true` if the slices are the same length,
     and all the elements are equal. It requires the elements of the slices to be comparable
  - `slices.EqualFunc` lets you pass in a function to determine equality and does not require the
    slice elements to be comparable.
  ```go
  x := []int{1,2,3,4,5}
  y := []int{1,2,3,4,5}
  z := []int{1,2,3,4,5,6}
  s := []int{"a", "b", "c"}
  
  fmt.Println(slices.Equal(x, y)) // true
  fmt.Println(slices.Equal(x, z)) // false
  fmt.Println(slices.Equal(x, s)) // does not compile
  ```
  ```go
  // slice.EqualFunc checks length first
  // The function receives individual elements, not slices (func is called once per element, not the whole slice)
  n := 3
  out := slices.EqualFunc(x[:n], z[:n], func(a, b int) bool {
      return a == b
  })
  fmt.Println(out) // true
  ```
- Before the inclusion of `slices.Equal` and `slices.EqualFunc`, `reflect.DeepEqual` was often used
  to compare slices. Don't use it in new code, as it is slower and less safe than using the functions in the `slices` package.

### len
- Passing a `nil` slice to `len` returns `0`

### append
- This built-in function is used to grow slices
  ```go
  var x[]int
  x = append(x, 10) // assign result to the variable that's passed in
  ```
  ```go
  // append does not modify the original slice, it returns a new one
  // Go is pass-by-value (give a copy not the original)
  // Python is pass-by-assignment or pass-by-object-reference (reference to the same object, not a copy)
  var x []int
	var y []int
	y = append(x, 10)
	fmt.Println(x, y) // [] [10]
  ```
- You can append more than one value at a time:
  ```go
	var x []int
	x = append(x, 10)
	fmt.Println(x) // [10]
	x = append(x, 5, 6, 7)
	fmt.Println(x) // [10 5 6 7]
  ```
- One slice is appended onto another by using the `...` operator to expand the source slice into individual values.
  This looks like Pythons unpacking (`*`)
  ```go
  y := []int{20, 30, 40}
  x = append(x, y...) // [10 5 6 7 20 30 40]
  ```
- It's a compile-time error if you forget to assign the value returned from `append`.
  Remember, Go is a _call-by-value_ language. Every time you pass a parameter to a function,
  Go makes a copy of the value that's passed in. Passing a slice to the `append` function actually passes _a copy_ of the slice
  to the function. The function adds the values to the copy of the slice and returns the copy. You then assign the returned
  copied slice with appended values back to the variable in the calling function.

### Capacity
- Each element in a slice is assigned to consecutive memory locations, which makes it quick to read or write these values.
  The length of the slice is the number of consecutive memory locations that have been assigned a value.
  Each slice also has a _capacity_, which is the number of consecutive memory locations reserved.
  - This can be larger than the length. Each time you append to a slice, one or more values is added to the end of the sliced.
    Each value added increases the length by one. When the length reaches the capacity, there's no more room to put values.
    If you try to add additional values when the length equals the capacity,
    the `append` function uses the Go runtime to allocate a _new_ backing array for the slice
    with a larger capacity. The values in the original backing array are copied to the new one,
    the new values are added to the end of the new backing array, and the slice is updated to refer to the new backing array.
    Finally, the updated slice is returned.
- When a slice grows via `append`, it takes time for the Go runtime to allocate new memory and copy the existing data from the old
  memory to the new. The old memory also needs to be garbage collected. For this reason, the Go runtime usually increases a slice
  by more than one each time it runs out of capacity (see pg 42 for logic).
- Just as the built-in `len` function returns the current length of a slice, the built-in `cap` function returns the current
  capacity of a slice. It is used far less frequently than `len`. Most of the time, `cap` is used to check if a slice is large
  enough to hold new data, or if a call to `make` is needed to crate a new slice.

  ```go
  var x[]int
  fmt.Println(x, len(x), cap(x))

  x = append(x, 10)
  fmt.Println(x, len(x), cap(x))
  
  x = append(x, 20)
  fmt.Println(x, len(x), cap(x))
  
  x = append(x, 30)
  fmt.Println(x, len(x), cap(x))
  
  x = append(x, 40)
  fmt.Println(x, len(x), cap(x))

  x = append(x, 50)
  fmt.Println(x, len(x), cap(x))

  // [] 0 0
  // [10] 1 1
  // [10 20] 2 2
  // [10 20 30] 3 4
  // [10 20 30 40] 4 4
  // [10 20 30 40 50] 5 8
  ```
> [!NOTE]
> While it's nice that slices grow automatically, it's far more efficient to size them once.
> If you know how many things you plan to put into a slices, create it with the correct initial capacity.
> You do that with the `make` function.
>
> You may ask, "If I know the length, why use a slice with capacity and not just an array?".
> Remember, the array size is part of the type and that arrays are copied (inefficient) while slices
> point to underlying data. Usually never use an array in go unless cryptographic hashes, truly fixed size data
> like RGB as `[3]byte`, or backing storage that you immediately slice.

### make
- So far seen slice literal (`var x []int{1,2,3}`) and the `nil` zero value (`var x []int`), but neither of these
  allows for creating an empty slice that already has a length or capacity specified.
- `make` allows you to specify the type, length, and optionally, the capacity.

  ```go
  // This creates an int slice with a length of 5 and a capacity of 5
  // Since it has a length of 5, x[0] through x[4] are valid elements and they all initialize to 0
  x := make([]int, 5)
  fmt.Println(x, len(x), cap(x)) // [0 0 0 0 0] 5 5

  x := make([]int, 5, 10)
	fmt.Println(x, len(x), cap(x)) // [0 0 0 0 0] 5 10

  x := make([]int, 0, 10)
	fmt.Println(x, len(x), cap(x)) // [] 0 10
	x = append(x, 5, 6, 7, 8)
	fmt.Println(x, len(x), cap(x)) // [5 6 7 8] 4 10
  ```

### Emptying a Slice
- `clear` takes in a slice and sets all of the slice's elements to their zero values.
  The length of the slice remains unchanged.
  ```go
	x := []int{1, 2, 3}
	fmt.Println(x, len(x), cap(x)) // [1 2 3] 3 3
	clear(x)
	fmt.Println(x, len(x), cap(x)) // [0 0 0] 3 3

	s := []string{"first", "second", "third"}
	fmt.Println(s, len(s), cap(s)) // [first second third] 3 3
	clear(s)
	fmt.Println(s, len(s), cap(s)) // [    ] 3 3
  ```
### Declaring Your Slice
- Which slice declaration style to use? The primary goal is to minimize the number of times the slice needs to grow.
  If it's possible the slice won't need to grow at all, use `var` declaration with no assigned value to create a `nil` slice.
  ```go
  var data []int
  ```
- If you have some starting values, or if a slice's values aren't going to change, then a slice literal is a good choice.
  ```go
  data := []int{2, 4, 6, 8}
  ```
- If you have a good idea of how large your slice needs to be, but don't know what those values will be when you are writing
  the program, use `make`. See pg 45 for further discussion on whether you should specify a nonzero length or a zero length and 
  a nonzero capacity.

### Slicing Slices
- A _slice expression_ crates a slice from a slice.
- Same as Python ("up to, not including"; if you leave off the starting offset, `0` is assumed)
- Different than Python: no negative indices
  ```go
	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("d:", d)
	fmt.Println("e:", e)

  // x: [a b c d]
  // y: [a b]
  // z: [b c d]
  // d: [b c]
  // e: [a b c d]
  ```
> [!IMPORTANT]
> When you take a slice from a slice, you are _not_ making a copy of the data.
> Instead, you now have two variables that are sharing memory.
> Changes to an element in a slice affect all slices that share that element.
  ```go
	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	z := x[1:]
	x[1] = "y"
	y[0] = "x"
	z[1] = "z"
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

  // x: [x y z d]
  // y: [x y]
  // z: [y z d]
  ```

- Slicing slices gets extra confusing when combined with `append`.
  Whenever you take a slice from another slice, the subslice's capacity is set to the capacity of the original slice,
  minus the starting offset of the subslice within the original slice. This means elements of the original slice beyond
  the end of the subslice, including unused capacity, are shared by both sides
  ```go
	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	fmt.Println(cap(x), cap(y))
	y = append(y, "z")
	fmt.Println("x:", x)
	fmt.Println("y:", y)

  // 4 4
  // x: [a b z d]
  // y: [a b z]
  ```
  ```go
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")
	y := x[:2]
	z := x[2:]
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println(cap(x), cap(y), cap(z))
	fmt.Println(strings.Repeat("-", 10))
  // x: [a b c d]
  // y: [a b]
  // z: [c d]
  // 5 5 3
  // ----------
	
  y = append(y, "i", "j", "k")
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println(cap(x), cap(y), cap(z))
	fmt.Println(strings.Repeat("-", 10))
  // x: [a b i j]
  // y: [a b i j k]
  // z: [i j]
  // 5 5 3
  // ----------
	
  x = append(x, "x")
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println(cap(x), cap(y), cap(z))
	fmt.Println(strings.Repeat("-", 10))
  // x: [a b i j x]
  // y: [a b i j x]
  // z: [i j]
  // 5 5 3
  // ----------
	
  z = append(z, "y")
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println(cap(x), cap(y), cap(z))
	fmt.Println(strings.Repeat("-", 10))
  // x: [a b i j y]
  // y: [a b i j y]
  // z: [i j y]
  // 5 5 3
  // ----------
  ```
- To avoid complicated slice situations, you should either never use `append` with a subslice,
  or make sure that `append` doesn't cause an overwrite by using a _full slice exprssion_.
  - The full slice expression includes a third part, which indicates the last position in the parent slice's capacity
    that's available for the subslice.
  - Subtract the starting offset from this number to get the subslice's capacity.
  - Both `y` and `z` have a capacity of `2`. Because you limited the capacity of the subslides to their lengths,
    appending additional elements onto `y` and `z` crated _new slices_ that didn't interact with the other slices.
  - In the previous example, both `y` and `z` had more capacity than their length, so didn't need a new slice when appending,
    and thus, interacted with same memory and changes other slices.
  ```go
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")
	y := x[:2:2]
	z := x[2:4:4]
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println(cap(x), cap(y), cap(z))
	fmt.Println(strings.Repeat("-", 10))
  // x: [a b c d]
  // y: [a b]
  // z: [c d]
  // 5 2 2
  // ----------

	y = append(y, "i", "j", "k")
	x = append(x, "x")
	z = append(z, "y")
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
  // x: [a b c d x]
  // y: [a b i j k]
  // z: [c d y]
  ```

### copy
- Use the built-in `copy` function to crate a slice that's independent of the original
  - The function takes two parameters: The first is the _destination_ slice,
    and the second is the _source_ slice.
  - The function copies as many values as it can from source to destination, limited by whichever slice is smaller
  - The function returns the number of elements copied
  - The _capacity_ of `x` and `y` doesn't matter; it's the length that's important
  ```go
	x := []int{1, 2, 3, 4}
	y := make([]int, 4)
	num := copy(y, x)
	fmt.Printf("copied %05d or %05d elements to y \n", num, len(x))
	fmt.Println("copied", num, "elements...")
	fmt.Println(y, num)

  // copied 00004 or 00004 elements to y 
  // copied 4 elements...
  // [1 2 3 4] 4
  ```
- You can copy a subset of a slice. The following copies the first two elements of a four-element slice
  into a two-element slice:
  ```go
	x := []int{1, 2, 3, 4}
	y := make([]int, 2)
	num := copy(y, x)
	fmt.Printf("copied %d or %d elements to y \n", num, len(x))
	fmt.Println(y, num)

  // copied 2 or 4 elements to 6
  // [1 2] 2
  ```
- You can also copy from the middle of the source slice:
  ```go
  x := []int{1, 2, 3, 4}
  y := make([]int, 2)
  num := copy(y, x[2:])

  // copied 2 or 4 elements to 6
  // [3 4] 2
  ```
### Converting Arrays to Slices
- If you have an array, you can take a slice from it using a slice expression.
  This is a useful way to bridge an array to a function that only takes slices.
  - Be aware that taking a slice from an array has the same memory-sharing properties as taking a slice from a slice.
  ```go
  xArray := [4]int{1, 2, 3, 4}
  xSlice := xArray[:]

  // Can convert a subslice of an array into a slice
  x := [4]int{5, 6, 7, 8}
  y := x[:2]
  z := x[2:]

  // Same memory sharing properties as slice from slice
  x := [4]int{5, 6, 7, 8}
	y := x[:2]
	z := x[2:]
	x[0] = 10
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

  // x: [10 6 7 8]
  // y: [10 6]
  // z: [7 8]
  ```
### Converting Slices to Arrays
- Use a type conversion to make an array variable from a slice.
  You can convert an entire slice to an array of the same type, or can create an array from a subset of the slice.
- When you convert a slice to an array, the data in the slice is copied to new memory.
  That means that changes to the slice won't affect the array, and vice versa.
- The size of the array must be specified at compile time.
  It is a compile-time error to use `[...]` in a slice to array type conversion.
  ```go
  xSlice := []int{1, 2, 3, 4}
  xArray := [4]int(xSlice)
  smallArray := [2]int(xSlice)
  xSlice[0] = 10
  fmt.Println(xSlice)
  fmt.Println(xArray)
  fmt.Println(smallArray)

  // [10 2 3 4]
  // [1 2 3 4]
  // [1 2]
  ```
- While the size of the array can be smaller than the size of the slice, it cannot be bigger.
  ```go
  panicArray := [5]int(xSlice)
  fmt.Println(panicArray)

  // panic: runtime error: cannot convert slice with length 4 to array or pointer to array with length 5
  ```
- You can also use a type conversion to convert a slice into a pointer to an array:
  ```go
  xSlice := []int{1, 2, 3, 4}
  xArrayPointer := (*[4]int)(xSlice)
  ```
- After converting a slice into an array pointer, the storage between the two is shared.
  A change to one will change the other:
  ```go
  xSlice[0] = 10
  xArrayPointer[1] = 20
  fmt.Println(xSlice) // [10 20 3 4]
  fmt.Println(xArrayPointer) // &[10 20 3 4]
  ```
## Strings and Runes and Bytes
- Under the hood, Go uses a sequence of bytes to represent a string.
  These bytes don't have to be in any particular character encoding, but several Go library functions assume that a string is
  composed of a sequence of UTF-8 code points.
- Just as you can extract a single value from an array or slice, you can extract a single value from a string using an _index
  expression_
  ```go
  var s string = "Hello there"
  var b byte = s[6] 
	fmt.Println(b) // 116 is the UTF-8 value of lowercase t
	
  var s2 string = s[4:7]
	var s3 string = s[:5]
	var s4 string = s[6:]
	fmt.Println(s2) // o t
	fmt.Println(s3) // Hello
	fmt.Println(s4) // there
  ```
- Since strings are immutable, they don't have the modification problems that slices do. They have a different problem.
  A string is composed of a sequence of bytes, while a code point in UTF-8 can be anywhere from one to four bytes long.
  When dealing with languages other than English or with emojis, you run into code points that are multiple bytes long in UTF-8.
  - In this example, `s3` will still equal `"Hello"`. The variable `s4` is set to the sun emoji. But `s2` is not set to `"o ☀️"`.
    Instead, you get `"o �"` because you copied only the first byte of the sun emoji's code point, which is not a valid code point
    on its own.
  ```go
	var s string = "Hello ☀️"
	var s2 string = s[4:7]
	var s3 string = s[:5]
	var s4 string = s[6:]
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

  // o �
  // Hello
  // ☀️

  fmt.Println(len(s3)) // 5
  fmt.Println(len(s4)) // 6
  fmt.Println(len(s)) // 12 (5 + space (1) + 6)
  ```
- Because of this complicated relationship between runes, strings, and bytes,
  Go has some interesting type conversions between these types.
- A single rune or byte can be converted to a string
  ```go
  var a rune = 'x'
  var s string = string(a)
  var b byte = 'y'
  var s2 string = string(b)

  fmt.Println(a, s, b, s2) // 120 x 121 y
  ```
- A common bug for new Go developers is to try to make an `int` into a `string` by using a type conversion:
  ```go
  var x int = 65
  var y string = string(x)
  fmt.Println(x, y) // 65 A

  // Correct way
  fmt.Println(strconv.Itoa(x)) // 65
  y := fmt.Sprintf("%d", x)
  fmt.Println(x, y) // 65 65
  ```
- A string can be converted back and forth to a slice of bytes or a slice of runes
  ```go
	var s string = "Hello, 🌞"
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	fmt.Println(bs)   // [72 101 108 108 111 44 32 240 159 140 158] this is UTF-8 bytes (most common)
	fmt.Println(rs)   // [72 101 108 108 111 44 32 127774] this is runes

  fmt.Println(string(bs)) // Hello, 🌞
	fmt.Println(string(rs)) // Hello, 🌞
  ```
> [!NOTE]
> Rather than use the slice and index expressions with strings, you should extract substrings and code points from strings
> using the functions in the `strings` and `unicode`/`uft8` packages in the standard library.  

## Maps
### Reading and Writing a Map
### The comma ok Idiom
### Deleting from Maps
### Emptying a Map
### Comparing a Map
### Using Maps
### Using Maps as Sets
## Structs
### Anonymous Structs
### Comparing and Converting Structs

