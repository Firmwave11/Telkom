package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func main() {

	convertCmd := flag.String("t", "json", "a string with value json(default) or text example [path] -t [json/text]")
	locCmd := flag.String("o", "", "a string with value example [path] -t [json/text] -o [path] || [path] -o [path]")
	var locSaveFile string
	if len(os.Args) < 2 {
		fmt.Println("expected '-t', '-o' or '-h' subcommands")
		fmt.Println(*convertCmd)
		fmt.Println(*locCmd)
		os.Exit(1)
	}
	flag.Parse()

	if len(os.Args) > 2 {
		checkFileExist(os.Args[1])
		//look at the 3rd argument's value
		switch os.Args[2] {
		case "-t": // if its the '-t' command
			if os.Args[3] == "json" {
				locSaveFile = "error.json"
				if len(os.Args) > 4 {
					if os.Args[5] != "" {
						locSaveFile = os.Args[5]
					}
				}

				HandleConvertToJson(os.Args[1], locSaveFile)
			} else {
				locSaveFile = "error.txt"
				if len(os.Args) > 4 {
					if os.Args[5] != "" {
						locSaveFile = os.Args[5]
					}
				}
				HandleConvertToText(os.Args[1], locSaveFile)
			}
		case "-o": // if its the '-O' command
			for k, arg := range os.Args {
				if arg == "-o" {
					if len(os.Args)-2 != k {
						fmt.Println("wrong value")
						os.Exit(1)
					}
					locSaveFile = os.Args[k+1]
					break
				}
			}
			HandleConvertToText(os.Args[1], locSaveFile)
		default: // if we don't understand the input
		}
	} else {
		locSaveFile = "error.txt"
		HandleConvertToText(os.Args[1], locSaveFile)
	}

}

func checkFileExist(loc string) {
	if _, err := os.Stat(loc); errors.Is(err, os.ErrNotExist) {
		fmt.Println("file doesn't exist")
		os.Exit(1)
	}
}
