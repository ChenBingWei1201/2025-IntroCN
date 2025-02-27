package main

import "fmt"
import "os"
import "bufio"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	inputName, outputName := "", ""
	fmt.Printf("Input filename: ")
	fmt.Scanf("%s", &inputName)
	fmt.Printf("Output filename: ")
	fmt.Scanf("%s", &outputName)

	input, inputErr := os.Open(inputName)
	check(inputErr)
	defer input.Close()

	output, outputErr := os.Create(outputName)
	check(outputErr)
	defer output.Close()

	scanner := bufio.NewScanner(input)
	writer := bufio.NewWriter(output)

	for lineNumber := 1; scanner.Scan(); lineNumber++ {
		line := scanner.Text()
		_, err := writer.WriteString(fmt.Sprintf("%d %s\n", lineNumber, line))
		check(err)
	}

	writer.Flush()
}