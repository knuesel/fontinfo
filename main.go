package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

import "github.com/ConradIrwin/font/sfnt"

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <font file>", os.Args[0])
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	font, err := sfnt.Parse(file)
	if err != nil {
		log.Fatal(err)
	}

	name, err := font.NameTable()
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range name.List() {
		// fmt.Printf("% -16v\t%v\n", entry.Label()+":", entry.String())
		if entry.Label() == "Version" {
			version := strings.TrimPrefix(entry.String(), "Version ")
			if version != "" {
				fmt.Println("Version:", version)
				break
			}
		}
	}
}
