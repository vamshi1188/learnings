package main

import (
	"fmt"
	"unicode"
)

func main() {

	// Recusrion(1, 1)

	s := "A man, a plan, a canal: Panama"

	fmt.Println(IsPalinfrome(s))
}

func Recusrion(n int, current int) {
	if current > n {
		return
	}
	fmt.Println(n)
	Recusrion(n, current+1)
}

// CLEAN + CHECK PALINDROME
func IsPalinfrome(s string) bool {

	// 1) Clean string: keep only letters/digits, convert to lowercase
	var cleaned []rune
	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			cleaned = append(cleaned, unicode.ToLower(ch))
		}
	}

	// 2) Convert cleaned string to slice of runes
	conv := cleaned
	lngth := len(conv)

	// 3) Create reversed version
	var reversed []rune
	for i := lngth - 1; i >= 0; i-- {
		reversed = append(reversed, conv[i])
	}

	// 4) Compare
	for i := 0; i < lngth; i++ {
		if conv[i] != reversed[i] {
			return false
		}
	}
	return true
}
