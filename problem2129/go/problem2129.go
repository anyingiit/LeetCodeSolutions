package main

import (
	"strings"
)

func capitalizeTitle(title string) string {
	letters := strings.Split(title, " ")

	for i := range letters {
		letters[i] = strings.ToLower(letters[i])

		if len(letters[i]) > 2 {
			// if it has chinese, can use
			// string([]rune(str)[0])
			firstLetter := strings.ToUpper(string(letters[i][0]))

			letters[i] = firstLetter + letters[i][1:]
		}
	}

	return strings.Join(letters, " ")
}
