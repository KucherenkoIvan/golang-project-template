package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func PrintASCIIBanner() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("unable to get wd: %v", err)
	}

	body, err := os.ReadFile(wd + "/assets/logo.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	for _, row := range strings.Split(string(body), "\n") {
		coloredRow := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 36, row)

		fmt.Println(coloredRow)
	}
}
