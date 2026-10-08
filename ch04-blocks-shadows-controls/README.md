 # Chapter 4: Blocks, Shadows, and Control Structures 

> [!IMPORTANT]
> My personal reminders:
> - Each places where a declaration occurs is called a _block_.
>   Variables, constants types, and functions declared outside of any functions are placed in the _package_ block.
> - When you have a declaration with the same name as an identifier in a containing block you _shadow_ the identifier created
>   in the outer block.
> - When shadowing a variable, the shadowed variable does not disappear or get reassigned;
>   there was just no way to access it once it is shadowed.
> - `:=` is more prone to accidental shadowing than `var` because it allows mixing existing and new variables on the left side.
>   In a new scope, this means `:=` can silently shadow a variable that looks like it's being reused. When using `:=`, make sure
    that you don't have any variables from an outer scope on the lefthand side unless you intend to shadow them.
> - Since shadowing can be useful in a few instances, `go vet` doesn't report it as a likely error, but there are third-party
    tools that can detect accidental shadowing in your code.
> - Any variable declared within the braces of an `if` or `else` statement exists only within that block. Go adds the ability
>   to declare variables that are scoped to the condition and to both the `if` and `else` blocks.
>   Consider using this convenience to keep variable scope as narrow as possible, which reduces the chance of accidentally
>   using or shadowing later.
> - `for` is the only looping keyword in the Go language. There are four formats: a complete, C-style `for`, a condition-only `for`,
>   an infinite `for`, a `for-range`.
> - The `for` statement has three parts: the initialization, the comparison (must return `bool`), the increment.
>   Go allows you to leave off one or more but most commonly the initialization (value calculated before the loop)
>   or increment (complicated inside the loop instead).
> - Go encourages short `if` statement bodies, as left aligned as possible.
>   Nested code is more difficult to understand and follow. Using a `continue` statement makes it
>   easier to understand what's going on.
> - The `for-range` format is for iterating over elements in some of Go's built-in types. You can use a `for-range` loop only
>   to iterate over the built-in compound types and user-defined types that are based on them (things like strings, arrays,
>   slices, maps, channels; but not int, struct). Think of this as "for each" loop. You get two loop variables:
>   the first variable is the position in the data structure being iterated, while the second is the value at that position
>   (think Python's `enumerate`).
> - If you don't need access to the key, use an underscore (`_`) as the variables name. This tells Go to ignore that value.
> - If you want the key, but don't want the value Go allows you to just leave off the second variable (`for k := range ...`) .




## Shadowing Variables
- A _shadowing variable_ is a variable that has the same name as a variable in a containing block.
  For as long as the shadowing variable exists, you cannot access a shadowed variable.
  ```go
  func main() {
    x := 10
    if x > 5 {
      fmt.Println(x)
      x := 5
      fmt.Println(x)
    }
    fmt.Println(x)
  }

  // 10
  // 5
  // 10
  ```
- In the above case, at the first `fmt.Println` inside of the `if` statement, you are able to access the `x` declared
  at the top level of the function. One the next line, though, you _shadow_ `x` by _declaring a new variable with the same name_
  inside the block created by the `if` statement's body. At the second `fmt.Println`, when you access the variable named `x`,
  you get the shadowing variable, which has the value of `5`. The closing brace for the `if` statement's body ends the block,
  where the shadowing `x` exists, and at the third `fmt.Println`, when you access the variable `x`, you get the variable declared
  at the op level of the function, which has the value of `10`.
- `x` didn't disappear or get reassigned; there was just no way to access it once it was shadowed in the inner block.
- It's easy to accidentally shadow a variable when using `:=` because you can use `:=` to create and assign to multiple variables 
  at once. Also, not all variables on the lefthand side have to be new for `:=` to be legal. **You can use `:=` as long as there
  is at least one new variable on the lefthand side.**
- This is what makes `:=` riskier than `var` for accidental shadowing
    - With `var`, **all** variables on the left must be new, so the intent is always clear.
    - With `:=`, you can mix existing and new variables, which makes it easy to _think_ you're reusing a variable when you've
      actually entered a new scope and silently created a shadow:
  ```go
  x := 10

  // In a new scope, var makes it obvious you're declaring a NEW x
  {
     var x, err = someFunc()  // clearly a new x; no abmbiguity
     fmt.Println(x, err)
  }

  // But := can trick you; it looks like you're reusing x:
  {
     x, err := someFunc()  // x is SHADOWED, not reassigned! Legal bc err is new
     fmt.Print(x, err)
  }

  fmt.Println(x)  // still 10 in both cases
  ```
- In the same scope, `var` won't even let you make this mistake. It will refuse to compile.
  But, `:=` happily allows it, since only one variable needs to be new:
  ```go
  x := 10
  var x, err = someFunc()  // Compile error: x already declared

  x := 10
  x, err := someFunc()  // err is new, x is reused (not shadowed here, since in the same scope)
  ```
- Aside on `:=` vs `var`
    - Use `:=` most of the time. It's the default choice inside functions. It's concise and idiomatic:
      ```go
      x := 10
      name := "hello"
      result, err := someFunc()
      ```
    - Use `var` when
      1. Declaring a variable without initializaing it (you want the zero value to be meaningful):
      ```go
      var total int       // zero value of 0 is intentional
      var names []string  // zero value of nil is intentional
      ```
      2. The type isn't what would be inferred:
      ```go
      var x float64 = 10  // without var, 10 would be inferred as int
      ```
      3. Package-level variables (`:=` isn't allowed outside functions):
      ```go
      var config Config
      ```
## `if`
- The `if` statement in Go is much like the `if` statement in most programming languages:
  ```go
  n := rand.Intn(10)  // returns a random int in the range [0, n) (0 inclusive, n exclusive)

  if n == 0 {
    fmt.Println("That's too low")
  } else if n > 5 {
    fmt.Println("That's too big")
  } else {
    fmt.Print("That's a good number:", n)
  }
  ```
- The most visible difference between `if` in Go and other languages is that you don't put parentheses around the condition.
- Any variable declared within the braces of an `if` or `else` statement exists only within that block. Go adds the ability
  to declare variables that are scoped to the condition and to both the `if` and `else` blocks.
  ```go
  if n := rand.Intn(); n == 0 {
    fmt.Println("That's too low")
  } else if n > 5 {
    fmt.Println("That's too big")
  } else {
    fmt.Print("That's a good number:", n)
  }
  ```
- Having this special scope is handy. It lets you create variables that are available **only where they are needed**.
  Once the series of `if/else` statements ends, `n` is undefined. In the first example above, `n` exists in the outer scope.
  In the second version, `n` is scoped to the `if/else` block. The second version is idiomatic Go when you only need `n`
  inside the `if/else` chain. It keeps the variable's scope as narrow as possible, which reduces the chance of accidentally
  using or shadowing it later.
- Example of out of scope (bad):
  ```go
  if n := rand.Intn(10); n == 0 {
    fmt.Println("That's too low")
  } else if n > 5 {
    fmt.Println("That's too big")
  } else {
    fmt.Println("That's a good number:", n)
  }
  fmt.Println(n)  // undefined n
  ```
- Be aware that just like any other block, a variable declared as part of an `if` statement will shadow variables with the same
  name that are declared in containing blocks.

## `for`, Four Ways
- `for` is the only looping keyword in the Go language. There are four formats:
    - A complete, C-style `for`
    - A condition-only `for`
    - An infinite `for`
    - `for-range`

### The Complete `for` Statement (C-style)

```go
for i := 0; i < 10; i ++ {
  fmt.Println(i)
}

// 0
// 1
// 2
// 3
// 4
// 5
// 6
// 7
// 8
// 9
```
- The `for` statement does not use parentheses around its parts. It has three parts separated by a semicolon (`;`):
    - The **initialization** that sets on or more variables before the loop begins (`i := 0`). You _must_ use `:=`
      to initialize the variables; `var` is _not_ legal here. You can shadow a variable here.
          - Because the `for` initialization requires `:=`, it always creates a _new_ variable, which
            will shadow any same-named variable from an outer scope.
            ```go
            i := 10
            fmt.Println(i) // prints 10

            for i := 0; i < 5; i++ {
                fmt.Println(i) // prints 0, 1, 2, 3, 4 — this is a NEW i, scoped to the loop
            }

            fmt.Println(i) // prints 10 — the outer i was never touched
            ```
    - The **comparison**. This must be an expression that evaluates to a `bool`. It is checked immediately _before_
      each iteration of the loop. If the expression evaluates to `true`, the loop is executed. Here, it's the `i < 10`
    - The **increment**. You usually see something like `i++` but any assignment is valid. It runs immediately after
      each iteration of the loop, before the condition is evaluated.
- Go allows you to leave out one or more of the three parts of the `for` statement.
  Most commonly, you'll either leave off the initialization if it is based on a value calculated before the loop:
  ```go
  i := 0
  for ; i < 10; i++ {
    fmt.Println(i)
  }
  ```
  or you'll leave off the increment because you have a more complicated increment rule _inside_ the loop:
  ```go
  for i := 0; i < 10; {
    fmt.Println(i)
    if i % 2 == 0 {
      i ++
    } else {
      i+=2
    }
  }
  ```

### The Condition-Only `for` Statement
- When you leave off _both_ the initialization and the increment in a `for` statement, do not include the semicolons.
  This leaves a `for` statement that functions like a `while` loop in other languages.

  ```go
  i := 1

  for i < 100 {
    fmt.Println(i)
    i = i * 2
  }

  // 1
  // 2
  // 4
  // 8
  // 16
  // 32
  // 64
  ```

### The Infinite `for` Statement 
- The loop, removes the condition as well as the initialization and the increment.
  ```go
  package main

  import "fmt"

  func main() {
    for {
      fmt.Println("hello")
    }
  }
  // hello
  // hello
  // hello
  // hello
  // hello
  // hello
  // hello
  // ...
  ```
- How do you get out of an infinite `for` loop? `break`. It exists the loop immediately.
- Go also includes the `continue` keyword, which skips over the rest of the `for` loop's body
  and proceeds directly to the next iteration. Technically, you don't need a `continue` statement.
  You _could_ write code like this:
  ```go
  import "fmt"

  func main() {
      for i := 1; i <= 100; i++ {
          if i%3 == 0 {
              if i%5 == 0 {
                  fmt.Println("fizzbuzz")
              } else {
                  fmt.Println("fizz")
              }
          } else if i%5 == 0 {
              fmt.Println("buzz")
          } else {
              fmt.Println(i)
          }
      }
  }
  ```
- But this is not idiomatic! Go encourages short `if` statement bodies, as left aligned as possible.
  Nested code is more difficult to understand and follow. Using a `continue` statement makes it
  easier to understand what's going on.
  ```go
  import "fmt"

  func main() {
      for i := 1; i <= 100; i++ {
          if i%3 == 0 && i%5 == 0 {
              fmt.Println("fizzbuzz")
              continue
          }
          if i%3 == 0 {
              fmt.Println("fizz")
              continue
          }
          if i%5 == 0 {
              fmt.Println("buzz")
              continue
          }
          fmt.Println(i)
      }
  }
  ```
- Replacing chains of `if`/`else` statements with a series of `if` statements that use `continue` makes the conditions line up.
  This improves the layout of your conditions, which means your code is easier to read and understand.

### The `for-range` Statement
- The `for-range` format is for iterating over elements in some of Go's built-in types. You can use a `for-range` loop only
  to iterate over the built-in compound types and user-defined types that are based on them (things like strings, arrays,
  slices, maps, channels; but not int, struct). Think of this as "for each" loop.
  ```go
  // Example using a slice
  evenVals := []int{2, 4, 6, 8, 10, 12}
  for i, v := range evenVals {
    fmt.Println(i, v)
  }
  // 0 2
  // 1 4
  // 2 6
  // 3 8
  // 4 10
  // 5 12
  ```
- What makes a `for-range` loop interesting is that you get two loop variables. The first variable is the position in the data
  structure being iterated, while the second is the value at that position (think Python's `enumerate`).
    - The idiomatic names for the two loop variables depend on what is being looped over.
    - For an array, slice, or string, an `i` for _index_ is commonly used and when iterating through a map, `k` for key is used instead.
    - For the `v` value, single-letter variable names work well but for longer or more complex loops use a more descriptive name.
- If you don't need access to the key, use an underscore (`_`) as the variables name. This tells Go to ignore that value.
  ```go
  evenVals := []int{2, 4, 6, 8, 10, 12}
  for _, v := range evenVals {
    fmt.Println(v)
  }
  ```
- If you want the key, but don't want the value Go allows you to just leave off the second variable:
  ```go
  uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
  for k := range uniqueNames {
    fmt.Println(k)
  }
  // Fred
  // Raul
  // Wilma
  ```


