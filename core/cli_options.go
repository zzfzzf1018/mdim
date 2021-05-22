package core

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// Command-line options.
type CliOptions struct {
	AbsDocFolder string
	AbsImgFolder string
	DoRelPathFix bool
	DoImgDel     bool
}

// Deal with cli params.
func GetOptions() *CliOptions {
	var help bool
	CliParams := &CliOptions{}

	flag.Usage = usage
	flag.BoolVar(&help, "h", false, "Show this help.")
	flag.BoolVar(&CliParams.DoRelPathFix, "f", false, "Set the option to fix image relative paths of markdown documents, otherwise print the paths only.")
	flag.BoolVar(&CliParams.DoImgDel, "d", false, "Set the option to delete no reference images, otherwise print the paths only.")
	flag.StringVar(&CliParams.AbsDocFolder, "m", "", "Must be not empty. The folder markdown documents save in")
	flag.StringVar(&CliParams.AbsImgFolder, "i", "", "Must be not empty. The folder images save in")

	flag.Parse()

	// Show usage and then exit directly.
	if help {
		flag.Usage()
		os.Exit(0)
	}

	if CliParams.AbsImgFolder == "" || CliParams.AbsDocFolder == "" {
		flag.Usage()
		os.Exit(1)
	}

	var err error
	if !filepath.IsAbs(CliParams.AbsImgFolder) {
		CliParams.AbsImgFolder, err = filepath.Abs(CliParams.AbsImgFolder)
		if err != nil {
			fmt.Printf("Cannot get the abs path of img folder\n%s\n%s", CliParams.AbsImgFolder, err.Error())
			os.Exit(2)
		}
	}

	if !filepath.IsAbs(CliParams.AbsDocFolder) {
		CliParams.AbsDocFolder, err = filepath.Abs(CliParams.AbsDocFolder)
		if err != nil {
			fmt.Printf("cannot get the abs path of doc folder\n%s\n%s", CliParams.AbsDocFolder, err.Error())
			os.Exit(3)
		}
	}

	return CliParams
}

func usage() {
	fmt.Fprintf(os.Stderr, `mdim - Markdown Images Maintainer

Description: The tool will help to maintain the image relative paths of markdown files and cleanup no reference images.

Github: https://github.com/bunnier/mdim

Usage: mdim [-h] [-d] [-f] [-i imageFolder] [-m markdownFolder] 

Options:
`)
	flag.PrintDefaults()
}
