package main

import (
	"fmt"
	"os"
	"strconv"
)

func HandleArgs() int {
	args := os.Args

	if len(args) <= 1 {
		fmt.Println("You must provide a file name.")
		os.Exit(1)
	}

	lim := 10 // default limit of results returned. Can be specified by user

	// Check if optional param was passed in
	if len(args) == 4 {
		flag := args[1]
		value, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Printf("Error parsing result limitor. Make sure this is a valid number.\nEx: ./program -n 5 sample.txt")
			os.Exit(1)
		}

		if flag != `-n` {
			fmt.Println("Incorrect params passed in.\nEx: ./program -n 5 sample.txt")
			os.Exit(1)
		}

		lim = value
	}

	return lim
}
