package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func init() {
	if len(os.Args) != 2 {
		printUsageAndExit()
	}
}

func printUsageAndExit() {
	fmt.Printf("Usage: %s %s \n", filepath.Base(os.Args[0]), "<url>")
	os.Exit(-1)
}

func main() {
	r, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(os.Stdout, r.Body)
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}
}
