package main

import "fmt"
import "os"
import "bufio"
import "io"

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

	reader := bufio.NewReader(input)
	writer := bufio.NewWriter(output)

  for lineNumber := 1; ; lineNumber++ {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			check(err)
		}

		_, writeErr := writer.WriteString(fmt.Sprintf("%d %s", lineNumber, line))
		check(writeErr)

		if err == io.EOF {
			break
		}
	}

	writer.Flush()
}