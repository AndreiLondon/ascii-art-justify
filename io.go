package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var usage = `Usage: go run . [OPTION] [STRING] [BANNER]

Example: go run . --align=right something standard`

// svar := string

func getArguments() (string, string, string) {
	// flag.StringVar (&svar, "svar", "bar", "a string var")
	fileName := flag.String("align", "type", "it returns a message to display with different alignment")
	flag.Parse()
	args := os.Args[1:]
	// go run . empty or --output=result.txt "hello" standard shadow
	if len(args) == 0 || len(args) > 3 {
		fmt.Println(usage)
		os.Exit(0)
	}
	if len(args) == 3 && strings.HasPrefix(args[0], "--align=") {
		//return *fileName, args[1], args[2]

	}
	if len(args) == 2 && strings.HasPrefix(args[0], "--align=") {
		//return *fileName, args[1], "standard"
	}
	if len(args) == 2 && !strings.HasPrefix(args[0], "--align=") {
		//return "result.txt", args[0], args[1]
	}
	if len(args) == 1 {
		//return "result.txt", args[0], "standard"
	}

	// set a flag to specify the message to display
	// 1st is the name of the flag that user needs to type --align
	// 2nd is a default value, if 1st value  is missing then what should the default value be. It would be displayed.
	// 3d is the help message. What this flag does
	// fmt.Print(*fileName)
	fmt.Println(usage)
	os.Exit(0)

	return "", "", ""
}

func readFile(s string) string {
	data, err := os.ReadFile(s)
	if err != nil {
		log.Fatal("invalid file. Cannot read " + s)
	}
	return string(data)
}

// writeToFile writes the completed string to a file called 'result.txt'
func writeToFile(fileName string, data []byte) {
	err := os.WriteFile(fileName, data, 0644)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
}
