# Chapter 2: Predeclared Types and Declarations

> [!IMPORTANT]
> My personal reminders:
> - Golang principle: write your programs in a way that makes your intentions clear.
> - In Go, single quotes and double quotes are NOT interchangeable.
> - Go enforces strict type safety, requiring explicit conversions to prevent subtle bugs from implicit type coercion.
>   (`x int` and `y float64`, need to use `float64(x) + y` or `x + int(y)`)
> - Go doesn't allow for "truthy" values. _No other type can be converted to a bool, implicitly or explicitly_.
>   To convert from another data type to boolean, you must use one of the comparison operators
>   (`x == 0` or `s ==""`).
> - Go has a lot of ways to declare variables. Each declaration style communicates something about how the
>   variable is used.
> - If you are declaring a variable at the package level, you must use `var` because `:=` is not legal outside of functions.
> - While it is legal to use a type conversion to specify the type of the value and use `:=` to write `x := byte(20)`
>   it is idiomatic to write `var x byte = 20`.
> - You should rarely declare variables outside of functions, in what's called the _package block_.
>   Package-level variables whose values change are a bad idea. As a general rule, you should only declare variables
>   in the package block that are effectively immutable. 
> - Constants in Go are a way to give names to literals. There is NO WAY in Go to declare that a variable is immutable.
> - Go requirement that _every declared local variable must be read_.
>   It is a _compile-time error_ to declare a local variable and to not read its value.
> - Idiomatic Go uses camel case (names like `indexCounter` and `numberTries`).
> - Go does not write constants in all uppercase (Go uses the case of the first letter in the name of a 
>   package-level declaration to determine if the item is accessible outside of the package).
> - Within a function, favor short variable names: _the smaller the scope for a variable, the shorter the name that's used for it_.



## The Predeclared Types
- Go has many types built into the language. These are called _predeclared_ types 
  (booleans, integers, floats, and strings).
- The Zero Value: go assigns a defalut _zero value_ to any variable that is declared
  but not assigned a value.

### Literals
- A Go _literal_ is an explicitly specified number, character, or string.
  Go programs have four common kinds of literals
- An _integer literal_ is a sequence of numbers (base 10 by default; can put underscores for thousands separator).
- A _floating-point literal_ has a decimal point to indicate the fractional portion of the value.
  They can also have an exponential specified with the letter `e`.
- A _rune literal_ represents a character and is surrounded by single quotes. **Unlike many other languages,
  in Go single quotes and double quotes are NOT interchangeable.** There are several backslash-escaped rune
  literals (newline '\n', tab '\t', single quote '\'', backslash '\\').
- There are two types of _string literals_: _interpreted string literal_ and _raw string literals_.
  - Interpreted string literal uses double quotes: `"Greetings and Salutations"`.
    These contain zero or more rune literals. They are called "interpreted" because they interpret rune literals
    into single characters.
  - If you need to include backslashes, double quotes, or newlines in your string, using a raw string literal
    is easier. These are delimited with backquotes and can contain any character except backquotes.
    There's no escape character in a raw string literal; all characters are included as is.
    
    ```go
    # interpreted string literal
    "Greetings and\n\"Salutations\""

    # raw string literal
    `Greetings and
    "Salutations"`
    ```
  - Literals are considered _untyped_.

