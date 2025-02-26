# Chapter 1: Setting Up Your  Go Environment

## Installing the Go Tools
- https://go.dev/dl/
    - Apple macOS (x86-64): https://go.dev/dl/go1.24.0.darwin-amd64.pkg
```
$ go version
go version go1.24.0 darwin/amd64
```

## Go Tooling
All of the Go development tools are accessed via the `go` command:
- `go build`: compiler
- `go fmt`: formatter
- `go mod`: dependency manager
- `go test`: test runner
- `go vet`: tool to scan for common coding mistakes

> [!NOTE]
> Since the introduction of Go in 2009, several changes have occurred
> in the way Go developers organize their code and their dependencies.
> Because of this churn, there's lots of conflicting advice, and most
> of it is obsolete (for example, you can safely ignore discussions
> about `GOROOT` and `GOPATH`).
>
> For modern Go development, the rule is simple: you are free to organize
> your projects as you see fit and store them anywhere you want.

## Personal Notes
To create a new module:
```
$ go mod init hello_world
go: creating new go.mod: module hello_world
```
- Every module has a `go.mod` file in its root directory.
  This declares the name of the module, the minimum supported
  version of Go for the module, and any other modules that
  a module depends on. **You can think of it as being similar
  to the `requirements.txt` file used by Python.**
- You shouldn't edit the `go.mod` file directly. Instead,
  use `go get` and `go mod tidy` commands to manage changes
  to the file.

Go enforces a standard format:
```
$ go fmt ./...
```
- Go programs use tabs to indent.
- It is a syntax error if the opening brace is not on the same line as the
  declaration or command that begins the block.
- Remember to run `go fnt` before you compile your code, and, at the very
  least, before you commit source code changes to your repo.
  If you forget, **make a separate commit** that does _only_ `go fmt ./...`
  so you don't hide logic changes in an avalanche of formatting changes.

> [!NOTE]
> GO requires a semicolon at the end of every statement. 
> However, **Go developers should never put the semicolons in themselves.**
> The Go compiler adds them automatically, following a simple rule described
> in Effective Go (https://oreil.ly/hTOHU).

Just as you should run `go fmt` to make sure your code is formatted properly,
run `go vet` to scan for possible bugs in valid code:
```
$ go vet ./...
```
- All Go developers should read through Effective Go (https://oreil.ly/hTOHU)
  and the Code Review Comments page on Go's wiki (https://oreil.ly/FHi_h) to
  understand what idiomatic Go code looks like.

