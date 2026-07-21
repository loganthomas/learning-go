# Exercises
- [x] Take the "Hello, world!" program and run it on The Go Playground.
   Share a link to the code in the playground with a coworker who would love to learn about Go.
- [x] Add a target to he Makefile called `clean` that removes the `hello_world` binary
   and any other temporary files created by `go build`. Take a look at the Go command documentation
   (https://oreil.ly/uqsMy) to find a `go` command to help implement this.
   (https://pkg.go.dev/cmd/go/internal/clean@go1.26.5)
- [x] Experiment with modifying the formatting in the "Hello, world!" program.
   Add blank lines, spaces, change indentation, insert new  lines.
   After making a modification, run `go fmt` to see if the formatting change is undone.
   Also, run `go build` to see if the code still compiles.
   You can also add additional `fmt.PrintLn` calls
   so you can see what happens if you put blank lines in the middle of a function.
