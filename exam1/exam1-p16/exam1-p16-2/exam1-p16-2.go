package main

import (
	"bufio"
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

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
