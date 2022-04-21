package asciiart

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const ASCII_START = ' '
const ASCII_END = '~'
const FONT_HEIGHT = 8

func GetFileLines(symbols []byte) []string {
	symbol := string(symbols)
	line := strings.Split(strings.ReplaceAll(symbol, "\r\n", "\n"), "\n")
	return line //returns 855 lines
}

func StoreInDictionary(line []string) map[rune][]string {
	q := 1
	m := make(map[rune][]string)
	for i := 0; i < 95; i++ {
		m[rune(i+32)] = line[q : q+FONT_HEIGHT]
		q = q + FONT_HEIGHT + 1
	}
	return m
}

func IsValidText(s string) bool {
	for _, ch := range s {
		if ch != 10 && ch != 13 && (ch < ASCII_START || ch > ASCII_END) {
			return false
		}
	}
	return true
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func PrintSymbols(arg, banner string) (string, error) {
	symbols, err := ioutil.ReadFile("ascii-art/" + banner + ".txt")
	if err != nil {
		return "", fmt.Errorf("an error occured while accessing/reading the %q banner", banner)
	}
	line := GetFileLines(symbols)
	m := StoreInDictionary(line) //easch symbol is saved as a key, f.ex., 'A' is the key of the A symbol
	arg = strings.ReplaceAll(arg, "\r\n", "\n")
	userText := strings.Split(arg, "\n")
	symbolText, count := "", 1
	for t := 0; t < len(userText); t++ {
		if userText[t] == "" {
			if count < len(userText) {
				symbolText += "\n"
				count++
			}
			continue
		}
		for i := 0; i < FONT_HEIGHT; i++ {
			for _, ch := range userText[t] {
				symbolText += m[ch][i]
			}
			symbolText += "\n"
		}
	}
	return symbolText, nil
}
