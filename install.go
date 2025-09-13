package main

import (
	//main.go
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/pelletier/go-toml"
)

type PackageInfo struct {
	Name    string
	Version string
}

func install() {
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
	url := pkglist.Get("packages." + pkg + ".url").(string)
	xzname := pkglist.Get("packages." + pkg + ".xzname").(string)
	foldername := pkglist.Get("packages." + pkg + ".foldername").(string)
	version := pkglist.Get("packages." + pkg + ".version").(string)
	// exec commands
	mkpkgdir := os.MkdirAll
	extpkg := tarxz

	if strings.Contains(s, pkg) {
		// if it contains a recognized package, ask with Y/N dialog to install
		fmt.Printf("Install " + pkg + " " + version + "?\n")
		fmt.Printf("Y/N\n")
		fmt.Scanln(&yn)
		if yn == "y" {
			// if yes, install
			fmt.Printf("Downloading Packages...\n")
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
			fmt.Printf("\nInstalling\n")
			// load TOML
			pkgconf, err := toml.LoadFile("/bin/arfpkg/temp/" + pkg + "-" + version + "/" + pkg + "-latest/" + pkg + ".toml")
			if err != nil {
				panic(err)
			}
			instdir := pkgconf.Get(pkg + "." + "install-location").(string)
			bins := pkgconf.Get(pkg + ".binaries" + ".mainexec").(string)
			// install binary to /bin
			inst := exec.Command("mv", "/bin/arfpkg/temp/"+foldername+version+"/"+pkg+"-latest/"+bins, instdir+"/"+pkg)
			// make the file executable with CHMOD
			chmod := exec.Command("chmod", "+x", instdir+"/"+pkg)
			// clean up temp files
			del := exec.Command("rm", "-rf", "/bin/arfpkg/temp/"+foldername+version)
			// Move the tarball to archive
			archive := exec.Command("mv", "/bin/arfpkg/temp/"+xzname, "/bin/arfpkg/package_archive/"+pkg+".tar.xz")
			//	cd.Run()
			inst.Run()
			chmod.Run()
			fmt.Printf("\nCleaning Up...")
			del.Run()
			archive.Run()

			pkginfo := PackageInfo{
				Name:    pkg,
				Version: version,
			}
			data, err := toml.Marshal(pkginfo)
			if err != nil {
				panic(err)
			}
			err = os.WriteFile("/bin/arfpkg/packages/"+pkg+".toml", data, 0644) // 0644 sets file permissions
			if err != nil {
				panic(err)
			}
			cleartmp := exec.Command("rm -rf", "/bin/arfpkg/temp/packages.toml")
			cleartmp.Run()
			fmt.Printf("\nAll Done\n")
		} else {
			fmt.Printf("Aborted.\n")
		}

	} else {
		fmt.Printf("Package Not Found")
	}
}
