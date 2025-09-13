package main

import (
	"fmt"
	"os"
)

// global vars
var oper string
var pkg string
var yn string

func main() {
	//check if sudo
	if os.Getuid() == 0 {
		// init
		err := download("packages.toml", "https://ixrepo.pages.dev/packages.toml")
		if err != nil {
			fmt.Printf("Could not get packages.\n")
			os.Exit(1)

		}
		//args := os.Args
		fmt.Printf("-----|arfpkg V1.2|-----")
		fmt.Printf("\nEnter operation ")
		fmt.Scanln(&oper)
		// if install, install
		if oper == "install" || oper == "fetch" {
			install()
			err := os.Remove("/bin/arfpkg/temp/packages.toml")
			if err != nil {
				return
			}
		} else if oper == "remove" {
			remove()
		} else {
			fmt.Printf("operation '" + oper + "' is not valid.\n")
		}
	} else {
		fmt.Printf("\nPlease run as root(or sudo)\n")
		os.Exit(1)
	}
	err := os.Remove("/bin/arfpkg/temp/packages.toml")
	if err != nil {

	}
	fmt.Printf("exiting...\n")
}
