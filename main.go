package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("---arfpkg v1.0 b1.1---")
	arg := os.Args[1]
	fmt.Println(" " + arg)
}
