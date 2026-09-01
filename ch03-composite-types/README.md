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
>   When defining a slice and no values are specified (no length by defn), the values will take on the zero value for the slice
>   which is `nil`.
> - In Go, `nil` is an identifier that represents the lack of a value for some types. `nil` has no type, so it can be
>   assigned or compared against values of different types. A `nil` slice contains nothing.
> - A slice is not comparable. It is a compile-time error to use `==` to see if two slices are identical
>   or `!=` to see if they are different. The only thing you can compare a slice with using `==` is `nil`.
>   Use `slices.Equal` or `slices.EqualFunc` instead.
> - `append` does not modify the original slice, it returns a new one. **Go is pass-by-value** (give a copy not the original).
>   **Python is pass-by-assignment** or pass-by-object-reference (reference to the same object, not a copy).
>   Said differently, Go is a _call-by-value_ language. Every time you pass a parameter to a function,
>   Go makes a copy of the value that's passed in. Passing a slice to the `append` function actually passes _a copy_ of the slice
>   to the function.
> - Just as the built-in `len` function returns the current length of a slice, the built-in `cap` function returns the current
>   capacity of a slice. It is used far less frequently than `len`. Most of the time, `cap` is used to check if a slice is large
>   enough to hold new data, or if a call to `make` is needed to crate a new slice.
> - While it's nice that slices grow automatically, it's far more efficient to size them once.
>   If you know how many things you plan to put into a slices, create it with the correct initial capacity.
>   You do that with the `make` function.
> - When you take a slice from a slice, you are _not_ making a copy of the data.
>   Instead, you now have two variables that are sharing memory.
>   Changes to an element in a slice affect all slices that share that element.
> - When should you use a map, and when should you use a slice? You should use slices for lists of data when data should
>   be processed sequentially or the order of the elements is important. Maps are useful when you need to organize values
>   using something other than an increasing integer value, such as a name.
> - In Go, **order of map is random**. Go intentionally randomizes map iteration to prevent developers from relying on it.
>   You have to use a slice of keys alongside the map for sorted order.
> - Go provides the _comma ok idiom_ to tell the difference between a key that's associated with a zero value
>   and a key that's not in the map. Go determines this at compile time based on the assignment context;
>   specifically, how many variables are on the left-hand side.
>   The compiler looks at what's on the left side of the assignment (if 1 variable generate the code to return 1 value;
>   if 2 variables, generate the code to return 2 values). The comma ok idiom is used in Go when you want to differentiate
>   between reading a value and getting back the zero value.
> - When you have related data that you want to group together, you should define a `struct`.
> - Go doesn't have classes, because it doesn't have inheritance. This doesn't mean Go doesn't have some of the features of
>   object-oriented languages, it just does things a little differently.
> - A struct type that's defined within a function can only be used within that function.

```go
// Parentheses = conversion
// Type Conversion: []type(value)
[]rune("hello")     // convert string → slice of runes
[]byte("hello")     // convert string → slice of bytes
int(3.14)           // convert float → int
float64(42)         // convert int → float
string([]rune{72})  // convert rune slice → string

// Curly braces = create new slice
// Literal Creation: []type{values}
[]string{"hello", "world"}
[]int{1, 2, 3}
[]rune{'H', 'e', 'l', 'l', 'o'}
[]byte{72, 101, 108}
```

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
- Go provides a built-in data type for situations where you want to associate one value to another.
  The `map` type is written as `map[keyType]valueType`.
- You can declare a `map` using a `var` to create a map variable that's set to its zero value.
  In this case (`var nilMap map[string]int`) `nilMap` is declared to be a map with `string` keys and `int` values.
  The zero value for a map is `nil`. A `nil` map has a length of `0`. Attempting to read a `nil` map always returns
  the zero value for the map's value type. However, _attempting to write to a `nil` map variable causes a panic._
  ```go
  // 
	var nilMap map[string]int
	fmt.Println(nilMap) // map[]
	fmt.Println(nilMap["test"]) // 0 (notice no panic and returns zero value for int)
  nilMap["test"] = 1 // panic: assignment to entry in nil map

  ```
- You can use a `:=` declaration to create a map variable by assigning it a _map literal_.
  In this case, you are using an empty map literal. This is not the same as a `nil` map.
  It has a length of `0`, but you can read and write to a map assigned an empty map literal.
  ```go
	teams := map[string]int{} // Initialize with an empty map literal
	teams["test"] = 1
	fmt.Println(teams) // map[test:1]
	fmt.Println(teams["test"]) // 1
  ```