### Booleans
- The `bool` type represents Boolean variables (can be `true` or `false`; the zero value for `bool` is `false).

  ```go
  var flag bool        // no value assigned, set to false
  var isAwesome = true // type is inferred automatically by Go compiler
  ```

### Numeric Types
- 12 types that are grouped into three categories (integer, floating point, complex).
    - (8) `int8`, `int16`, `int32`, `int64`, `uint8`, `uint16`, `uint32`, `uint64`
    - (2) `float32`, `float64`
    - (2) `complex64`, `complex128`
- Go provides more integer types than some other languages. Follow these three rules:
    - If you are working with a binary file format or network protocol that has an integer
      of a specific size or sign, use the corresponding integer type.
    - If you are writing a library function that should work with an integer type,
      take advantage of Go's generics support and use a generic type parameter to
      represent any integer type.
    - In all other cases, just use `int`.

### A Taste of Strings and Runes
- The zero value for a string is an empty string.
- Like integers and floats, strings are compared for equality using `==`, difference with `!=`,
  or ordering with `>`, `>=`, `<`, or `<=`. They are concatenated using the `+` operator.
- Strings in Go are immutable; you can reassign the value of a string variable, but you cannot
  change the value of the string that is assigned to it.
- Go has a type that represents a single code point. The _rune_ type is an alias for the `int32` type,
  just as `byte` is an alias for `uint8`.
- A rune literal's default type is a rune, and a string literal's default type is a string.
- If you are referring to a character, use the `rune` type, not the `int32` type:

  ```go
  var myFirstInitial rune = 'J' // good - the type name matches the usage
  var myLastInitial int32 'B'   // bad - legal but confusing
  ```

### Explicit Type Conversion
- Go DOES NOT allow _automatic type promotion_ (where numeric types automatically convert from one
  to another when needed). You must use a _type conversion_ when variable types do not match.
  This makes it clear exactly what type you want without having to memorize any type conversion rules.

  ```go
  var x int = 10
  var y float64 = 30.2
  var sum1 float64 = float64(x) + y
  var sum2 int = x + int(y)
  fmt.Println(sum1, sum2)
  ```
- This strictness around types has other implications. Since all type conversions in Go are explicit,
  you cannot treat another Go type as a boolean. In many other languages, a nonzero number or a nonempty
  string can be interpreted as a boolean `true`. Go doesn't allow for "truthy" values. _No other type
  can be converted to a bool, implicitly or explicitly_. If you want to convert from another data type
  to boolean, you must use one of the comparison operators (`==`, `!=`, `>`, `>=`, `<`, `<=`).
  - To check if `x` is equal to `0`, the code would be `x == 0`.
  - To check if string `s` is empty, use `s == "`.

### Literals Are Untyped
- While you can't add two integer variables together if they are declared to be of different types of integers,
  Go lets you use an integer literal in floating-point expressions or even assign an integer literal to a 
  floating-point variable:
- Remember, Go literals are untyped. Go tries to avoid forcing a type until the developer specifies one.

  ```go
  var x float64 = 10
  var y float 64 = 200.3 * 5
  ```

## `var` Versus `:=`
- Go has a lot of ways to declare variables. Each declaration style communicates something about how the
  variable is used.
- The most verbose way to declare a variable in Go uses the `var` keyword, an explicit type,
  and an assignment.

  ```go
  # var, type, assignment
  var x int = 10
  ```

- If the type on the righthand side of the `=` is the expected type of your variable,
  you can leave off the type from the left side of the `=`.

  ```go
  // Since the default type of an integer literal is int,
  // the following declares x to be a variable of type int
  var x = 10
  ```
- You can declare multiple variables at once with `var`,
  and they can be of the same, declared to be all zero values of the same type, or of different types:

  ```go
  // All same type
  var x, y int = 10, 20

  // All zero values of same type
  var x, y int

  // Different types
  var x, y = 10, "hello"
  ```
- If you are declaring multiple variables at once, you can wrap them in a _declaration list_:

  ```go
  var (
      x    int
      y         = 20
      z    int  = 30
      d, e      = 40, "hello"
      f, g string

  )
  ```
- Go supports a short declaration and assignment format.
  **When you are within a function**, you can use the `:=` operator to replace the `var` declaration
  that uses type inference. The following two statements do exactly the same thing:

  ```go
  var x = 10
  x := 10
  ```
- As with `var`, you can declare multiple variables at once:

  ```go
  var x, y = 10, "hello"
  x, y := 10, "hello"
  ```
- The `:=` operator can do one trick that you cannot do with `var`:
  it allows you to assign values to existing variables too. As long as at least one new variable is on the lefthand
  side of the `:=`, any of the other variables can already exist:

  ```go
  x := 10
  x, y := 30, "hello"
  ```
- Using `:=` has one limitation. If you are declaring a variable at the package level, you must use `var`
  because `:=` is not legal outside of functions.
- How do you know which style to use?
    - Choose what makes your intent clearest.
    - The most common declaration style _within a function_ is `:=`.
    - Outside of a function, use declaration lists on the rare occasion when you are declaring multiple package-level variables.
    - In some situations, you should avoid using `:=`:
        - When initializing a variable to its zero value, use `var x int`. This makes it clear that the zero value is intended.
        - When assigning an untyped constant or literal to a variable and the default type for the constant or literal
          isn't the type you want for the variable, use the long `var` form with the type specified. While it is legal to use
          a type conversion to specify the type of the value and use `:=` to write `x := byte(20)` it is idiomatic to write
          `var x byte = 20`.
        - Because `:=` allows you to assign to both new and existing variables, it sometimes creates new variables when you
          think you are reusing existing ones. In those situations, explicitly declare all your new variables with `var` to
          make it clear which variables are new, and then use the assignment operator (`=`) to assign values to both new
          and old variables.
- You should rarely declare variables outside of functions, in what's called the _package block_.
    - Package-level variables whose values change are a bad idea.
    - When you have a variable outside of a function, it can be difficult to track the changes made to it, which makes it hard
      to understand how data is flowing through your program.
    - **As a general rule, you should only declare variables in the package block that are effectively immutable.** 

## Using `const`
- Many languages have a way to declare a value as immutable. In Go, this is done with the `const` keyword.
- You can declare a constant at the package level or within a function. Just as with `var`,
  you can (and should) declare a group of related constants within a set of parentheses:

  ```go
  package main

  import "fmt"

  const x int64 = 10

  const (
    idKey   = "id"
    nameKey = "name"
  )

  const z = 20 * 10

  // this code will not compile
  // ./main.go:23:2: cannot assign to x (constant 10 of type int64)
  // ./main.go:24:2: cannot assign to y (untyped string constant "hello")
  // on the Go Playground at https://oreil.ly/FdG-W
  func main() {
    const y = "hello"

    fmt.Println(x)
    fmt.Println(y)

    x = x + 1
    y = "bye"

    fmt.Println(x)
    fmt.Println(y)
  }
  ```
- Go doesn't provide a way to specify that a value calculated at runtime is immutable.
  For example, the following code will fail to compile with the error `x + y (value of type int) is not constant`:

  ```go
  x := 5
  y := 10
  const z = x + y // this won't compile
  ```
- There are no immutable arrays, slices, maps, or structs, and there's no way to declare that a field in a struct
  is immutable. This is less limiting than is sounds. Within a function, it is clear if a variable is being modified,
  so immutability is less important.

## Typed and Untyped Constants
- Constants can be typed or untyped.
    - An untyped constant works exactly like a literal; it has no type of its own
      but does have a default type that is used when no other type can be inferred.
    - A typed constant can be directly assigned only to a variable of that type.
- Whether to make a constant typed depends on why the constant was declared.
    - If you are giving a name to a mathematical constant that should be used with multiple numeric types,
      keep the constant _untyped_. (In general, leaving a constant untyped gives you more flexibility).
    - In certain situations, you'll want to enforce a type. (See enumerations with `iota` later on).
- An untyped constant declaration:

  ```go
  const x = 10

  // All of the following assignments are legal
  var y int = x
  var z float64 = x
  var d byte = x
  ```
- A typed constant declaration:

```go
const typedX int 10

// Can be assigned directly only to an int
// Assigning to any other type produces a compile-time error:
// cannot use typedX (type int) as type float64 in assignment.
```

## Unused Variables
- Go requirement that _every declared local variable must be read_. It is a _compile-time error_ to declare a local variable
  and to not read its value.
    - The compiler's unused variable check is not exhaustive. As long as a variable is read once, the compiler won't complain,
      even if there are writes to the variable that are never read.
- The following is a valid Go program. While the compiler and `go vet` do not catch the unused assignments of `10` and `30` to `x`,
  third-party tools can detect them.

  ```go
  func main() {
    x := 10 // this assignment isn't read!
    x = 20
    fmt.Println(x)
    x = 30 // this assignment isn't read!
  ```
- Surprisingly, Go will allow you to create unread constants with `const`. This is because constants in Go are calculated
  at compile time and cannot have any side effects. This makes them easy to eliminate: if a constant isn't used, it is simply
  not included in the compiled binary.

## Naming Variables and Constants
- There is a difference between Go's rules for naming variables and the patterns that Go developers follow when naming
  their variables and constants.
- Go requires identifier names to start with a letter or underscore, and the name can contain numbers, underscores, and letters.
    - Any Unicode character considered a letter or digit is allowed (but don't be clever or cheeky).
- Even though the underscore is a valid character in a variable name, it is rarely used, because idiomatic Go doesn't use
  snake case (names like `index_counter` or `number_tries`). Instead, idiomatic Go uses camel case (names like `indexCounter`
  and `numberTries`).
- In other languages, constants are always written in all uppercase letters, with words separated by underscores (names like `INDEX_COUNTER` or
  `NUMBER_TRIES`). **Go does not follow this pattern!** This is because Go uses the case of the first letter in the name of a 
  package-level declaration to determine if the item is accessible outside of the package.
- Within a function, favor short variable names: _the smaller the scope for a variable, the shorter the name that's used for it_.
    - It is common in Go to see single-letter variable names used with `for loops`. For example, the names `k` and `v` (short
      for `key` and `value`) are used as variable names in a `for-range` loop. If you are using a standard `for loop`,
      `i` and `j` are common names for the index variable.
    - These shore names sever two purposes. The first is that they eliminate repetitive typing, keeping your code shorter.
      Second, they serve as a check on how complicated your code is. If you find it hard to keep track of your short-named
      variables, your block of code is likely doing too much.
- When naming variables and constants in the package block, use more descriptive names. The type should still be excluded from
  the name, but since the scope is wider, you need a more complete name to clarify what the value represents.
