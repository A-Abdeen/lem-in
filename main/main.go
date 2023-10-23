package main

import (
	"fmt"
	lemin "lemin/adnanok"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage case requires txt file")
		os.Exit(0)
	}
	inputFile := os.Args[1]
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	// fmt.Println(string(data))
	dataStr := string(data)
	x := lemin.Validate(dataStr)
	fmt.Println(x)
}
