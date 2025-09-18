package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// FINALLY CLEAN DOWNLOADING!!!1!!!1!!

func download(pkg string, url string) error {
	packagenm := "/bin/arfpkg/temp/" + pkg
	err := downloadFile(url, packagenm)
	if err != nil {
		fmt.Printf("ERROR FETCHING!!!!\n exiting...")
	} else {
		fmt.Printf("")
	}
	return err
}

func downloadFile(url string, filePath string) error {
	output, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer func(output *os.File) {
		err := output.Close()
		if err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}(output)

	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error making HTTP request: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("Error closing body: %v\n", err)
		}
	}(response.Body)

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status code: %d %s", response.StatusCode, response.Status)
	}

	_, err = io.Copy(output, response.Body)
	if err != nil {
		return fmt.Errorf("error copying data: %w", err)
	}

	return nil
}
