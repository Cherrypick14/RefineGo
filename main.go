package main

import (
	"fmt"
	"os"
	"strings"
	"reloaded/functions"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Usage: <sample.txt> <result.txt>")
		return
	}

	inputtext := args[0]

	content, err := os.ReadFile(inputtext)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	words := strings.Split(string(content), " ")

	contents := functions.Capitalize(strings.Split(functions.Upper(words), " "))


	output := os.WriteFile(args[1], []byte(contents),0644)

	if output != nil {
		fmt.Println("err", output)
	}
}


