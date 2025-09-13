package main

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/ulikunitz/xz"
)

// It would be greatly appreciated if someone could annotate. I don't really have time :(
func tarxz(xzname string, foldername string, version string) {
	xzFile, err := os.Open("/bin/arfpkg/temp/" + xzname)
	if err != nil {
		log.Fatalf("Failed to open xz file: %v", err)
	}
	defer xzFile.Close()
	xzReader, err := xz.NewReader(xzFile)
	if err != nil {
		log.Fatalf("Failed to create xz reader: %v", err)
	}
	tarReader := tar.NewReader(xzReader)
	destDir := "/bin/arfpkg/temp/" + foldername + version
	if err := os.MkdirAll(destDir, 0755); err != nil {
		log.Fatalf("Failed to create destination directory: %v", err)
	}
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break // EOF
		}
		if err != nil {
			log.Fatalf("Error reading tar header: %v", err)
		}

		targetPath := filepath.Join(destDir, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(targetPath, os.FileMode(header.Mode)); err != nil {
				log.Fatalf("Failed to create directory %s: %v", targetPath, err)
			}
		case tar.TypeReg:
			outFile, err := os.OpenFile(targetPath, os.O_CREATE|os.O_WRONLY, os.FileMode(header.Mode))
			if err != nil {
				log.Fatalf("Failed to create file %s: %v", targetPath, err)
			}
			defer outFile.Close()

			if _, err := io.Copy(outFile, tarReader); err != nil {
				log.Fatalf("Failed to write file %s: %v", targetPath, err)
			}
		default:
			fmt.Printf("Skipping unsupported tar entry type: %v for %s\n", header.Typeflag, header.Name)
		}
	}

	fmt.Printf("")
}
