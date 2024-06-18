package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	name    string
	version int
}

func main() {
	var config Config
	var oper string
	var pkg string
	_, err := toml.DecodeFile("packages.toml", &config)
	if err != nil {
		panic(err)
	}
	fmt.Println("-----|arfpkg V1.1|-----")
	fmt.Println(config.name)
	fmt.Println("Enter operation ")
	fmt.Scanln(&oper)
	if oper == "install" {
		fmt.Println("Enter package name: ")
		fmt.Scanln(&pkg)

		fmt.Printf("Package not found.")
	}
	fmt.Printf("\nexiting... ")
	os.Exit(0)
}
