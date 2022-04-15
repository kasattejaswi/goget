/*
Copyright © 2022 Tejaswi Kasat (kasattejasvi@gmail.com)

*/
package main

import (
	"flag"
	"github.com/kasattejaswi/goget/internal/cli"
	"os"
)

func main() {
	// Defining command structure
	cliStructure := cli.Structure{
		CliName: "goget",
		Args: []string{
			"[OPTIONS]",
		},
		FlagHeading: "Options:",
		FlagData: []cli.Flags{
			{
				Name:    "-h",
				Message: "Show help",
			},
			{
				Name:    "-u",
				Message: "Url of the file to be downloaded",
			},
			{
				Name:    "-t",
				Message: "Number of concurrent threads to be used",
			},
			{
				Name:    "-o",
				Message: "File path where file will be downloaded",
			},
			{
				Name:    "-n",
				Message: "Name of file with which it will be created",
			},
		},
		TabWidth: 15,
	}
	// Flag variables creation
	var url string
	var help bool
	var threads int
	var output string
	var name string

	// Flag definition
	flag.StringVar(&url, "u", "", "Url of the file to be downloaded")
	flag.BoolVar(&help, "h", false, "Show help")
	flag.IntVar(&threads, "t", 1, "Number of concurrent threads to be used")
	flag.StringVar(&output, "o", ".", "File path where file will be downloaded")
	flag.StringVar(&name, "n", "", "Name of file with which it will be created")

	// Flag parsing
	flag.Parse()

	// Printing commandline help
	if help {
		cliStructure.PrintHelp(os.Stdout)
	}

}
