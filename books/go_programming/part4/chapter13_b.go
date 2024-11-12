package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Using pipes, stdin and stdout to apply a Rot13 encoding to a file")
	fmt.Println()

	stat, _ := os.Stdin.Stat()

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		processStdin()
	} else {
		processFileOrInput()
	}
}

func rot13(s string) string {
	result := make([]byte, len(s))

	for i := 0; i < len(s); i++ {
		char := s[i]

		switch {
		case char >= 'a' && char <= 'z':
			result[i] = 'a' + (char-'a'+13)%26
		case char >= 'A' && char <= 'Z':
			result[i] = 'A' + (char-'A'+13)%26
		default:
			result[i] = char
		}
	}

	return string(result)
}

func processStdin() {
	reader := bufio.NewReader(os.Stdin)

	for {
		input, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error reading from stdin")
			return
		}

		encoded := rot13(input)
		fmt.Print(encoded)
	}
}


func processFileOrInput() {
	var inputReader io.Reader
	// Check if the input is a file or not
	if len(os.Args) > 1 {
		file, err := os.Open(os.Args[1])

		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}

		defer file.Close()

		inputReader = file
	} else {
		// No file path provided, read from stdin
		fmt.Print("Enter text to encode: ")
		inputReader = os.Stdin
	}

	// Process input and apply rot13 encoding

	scanner := bufio.NewScanner(inputReader)
	for scanner.Scan() {
		encoded := rot13(scanner.Text())
		fmt.Println(encoded)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}

