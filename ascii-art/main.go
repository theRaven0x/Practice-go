package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println ("Invalid number of arguments")
	}

	input := os.Args[1]
	bannerName := os.Args[2]

	// if !validBanners[bannerName] {
	// 	fmt.Println ("Invalid banner use: standard, shadow and thinkertoy")
	// }

	fileName := bannerName + ".txt"

	data, err := os.ReadFile(fileName) 
		if err != nil {
			fmt.Println("error reading file")
		}

	lines := strings.Split(strings.ReplaceAll(string(data), "\r", ""), "\n")
	store := strings.Split(input, "\\n")

	for _, word := range store {
		if word == "" {
			fmt.Print()
			continue
		}

		for row := 1; row < 9; row++ {
			for _, char := range word {
				ascii := int(char)

				if ascii < 32 || ascii > 126 {
					continue
				}

				value := (ascii-32) * 9 
				fmt.Print(lines[value+row])
			}
			fmt.Println()
		}
	}
}

// var validBanners = map[string] bool {
// 	"standard": true,
// 	"thinkertoy": true,
// 	"shadow": true,
// }