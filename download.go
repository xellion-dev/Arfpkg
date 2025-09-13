package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/go-getter/v2"
	"github.com/pelletier/go-toml"
)

func downaload() {
	pkglist, err := toml.LoadFile("/bin/arfpkg/temp/packages.toml")
	url := pkglist.Get("packages." + pkg + ".url").(string)
	xzname := pkglist.Get("packages." + pkg + ".xzname").(string)
	// Download a Git repository
	sourceURL := url
	destDir := "/bin/arfpkg/temp/" + xzname // The destination directory

	result, err := getter.Get(destDir, sourceURL)
	if err != nil {
		log.Fatalf("Error downloading: %v", err)
	}

	fmt.Printf("Successfully downloaded from %s to %s\n", sourceURL, result)

	// Example of downloading a specific file from HTTP
	// sourceURL := "https://example.com/myfile.txt"
	// destDir := "./myfile.txt"
	// result, err := getter.Get(destDir, sourceURL)
	// if err != nil {
	// 	log.Fatalf("Error downloading: %v", err)
	// }
	// fmt.Printf("Successfully downloaded from %s to %s\n", sourceURL, result)
}
