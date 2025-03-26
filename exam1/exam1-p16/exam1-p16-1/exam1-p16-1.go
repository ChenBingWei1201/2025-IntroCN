package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	inputName := ""
	fmt.Printf("Input filename: ")
	fmt.Scanf("%s", &inputName)

	input, inputErr := os.Open(inputName)
	check(inputErr)
	defer input.Close()

	fileInfo, fileInfoErr := input.Stat()
	check(fileInfoErr)
	fileSize := fileInfo.Size()
	fmt.Printf("%d\n", fileSize)
}