- Here is what a non-empty map literal looks like:
  ```go
	teams := map[string][]string{
		"Orcas":   []string{"Fred", "Ralph", "Bijou"},
		"Lions":   []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	}
	fmt.Println(teams)
	fmt.Println(teams["Lions"])
	fmt.Println(teams["test"])

  // map[Kittens:[Waldo Raul Ze] Lions:[Sarah Peter Billie] Orcas:[Fred Ralph Bijou]]
  // [Sarah Peter Billie]
  // []
  ```
- A map literal's body is written as the key, followed by a colon (`:`), then the value.
  A comma separates each key-value pair in the map, even on the last line.
- If you know how many key-value pairs you intend to put in the map but don't know the exact values,
  you can use `make` to create a map with a default size (capacity).
- Maps created with `make` still have a length of `0`, and they can grow past the initially specified size (capacity).
  ```go
	ages := make(map[int][]string, 10)
	fmt.Println(ages) // map[]
	fmt.Println(ages[1]) // []

	ages[1] = []string{"one"}
	fmt.Println(ages, len(ages)) // map[1:[one]] 1
  
  // Maps in Go use a complex internal hash table structure.
  // The concept of "capacity" doesn't translate cleanly to the outside.
  // The capacity hint you pass to make is just an internal optimization; Go doesn't expose it back to you.
  fmt.Println(cap(ages)) // invalid argument: ages (variable of type map[int][]string) for built-in cap
  ```
- Maps are like slices in several ways:
  - Maps automatically grow as you add key-value pairs to them
  - If you know how many key-value pairs you plan to insert into a map, you can use `make` to create a map
    with a specific size (capacity)
  - Passing a map to the `len` function tells you the number of key-value pairs in a map
  - The zero value of a map is `nil`
  - Maps are not comparable. You can check if they are equal to `nil`, but you _cannot_ check if two maps
    have identical keys and values using `==` or differ using `!=`.
- The key for a map can be any comparable type. This means _you cannot use a slice or a map as the key for a map_.
- When should you use a map, and when should you use a slice? You should use slices for lists of data when data should
  be processed sequentially or the order of the elements is important. Maps are useful when you need to organize values
  using something other than an increasing integer value, such as a name.

> [!WARNING]
> In Go, order of map is random. Go intentionally randomizes map iteration to prevent developers from relying on it.
> You have to use a slice of keys alongside the map for sorted order.

### Reading and Writing a Map
- You assign a value to a map key by putting the key within brackets and using `=` to specify the value,
  and you read the value assigned to a map key by putting the key within brackets.
  You cannot use `:=` to assign a value to a map key.
  - When you try to read the value assigend to a mpa key that was never set,
    the map returns the zero value for the map's value type. In this case, the value type is an `int`,
    so you get back a `0`.
  - You can use the `++` operator to increment the numeric value for a map key. Because a map returns its zero value by default,
    this works even when there's no existing value associated with the key.
  ```go
	totalWins := map[string]int{}
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2

	fmt.Println(totalWins["Orcas"]) // 1
	fmt.Println(totalWins["Kittens"]) // 0

	totalWins["Kittens"]++
	fmt.Println(totalWins["Kittens"]) // 1

	totalWins["Lions"] = 3
	fmt.Println(totalWins["Lions"]) // 3
  ```
### The comma ok Idiom
- A map returns the zero value if you ask for the value associated with a key that's not in the map.
  Sometimes, you need to find if a key is in a map. Go provides the _comma ok idiom_ to tell the difference between
  a key that's associated with a zero value and a key that's not in the map:
  - Go determines this at compile time based on the assignment context;
    specifically, how many variables are on the left-hand side.
    The compiler looks at what's on the left side of the assignment (if 1 variable generate the code to return 1 value;
    if 2 variables, generate the code to return 2 values).
  ```go
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}

	v, ok := m["hello"]
	fmt.Println(v, ok) // 5, true
	fmt.Println(m["hello"]) // 5

  v2 := m["hello"]
  fmt.Prinln(v2) // 5

	v, ok = m["world"]
	fmt.Println(v, ok) // 0, true
	fmt.Println(m["world"]) // 0

	v, ok = m["goodbye"]
	fmt.Println(v, ok) // 0, false
	fmt.Println(m["goodbye"]) // 0
  ```
