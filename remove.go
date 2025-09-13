// MODIFIED VERSION OF INSTALL SCRIPT
package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/pelletier/go-toml"
	"github.com/pieterclaerhout/go-log"
)

func remove() {
	fmt.Println("Enter package name: ")
	fmt.Scanln(&pkg)
	// TOML loading
	pkglist, err := toml.LoadFile("/bin/arfpkg/temp/packages.toml")
	if err != nil {
		panic(err)
	}

	b, err := os.ReadFile("/bin/arfpkg/temp/packages.toml")
	if err != nil {
		fmt.Printf("\nFatal error: ")
		fmt.Println(err)
		fmt.Printf("\nexiting...\n")
		os.Exit(1)
	}
	s := string(b)
	url := pkglist.Get("packages." + pkg + ".url").(string)
	xzname := pkglist.Get("packages." + pkg + ".xzname").(string)
	foldername := pkglist.Get("packages." + pkg + ".foldername").(string)
	version := pkglist.Get("packages." + pkg + ".version").(string)
	// exec commands

	// download tarball with cURL
	geturl := exec.Command("curl", "-#", "-o", "/bin/arfpkg/temp/"+xzname, url)
	// make the temp dir for metadata
	mkpkgdir := exec.Command("mkdir", "/bin/arfpkg/temp/"+foldername+version)
	// extract tarball with tar
	extpkg := exec.Command("tar", "-xvf", "/bin/arfpkg/temp/"+xzname, "-C", "/bin/arfpkg/temp/"+foldername+version)
	// enter directory from extracted tarball
	cd := exec.Command("cd", "/bin/arfpkg/temp/"+foldername+version+"/"+pkg+"-latest")

	if strings.Contains(s, pkg) {
		// if it contains a recognized package, ask with Y/N dialog to install
		fmt.Printf("Remove " + pkg + " " + version + "?\n")
		fmt.Printf("Y/N\n")
		fmt.Scanln(&yn)
		if yn == "y" {
			// if yes, start removing
			fmt.Printf("Downloading package data...\n")
			// output cURL progress
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
			fmt.Printf("Extracting...")
			// extract tarball with tar

			mkpkgdir.Run()

			extpkg.Run()
			fmt.Printf("\nBegin Removal...\n")
			// load TOML
			tree2, err := toml.LoadFile("/bin/arfpkg/temp/" + pkg + "-" + version + "/" + pkg + "-latest/" + pkg + ".toml")
			if err != nil {
				panic(err)
			}
			// get dir where binary was installed
			instdir := tree2.Get(pkg + "." + "install-location").(string)
			// remove binary(ies)
			rem := exec.Command("rm", instdir+"/"+pkg)
			// remove temp files
			del := exec.Command("rm", "-rf", "/bin/arfpkg/temp/"+foldername+version)
			// move the tarball to archive
			archive := exec.Command("mv", "/bin/arfpkg/temp/"+xzname, "/bin/arfpkg/package_archive/"+pkg+".tar.xz")
			// remove app from index
			rmindx := exec.Command("rm", "/bin/arfpkg/packages/"+pkg+".toml")
			cd.Run()
			rem.Run()
			fmt.Printf("\nCleaning Up...")
			del.Run()
			archive.Run()
			rmindx.Run()
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
