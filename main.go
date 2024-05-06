package main

import (
	"fmt"
	"strings"
	"unicode"
	// "os"
	// "strings"
	//  "reloaded/utils"
)

func main() {
	// grab arguments from the command line
	//    args := os.Args[1:]

	// throw error if args is not 3
	//    if len(args) != 3 {
	// 	 fmt.Println("Usage: <sample> <result>")
	//    }
	// Read file contents
	//    inputText, err := os.ReadFile(args[0])

	// throw error reading file
	//    if err != nil {
	// 	  fmt.Println("Error reading file")
	//    }
	//split text and separate with spaces
	//    words := strings.Split(string(inputText), " ")
	word := []string{"it(cap) should happen"}
	capitalized := Capitalize(word)
	fmt.Println(capitalized)

	// Write to file
	//    os.WriteFile(args[1], []byte(word), 0644)

}

func Capitalize(words []string) string {
	for i := 0; i < len(words); i++ {
		word := words[i]
		if word == "cap" {
			prevWord := words[i-1]
			capitalizeWord := CapitalizeFirstWord(prevWord)
			words[i-1] = capitalizeWord
			words = append(words[:i], words[i+1:]...)
		}
	}
	return strings.Join(words, "")
}

func CapitalizeFirstWord(word string) string {
	if len(word) == 0 {
		return word
	}
	runes := []rune(word)
	runes[0] = unicode.ToUpper(runes[0])

	return string(runes)
}