- Rather than assign the result of a map read to a single variable, with the comma ok idiom
  you  assign the result of a map read to two variables. The first gets the value associated with the key.
  The second value returned a `bool`. It is usually named `ok`.
  - If `ok` is `true`, the key is present in the map.
  - If `ok` is `false`, the key is not present.
- The comma ok idiom is used in Go when you want to differentiate between reading a value and getting back the zero value.

### Deleting from Maps
- Key-value pairs are removed from a map via the built-in `delete` function
  - The `delete` function takes a map and a key and then removes the key-value pair with the specified key.
    If the key isn't present in the map or if the map is `nil`, nothing happens.
    The `delete` function doesn't return a value.
  ```go
  	m := map[string]int{
		"hello": 5,
		"world": 10,
	}
	fmt.Println(m, len(m)) // map[hello:5 world:10] 2
	delete(m, "hello")
	fmt.Println(m, len(m)) // map[world:10] 1
  ```

### Emptying a Map
- The `clear` function works on a map as well. A cleared map has its length set to zero, unlike a cleared slice
  (which sets all values in the slice to their zero values).
  ```go
	m := map[string]int{
		"hello": 5,
		"world": 10,
	}
	fmt.Println(m, len(m)) // map[hello:5 world:10] 2
	clear(m)
	fmt.Println(m, len(m)) // map[] 0

	s := []string{"hello", "world"}
	s2 := []int{1, 2, 3, 4}
	fmt.Println(s, len(s)) // [hello world] 2
	fmt.Println(s2, len(s2)) // [1 2 3 4] 4
	clear(s)
	clear(s2)
	fmt.Println(s, len(s)) // [  ] 2
	fmt.Println(s2, len(s2)) // [0 0 0 0] 4
  ```

### Comparing a Map
- There is a `maps` package in the standard library that has two useful functions for comparing if two maps are equal:
  `maps.Equal` and `maps.EqualFunc` (they are analogous to `slices.Equal` and `slices.EqualFunc`).
  ```go
  	m := map[string]int{
		"hello": 5,
		"world": 10,
	}
	n := map[string]int{
		"world": 10,
		"hello": 5,
	}
	fmt.Println(maps.Equal(m, n)) // true
  ```

