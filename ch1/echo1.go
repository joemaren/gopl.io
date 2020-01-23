// Implementing the echo unix utility: prints out it's command line arguments
package main

import (
	"os"
	"fmt"
)

func main() {
	var sep, s string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
