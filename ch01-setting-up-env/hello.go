package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")                    // new line automatically at end of print
	fmt.Printf("Hello,\nworld!")                    // no new line automatically at end of print
	fmt.Printf("Hello,%s\n", "world!")              // force new line
	fmt.Println(fmt.Sprintf("Hello, %s", "world!")) // Hello, world!\n but this is discouraged

}
