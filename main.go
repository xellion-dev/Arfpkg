package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/pelletier/go-toml"
)

func main() {
	var oper string
	var pkg string
	installcmd := exec.Command("wget", "https://www.nano-editor.org/dist/v8/nano-8.2.tar.xz")
	tree, err := toml.LoadFile("packages.toml")
	if err != nil {
		fmt.Printf("An error has occured")
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
			out, err := installcmd.Output()
			if err != nil {
				fmt.Println("could not run command: ", err)
			}
			fmt.Println("Output: ", string(out))

		} else {
			fmt.Printf("invalid package given")
		}

	}
	fmt.Printf("\nexiting... ")
	os.Exit(0)
}
