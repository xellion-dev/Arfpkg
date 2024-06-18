package main

import (
	"fmt"
)

func main() {

	var oper string
	var pkg string
	fmt.Println("-----|arfpkg V1.1|-----")
	fmt.Println("Enter operation ")
	fmt.Scanln(&oper)

	if oper == "install" {
		fmt.Println("Enter package name: ")
		fmt.Scanln(&pkg)

		fmt.Printf("Package not found.")
	}
	fmt.Printf("\nexiting...")
}
