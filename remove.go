// MODIFIED VERSION OF INSTALL SCRIPT
// if someone would like to clean up and rewrite, that would be greatly appreciated
package main

import (
	//main.go
	"fmt"
	"os"
	"strings"

	"github.com/pelletier/go-toml"
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
		os.Exit(1)
	}
	s := string(b)

	// TOML refs
	url := pkglist.Get("packages." + pkg + ".url").(string)
	xzname := pkglist.Get("packages." + pkg + ".xzname").(string)
	foldername := pkglist.Get("packages." + pkg + ".foldername").(string)
	version := pkglist.Get("packages." + pkg + ".version").(string)

	// easier just to point and not make the code unbearbly long
	move := os.Rename
	mkpkgdir := os.MkdirAll
	extpkg := tarxz
	del := os.Remove
	archive := move
	cleartmp := os.Remove

	if strings.Contains(s, pkg) {
		// if it contains a recognized package, ask with Y/N dialog to install
		fmt.Printf("remove " + pkg + " " + version + "?\n")
		fmt.Printf("Y/N\n")
		fmt.Scanln(&yn)
		if yn == "y" {
			// if yes, begin removing
			fmt.Printf("Downloading Package Metadata...\n")
			// download files with download function
			err := download(xzname, url)
			if err != nil {
				return
			}

			fmt.Printf("Extracting...")
			// extract tarball with tar
			// make package directory
			err = mkpkgdir("/bin/arfpkg/temp/"+foldername+version, 0755)
			if err != nil {
				fmt.Println("Error (cannot continue): ", err)
			}
			extpkg(xzname, foldername, version)

			fmt.Printf("\nRemoving...\n")
			// load package TOML
			pkgconf, err := toml.LoadFile("/bin/arfpkg/temp/" + pkg + "-" + version + "/" + pkg + "-latest/" + pkg + ".toml")
			if err != nil {
				panic(err)
			}
			instdir := pkgconf.Get(pkg + "." + "install-location").(string)
			//	bins := pkgconf.Get(pkg + ".binaries" + ".mainexec").(string)
			err = del(instdir + "/" + pkg)
			if err != nil {
				return
			}

			fmt.Printf("\nCleaning Up...\n")
			err = del("/bin/arfpkg/temp/" + foldername + version)
			if err != nil {
				return
			}
			err = archive("/bin/arfpkg/temp/"+xzname, "/bin/arfpkg/package_archive/"+pkg+".tar.xz")
			if err != nil {
				return
			}

			err = del("/bin/arfpkg/packages/" + pkg + ".toml") // delete package TOML; Unused for now
			if err != nil {
				panic(err)
			}

			err = cleartmp("/bin/arfpkg/temp/packages.toml")
			if err != nil {
				return
			}
			fmt.Printf("\nAll Done! 🦴\n")
		} else {
			// if no, abort
			fmt.Printf("Aborted.\n")
		}

	} else {
		// if no package found, exit
		fmt.Printf("Package Not Found/Installed")
	}
}
