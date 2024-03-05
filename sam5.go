package main

import (
	"fmt"
)

func clousre() func() int {
	return func() int {
		return 1
	}
}
func main() {
	fmt.Println(clousre()())
	a := clousre()
	fmt.Println(a())
}
