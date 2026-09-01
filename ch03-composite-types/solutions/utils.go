package main

import (
	"fmt"
	"strings"
)

func header(title string) {
	fmt.Println(strings.Repeat("-", 5) + " " + title + " " + strings.Repeat("-", 5))
}
