package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	a := os.Args[1]

	for j := 0; j < len(a); j++ {
		if a[j] == ' ' {
			break
		}
		z01.PrintRune(rune(a[j]))
	}
	z01.PrintRune('\n')

}
