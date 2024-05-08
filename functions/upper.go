package functions

import (
	"fmt"
	"strconv"
	"strings"
)

// Uppercase words in a slice based on "(up)" or "(up," markers and their following numbers.

func Upper(words []string) string {
	// Iterate through each word in the slice
	for i, word := range words {
		switch word {
		case "(up)":
			if i > 0 {
				prevWord := words[i-1]
				prevWord = strings.ToUpper(string(prevWord[0])) + prevWord[1:]
				words[i-1] = prevWord
				words = append(words[:i], words[i+1:]...)
			}
		case "(up,":
			b := strings.Trim(string(words[i+1]), words[i+1][1:])
			number, err := strconv.Atoi(b)
			if err != nil {
				fmt.Println("Error converting number")
			}
			for j := 1; j <= number; j++ {
				if i-j >= 0 { // Ensure index does not go out of bounds
					prevWord := words[i-j]
					prevWord = strings.ToUpper(string(prevWord[0])) + prevWord[1:]
					words[i-j] = prevWord
				}
			}
			words = append(words[:i], words[i+2:]...)
		}
	}
	return strings.Join(words, " ")
}
