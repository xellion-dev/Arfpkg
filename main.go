package main

import (
	"fmt"
	"os"
	"os/exec"
)

// global vars
var oper string
var pkg string
var yn string

func main() {
	//check if sudo
	if os.Getuid() == 0 {
		// init
		getlist := exec.Command("wget", "-P", "/bin/arfpkg/temp", "https://ixrepo.pages.dev/packages.toml")
		getlist.Run()

		fmt.Printf("-----|arfpkg V1.1|-----")
		fmt.Printf("\nEnter operation ")
		fmt.Scanln(&oper)
		// if install, install
		if oper == "install" {
			install()

		} else if oper == "remove" {
			remove()
		} else {
			fmt.Printf("operation '" + oper + "' is not valid.\n")
		}
	} else {
		fmt.Printf("\nPlease run as root(or sudo)\n")
		os.Exit(1)
	}
}
