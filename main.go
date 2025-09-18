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

		//check if run with args or TUI
		if len(os.Args) >= 3 { //if more than 2 args run regular. else; run TUI (was getting null errors)
			oper = os.Args[1]
			if oper == "install" || oper == "fetch" {
				pkg = os.Args[2]
				install()
			}
		} else {
			// run as TUI
			fmt.Printf("-----|arfpkg V1.3|-----")
			fmt.Printf("\nEnter operation ")
			fmt.Scanln(&oper)
			// if install, install
			if oper == "install" || oper == "fetch" {
				fmt.Println("Enter package name: ")
				fmt.Scanln(&pkg)
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
		}
	} else {
		fmt.Printf("\nPlease run as root(or sudo)\n")
		os.Exit(1)
	}
	//ALWAYS CLEAR TEMP!!!!!
	err := os.Remove("/bin/arfpkg/temp/packages.toml")
	if err != nil {
		fmt.Printf("Temp already cleared.\n")
	}
	fmt.Printf("\nAll Done! 🦴")
	fmt.Printf("\nexiting...\n")
	os.Exit(0)
}
