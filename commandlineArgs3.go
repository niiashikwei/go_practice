package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(os.Args[0])

	var s string
	sep := ""
	for index, arg := range os.Args[1:] {
		s += fmt.Sprint(sep, index, sep, arg)
		sep = " "
	}

	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])
}
