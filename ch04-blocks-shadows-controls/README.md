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