### Using Maps as Sets
- Go doesn't include a set, but you can use a map to simulate some of its features
  (see third party `golang-set` at github.com/deckarep/golang-set/v2 v2.8.0`).
  Use the key of the map for the type you want to put into the set and use a `bool` for the value.
  ```go
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[50])
	if intSet[100] {
		fmt.Println("100 is in the set")
	}
	if intSet[10] {
		fmt.Println("10 is in the set")
	}
	if !intSet[100] {
		fmt.Println("100 is not in the set")
	}

  // 11 8
  // true
  // false
  // 10 is in the set
  // 100 is not in the set
  ```
- In the above, we want a set of `int`s, so we create a map where the keys are of `int` type and the values are `bool` type.
  We iterate over the values in `vals` using a `for-range` loop to place them into `intSet`, associating each `int` with the
  boolean value `true` (being a member of the set).
- We wrote `11` values into `intSet`, but the length is `8`, because you cannot have duplicate keys in a map.
- Looking for `50` or `100` returns `false` because it's not in `intSet`, which causes the map to return the zero value for the map
  value, and the zero value for a `bool` is `false`.

> [!NOTE]
> Some people prefer to use `struct{}` for the value when a map is being used to implement a set. 
> The advantage is that an empty struct uses zero bytes, while a boolean uses one byte.
> The disadvantage is that using a `struct{}` makes your code clumsier.
> You have a less obvious assignment, and you need to use the comma ok idiom to check if a value is in the set:
> ```go
> intSet := map[int]struct{}{}
> vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
> for _, v := range vals {
> 	intSet[v] = struct{}{}
> }
> fmt.Println(len(vals), len(intSet))
> fmt.Println(intSet[5])
> fmt.Println(intSet[50])
> if _, ok := intSet[5]; ok {
> 	fmt.Println("5 is in the set")
> }
> if _, ok := intSet[50]; !ok {
> 	fmt.Println("50 is not in the set")
> }
>
> // 11 8
> // {}
> // {}
> // 5 is in the set
> // 50 is not in the set
> ```

## Structs
- Maps are a convenient way to store some kind of data, but they have limitations.
  They don't define an API since there's no way to constrain a map to allow only certain keys.
  Also, all values in a map must be of _the same type_. For these reasons, maps are not an ideal way to pass data from
  function to function. When you have related data that you want to group together, you should define a `struct`
  (short for structure; think a class in Python).
- Go doesn't have classes, because it doesn't have inheritance. This doesn't mean Go doesn't have some of the features of
  object-oriented languages, it just does things a little differently.
  ```go
  type person struct {
		name string
		age  int
		pet  string
	}
  ```
- A struct type is defined with the keyword `type`, the name of the struct type, the keyword `struct`, and a pair of braces (`{}`).
  Within the braces, you list the fields in the struct. Just as you put the variable name first
  and the variable type second in a `var` declaration, you put the struct field name first and the struct field type second.
  Also not that unlike in map literals, no commas separate the fields in a struct declaration.
- You can declare a struct inside or outside of a function.
  A struct type that's defined within a function can only be used within that function.
- Once a struct type is declared, you can define variables of that type
  - Here we are using a `var` declaration. Since no value is assigned to `fred`, it gets the zero value for the `person` struct type
    A zero value struct has every field set to the field's zero value.
  ```go
  var fred person
  fmt.Println(fred) // { 0 } (all zero values)
  ```
- A _struct literal_ can be assigned to a variable as well
  ```go
  bob := person{}
  fmt.Println(bob) // { 0 } 
  ```
- Unlike maps, there is no difference between assigning an empty struct literal and not assigning a value at all.
  Both initialize all fields in the struct to their zero values.
  ```go
  type person struct {
      name string
      age  int
  }

  var p1 person = person{}
  var p2 person
  fmt.Println(p1) // {"" 0 false}
  fmt.Println(p2) // {"" 0 false}
  fmt.Println(p1 == p2) // true — they are exactly the same
  
  // Empty map literal — creates a usable map
  m1 := map[string]int{}
  m1["hello"] = 5 // works fine

  // No assignment — map is nil
  var m2 map[string]int
  m2["hello"] = 5 // PANIC! assignment to nil map
  ```
- There are two styles for a nonempty struct literal. First, a struct literal can be specified as a comma-separated list
  of values for the fields inside of braces.
  - When using this struct literal format, a value for every field in the struct must be specified,
    and the values are assigned to fields in the order they were declared in the struct definition (order matters).
  ```go
  type person struct {
		name string
		age  int
		pet  string
	}
	julia := person{
		"Julia",
		40,
		"cat",
	}
	fmt.Println(julia) // {Julia 40 cat}
  ```
- The second style for a nonempty literal style looks like the map literal style.
  - You use the name of the fields in the struct to specify the values.
    This style has some advantages. It allows you to specify the fields in any order, and you don't need to provide
    a value for all fields. Any field not specified is set to its zero value.
  ```go
  type person struct {
		name string
		age  int
		pet  string
	}
	beth := person{
		age:  30,
		name: "Beth",
	}
	fmt.Println(beth) // {Beth 30 } (notice pets is "")
  ```
- You cannot mix the two struct literal styles: either all fields are specified with names, or none of them are.
  - For small structs where all fields are always specified, the simpler struct literal style is fine.
    In other cases, use names. It's more verbose, but it makes clear what value is being assigned to what field
    without having to reference the struct definition. It's also more maintainable. If you initialize a struct without using
    the field names and a future version of the struct adds additional fields, your code will no longer compile.
- A field in a struct is accessed with dot annotation:
  ```go
  bob.name = "Bob"
  fmt.Println(bob.name)
  ```

### Anonymous Structs
- You can declare that a variable implements a struct type without first giving the struct type a name.
  This is called an _anonymous struct_. In the below example, the types of variables `person` and `pet`
  are anonymous structs. You can assign (and read) fields in an anonymous struct just as you do for a named struct type.
  Just as you can initialize an instance of a named struct with a struct literal, you can do the same for an anonymous struct as well.
  ```go
  var person struct {
		name string
		age  int
		pet  string
	}
	person.name = "bob"
	person.age = 50
	person.pet = "dog"

	pet := struct {
		name string
		kind string
	}{
		name: "fido",
		kind: "dog",
	}
	fmt.Println(person) // {bob 50 dog}
	fmt.Println(pet) // {fido dog}
  ```
- Anonymous structs are handy in two common situations: (1) when you translate external data into a struct or a struct into
  external data (like JSON or Protocol Buffers). This is called _unmarshalling_ (translated external into a struct) and 
  _marshalling_ (translating a struct into external data) data. (2) Writing tests: you'll use a slice of anonymous structs
  when writing table-driven tests.
  ```go
  // Example for JSON
  // Instead of defining a named type...
  type person struct {
      Name string `json:"name"`
      Age  int    `json:"age"`
  }

  // You can use an anonymous struct
  person := struct {
      Name string `json:"name"`
      Age  int    `json:"age"`
  }{
      Name: "Alice",
      Age:  30,
  }
  ```
  ```go
  // Unmarshalling (JSON to struct)
  data := `{"name": "Alice", "age": 30, "email": "alice@example.com"}`

  // Only extract the fields you care about
  var result struct {
      Name string `json:"name"`
      Age  int    `json:"age"`
  }

  json.Unmarshal([]byte(data), &result)
  fmt.Println(result.Name) // Alice
  fmt.Println(result.Age)  // 30
  // email is ignored since we didn't define a field for it
  ```
  ```go
  // Marshalling (strcut to JSON)
  data, _ := json.Marshal(struct {
      Name  string `json:"name"`
      Age   int    `json:"age"`
      Admin bool   `json:"is_admin"`
  }{
      Name:  "Alice",
      Age:   30,
      Admin: true,
  })

  fmt.Println(string(data))
  // {"name":"Alice","age":30,"is_admin":true}
  ```

### Comparing and Converting Structs
- Whether a struct is comparable depends on the struct's fields. Structs that are entirely composed of comparable types are
  comparable; those with slice or map fields are not.
- Unlike Python, in Go there's no magic method that can be overridden to redefine equality and make `==` and `!=` work
  for incomparable structs. You can write your own function that you use to compare structs.
- Just as Go doesn't allow comparisons between variables of different primitive types, Go doesn't allow comparisons
  between variables that represent structs of different types. Go does allow you to perform a type conversion from one
  struct type to another _if the fields of both structs have the same names, order, and types_.
- In the below example, you can use a type conversion to convert an instance of `firstPerson` to `secondPerson`,
  but you can't use `==` to compare an instance of `firstPerson` and an instance of `secondPerson`, because they are different types.
  - You can't convert an instance of `firstPerson` to `thirdPerson`, because the fields are in a different order.
  - You can't convert an instance of `firstPerson` to `fourthPerson`, because the field names don't match.
  - You can't convert an instance of `firstPerson` to `fifthPerson` because there's an additional field.
  ```go
  type firstPerson struct {
		name string
		age int
	}
	type secondPerson struct {
		name string
		age int
	}
	type thirdPerson struct {
		age int
		name string
	}
	type fourthPerson struct {
		firstName string
		age int
	}
	type fifthPerson struct {
		name string
		age int
		favoriteColor string
	}

  f := firstPerson{name: "Alice", age: 30}
	s := secondPerson{name: "Alice", age: 30}

  converted := secondPerson(f)
  fmt.Println(converted) // {Alice 30}
  fmt.Println(converted == s) // true
  fmt.Println(f == s) // invalid operation: f == s (mismatched types firstPerson and secondPerson)
  ```
- Anonymous structs add a small twist: if two struct variables are being compared and at least one has a type that's anonymous,
  you can compare them without a type conversion if the fields of both structs have the same names, order, and types.
  You can also assign between names and anonymous struct types if the fields of both structs have the same names, order, and types.
  ```go
  type firstPerson struct {
		name string
		age  int
	}
	f := firstPerson{
		name: "bob",
		age:  40,
	}
	var g struct {
		name string
		age  int
	}
	// compiles -- can use = and == between identical named and anonymous structs
	g = f
	fmt.Println(f == g) // true
  ```
- Said differently: Go relaxes its strict type rules when at least one of the structs is anonymous.
  Two named types are always considered different (`firstPerson` and `secondPerson`), even if their fields are identical.
  - The exception to this is anonymous structs.
  - Named types have an explicit identity; `firstPerson` and `secondPerson` are distinct types by name,
    even if structurally identical
  - Anonymous structs have no name; they have no identity other than their structure, so Go just checks if the fields match.
    The same rules about field names, order, and types still apply (can't be `age` then `name`, can't be `firstName`,
    can't have `favoriteColor`).
  ```go
  type firstPerson struct {
      name string
      age  int
  }

  f := firstPerson{name: "bob", age: 40}

  // g is an anonymous struct
  var g struct {
      name string
      age  int
  }

  // Assignment works — no conversion needed!
  g = f

  // Comparison works — no conversion needed!
  fmt.Println(f == g) // true
  ```
