/*
Copyright © 2022 Tejaswi Kasat (kasattejasvi@gmail.com)

*/
package main

import (
	"flag"
	"fmt"
	"github.com/kasattejaswi/goget/internal/cli"
	"github.com/kasattejaswi/goget/internal/downloader"
	"os"
	"strings"
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

	if url == "" {
		fmt.Printf("Error: Please pass a URL to download a file\n")
		cliStructure.PrintHelp(os.Stdout)
	}

	if name == "" {
		urlP := strings.Split(strings.ReplaceAll(url, " ", ""), "/")
		for i := len(urlP) - 1; i >= 0; i-- {
			if urlP[i] != "" {
				name = urlP[i]
				break
			}
		}
	}
	// Building download options
	downloadOptions := downloader.DownloadOptions{
		Url:      url,
		Threads:  threads,
		Output:   output,
		FileName: name,
	}
	err := downloadOptions.Download()
	if err != nil {
		fmt.Println(err)
		return
	}
}
