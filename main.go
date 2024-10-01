package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/pelletier/go-toml"
)

func main() {
	var oper string
	var pkg string
	tree, err := toml.LoadFile("packages.toml")
	if err != nil {
		panic(err)
	}

	// get value
	fmt.Printf("-----|arfpkg V1.1|-----")
	fmt.Printf("\nEnter operation ")
	fmt.Scanln(&oper)
	if oper == "install" {
		fmt.Println("Enter package name: ")
		fmt.Scanln(&pkg)
		b, err := ioutil.ReadFile("packages.toml")
		if err != nil {
			panic(err)
		}
		s := string(b)

		if strings.Contains(s, pkg) {
			tree.Get("packages." + pkg + ".url")
		} else {
			fmt.Printf("invalid package given")
		}

	}
	fmt.Printf("\nexiting... ")
	os.Exit(0)
}
