# learning-go
_Learning Go_ (Jon Bodner)
- https://www.oreilly.com/library/view/learning-go-2nd/9781098139285/
- https://github.com/learning-go-book-2e



## Exercise Scaffolding
Each chapter's exercises live in their own Go module under a `solutions/` directory.
The pattern below is established in `ch02-predeclared-types/` and `ch03-composite-types/`
and should be reused for every other chapter.

### File Structure
```
chNN-chapter-name/
├── README.md          # chapter notes and reminders
├── exercises.md       # exercise prompts, checked off as they're solved
└── solutions/
    ├── go.mod         # module solutions
    ├── main.go        # calls each exercise in order
    ├── utils.go       # shared helpers (header, etc.)
    ├── ex1.go
    ├── ex2.go
    └── ex3.go
```

### Setting Up a New Chapter
```
$ mkdir -p chNN-chapter-name/solutions
$ cd chNN-chapter-name/solutions
$ go mod init solutions
```

The module is always named `solutions` so the setup is identical in every chapter.
Each chapter is its own module, so nothing is shared across chapters
and each one can pin its own Go version.

### Running
Run from inside the chapter's `solutions/` directory:
```
$ go run .
```
`go run .` compiles every `.go` file in the package,
which is what makes the one-file-per-exercise split work.
`go run main.go` would fail, since it wouldn't pick up `ex1.go` and friends.

To build a binary instead, send it to `bin/` so it stays out of git:
```
$ go build -o bin/solutions
$ ./bin/solutions
```

### Structuring Exercise Files
Every file in `solutions/` is `package main`.

`main.go` is just the driver, one call per exercise, in order:
```go
package main

func main() {
	ex1()
	ex2()
	ex3()
}
```

`utils.go` holds anything shared across exercises.
Today that's the section header used to separate output:
```go
package main

import (
	"fmt"
	"strings"
)

func header(title string) {
	fmt.Println(strings.Repeat("-", 5) + " " + title + " " + strings.Repeat("-", 5))
}
```

`exN.go` holds a single exercise as a single function named `exN`.
The function opens with a comment restating the problem and what it's testing,
then calls `header` before doing any work:
```go
package main

import (
	"fmt"
)

func ex1() {
	// Exercise 1
	// Short restatement of the problem and the concept it exercises.
	header("exercise 1")

	greetings := []string{"Hello", "Hola"}
	fmt.Println(greetings)
}
```

Longer takeaways that don't belong inside the function
go at the bottom of the file as `// Note` blocks,
so the exercise itself stays readable.

### Adding an Exercise
1. Add the prompt to the chapter's `exercises.md` as an unchecked `- [ ]` item.
2. Create `solutions/exN.go` with the `exN` function.
3. Add the `exN()` call to `main.go`.
4. Run `go run .` to confirm the output.
5. Check the box in `exercises.md`.
