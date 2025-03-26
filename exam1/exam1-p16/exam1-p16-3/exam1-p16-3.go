package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func shiftCharacter(c rune) rune {
	if c >= 'A' && c < 'Z' {
		return c + 1
	} else if c == 'Z' {
		return 'A'
	} else if c >= 'a' && c < 'z' {
		return c + 1
	} else if c == 'z' {
		return 'a'
	}
	return c
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
		line := scanner.Text()
		var shiftedLine strings.Builder
		for _, char := range line {
			shiftedLine.WriteRune(shiftCharacter(char))
		}
		fmt.Println(shiftedLine.String())
	}
}
