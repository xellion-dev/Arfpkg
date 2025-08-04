package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/pelletier/go-toml"
	"github.com/pieterclaerhout/go-log"
)

func main() {
	var oper string
	var pkg string
	var yn string
	//check if sudo
	if os.Getuid() == 0 {
		tree, err := toml.LoadFile("packages.toml")
		if err != nil {
			panic(err)
		}
		// init
		fmt.Printf("-----|arfpkg V1.1|-----")
		fmt.Printf("\nEnter operation ")
		fmt.Scanln(&oper)
		// if install, install
		if oper == "install" {
			fmt.Println("Enter package name: ")
			fmt.Scanln(&pkg)
			// TOML loading
			b, err := ioutil.ReadFile("packages.toml")
			if err != nil {
				panic(err)
			}
			s := string(b)
			url := tree.Get("packages." + pkg + ".url").(string)
			xzname := tree.Get("packages." + pkg + ".xzname").(string)
			foldername := tree.Get("packages." + pkg + ".foldername").(string)
			version := tree.Get("packages." + pkg + ".version").(string)
			if strings.Contains(s, pkg) {
				// if it contains a recognized package, ask with Y/N dialog to install
				fmt.Printf("Install " + pkg + " " + version + "?\n")
				fmt.Printf("Y/N\n")
				fmt.Scanln(&yn)
				if yn == "y" {
					// if yes, install
					fmt.Printf("Downloading Packages...\n")
					geturl := exec.Command("curl", "-#", "-o", "/bin/arfpkg/temp/"+xzname, url)
					r, _ := geturl.StdoutPipe()
					geturl.Stderr = geturl.Stdout
					done := make(chan struct{})
					scanner := bufio.NewScanner(r)
					go func() {
						for scanner.Scan() {
							line := scanner.Text()
							log.Info(line)
						}
						done <- struct{}{}
					}()
					err := geturl.Start()
					log.CheckError(err)
					<-done
					err = geturl.Wait()
					log.CheckError(err)
					geturl.Run()
					fmt.Printf("Extracting...")
					mkpkgdir := exec.Command("mkdir", "/bin/arfpkg/temp/"+foldername+version)
					mkpkgdir.Run()
					extpkg := exec.Command("tar", "-xvf", "/bin/arfpkg/temp/"+xzname, "-C", "/bin/arfpkg/temp/"+foldername+version)
					extpkg.Run()
					fmt.Printf("\nInstalling\n")

					tree2, err := toml.LoadFile("/bin/arfpkg/temp/tar-1.35/tar-latest/tar.toml")
					if err != nil {
						panic(err)
					}
					instdir := tree2.Get(pkg + "." + "install-location").(string)
					bins := tree2.Get(pkg + ".binaries" + ".mainexec").(string)
					cd := exec.Command("cd", "/bin/arfpkg/temp/"+foldername+version+"/"+pkg+"-latest")
					cd.Run()
					inst := exec.Command("mv", "/bin/arfpkg/temp/"+foldername+version+"/"+pkg+"-latest/"+bins, instdir+"/"+pkg)
					inst.Run()
					chmod := exec.Command("chmod", "+x", instdir+"/"+pkg)
					chmod.Run()
					fmt.Printf("\nCleaning Up...")
					del := exec.Command("rm", "-rf", foldername+version)
					del.Run()
					fmt.Printf("\nAll Done\n")
				} else {
					fmt.Printf("Aborted.\n")
				}

			} else {
				fmt.Printf("Package Not Found")
			}
			fmt.Printf("\nexiting... \n")
			os.Exit(0)
		}
	} else {
		fmt.Printf("\nPlease run as root(or sudo)\n")
		os.Exit(1)
	}
}
