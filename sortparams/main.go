package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	lenArgs := len(args)
	swapped := true

	for swapped { // while swapped is true
		swapped = false
		for i := 1; i < lenArgs; i++ {
			fmt.Println(i)
			fmt.Println("- ", args[i-1], args[i])
			if args[i-1] > args[i] { // if before is bigger than this value
				args[i-1], args[i] = args[i], args[i-1] // Swap elements if they are in the wrong order
				fmt.Println(args)
				swapped = true // will keep doing this for till not entering into this if
			}
			fmt.Println(swapped)
			fmt.Println(".")
		}
	}
	for _, arg := range args {
		// Iterate through the characters in each string
		for _, char := range arg {
			z01.PrintRune(char)
			z01.PrintRune('\n')
		}
	}
}
