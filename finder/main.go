package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/loganstone/go_in_action/finder/matchers"
	"github.com/loganstone/go_in_action/finder/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func printUsageAndExit() {
	fmt.Printf("Usage: %s %s \n", filepath.Base(os.Args[0]), "<search term>")
	os.Exit(0)
}

func main() {
	if len(os.Args) != 2 {
		printUsageAndExit()
	}

	searchTerm := strings.TrimSpace(os.Args[1])
	if searchTerm == "" {
		printUsageAndExit()
	}

	search.Run(searchTerm)
}
