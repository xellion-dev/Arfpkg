package main

import (
	//"bufio"
	"fmt"
	//"io/ioutil"

	//	"io/ioutil"
	"os"
	"os/exec"
	//"strings"
	//"github.com/pelletier/go-toml"
	//"github.com/pieterclaerhout/go-log"
)

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
		}
	} else {
		fmt.Printf("\nPlease run as root(or sudo)\n")
		os.Exit(1)
	}
}
