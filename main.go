package main

import (
	"fmt"
	"strings"
	"os"
	
)

func main() {
	// grab arguments from the command line
	   args := os.Args[1:]

	   inputtext := args[0]

	   words, err := os.ReadFile(inputtext)
	   if err != nil {
		  fmt.Println("Error reading file")
		  return
	   }
	   content := strings.Split(string(words), " ")
	   capitalized := CapitalizeAll(content)
	   fmt.Println(capitalized)
	// Write to file
	//    os.WriteFile(args[1], []byte(word), 0644)

}

func CapitalizeAll(content []string) string {
	for i := 0; i < len(content); i++ {
		if content[i] == "(cap)" && i > 0 {
			prevWord := content[i-1]
			if len(prevWord) > 0 {
				// Capitalize the first letter of the previous word
				firstChar := strings.ToUpper(string(prevWord[0]))
				prevWord = firstChar + prevWord[1:]
				content[i-1] = prevWord
			}
		}
	}
	return strings.Join(content, " ")
}
