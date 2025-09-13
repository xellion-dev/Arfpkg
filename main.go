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
	// ALWAYS clear temp. I've had issues with new packages not showing up
	cleartmp := exec.Command("rm", "-rf", "/bin/arfpkg/temp/packages.toml")
	cleartmp.Run()
	//check if sudo
	if os.Getuid() == 0 {
		// init

		getlist := exec.Command("wget", "-P", "/bin/arfpkg/temp", "https://ixrepo.pages.dev/packages.toml")
		getlist.Run()
		//args := os.Args
		fmt.Printf("-----|arfpkg V1.2|-----")
		fmt.Printf("\nEnter operation ")
		fmt.Scanln(&oper)
		// if install, install
		if oper == "install" || oper == "fetch" {
			//fmt.Printf(args[1])
			install()
			cleartmp.Run()
		} else if oper == "remove" {
			remove()
		} else {
			fmt.Printf("operation '" + oper + "' is not valid.\n")
		}
	} else {
		fmt.Printf("\nPlease run as root(or sudo)\n")
		os.Exit(1)
	}
	cleartmp.Run()
	fmt.Printf("exiting...\n")
}
