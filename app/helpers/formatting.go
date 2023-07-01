package helpers

import (
	"strings"
)

// FormatName adalah fungsi untuk memformat nama menjadi huruf kapital di awal kata
func FormatName(name string) string {
	words := strings.Fields(name)
	for i, word := range words {
		words[i] = strings.Title(strings.ToLower(word))
	}
	return strings.Join(words, " ")
}
