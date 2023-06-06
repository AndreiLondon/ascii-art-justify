package main

import (
	"strings"
)

func parseTemplate(s string) map[rune][][]rune {
	myMap := make(map[rune][][]rune)
	ascii := 32
	data := [][]rune{}
	rowData := []rune{}

	for _, char := range s {
		if char != 10 && char != 13 {
			// We're adding a char to this rowData if the char != 10
			rowData = append(rowData, char)
		}
		// Adding the rowData to the two dimen. array called data
		if char == 10 || char == 13 {
			// We only add rowData to data if rowData is not empty
			if len(rowData) > 1 {
				data = append(data, rowData)
				rowData = []rune{}

			}
			if len(data) == 8 {
				myMap[rune(ascii)] = data
				ascii++
				data = [][]rune{}
			}

		}

	}
	return myMap
}

func printMessage(s string, myMap map[rune][][]rune) string {
	text := ""
	split := strings.Split(s, "\\n")

	// We need this check when we pass \n twice in a row, so we don't need to print an empty string
	for _, v := range split {
		if v == "" {
			text += "\n"
			continue
		}
		// 1st loop is for counting the lines
		for i := 0; i < 8; i++ {
			// 2nd loop is for grabbing each char from the string
			for _, ch := range v {
				// do whatever you want to handle errors - this means this wasn't a string
				if data, ok := myMap[ch]; ok {
					rowData := data[i]
					for _, value := range rowData {
						text += string(value)
					}
				}
			}
			text += "\n"
		}
	}
	return text
}
