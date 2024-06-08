package main

import (
	"fmt"
)

func maxPalindrome(palindrome string, k int) string {
	s := []rune(palindrome)
	return maxPalindromeRecursive(s, 0, len(s)-1, k)
}

func maxPalindromeRecursive(s []rune, left, right, k int) string {
	if left >= right {
		return string(s)
	}

	if s[left] != s[right] {
		maxChar := max(s[left], s[right])
		s[left] = maxChar
		s[right] = maxChar
		k--
	}

	if k < 0 {
		return "-1"
	}

	return maxPalindromeRecursive(s, left+1, right-1, k)
}

func max(a, b rune) rune {
	if a > b {
		return a
	}
	return b
}

func main() {
	palindrome := "3943"
	k := 1
	result := maxPalindrome(palindrome, k)
	fmt.Println("Input : ", palindrome, " k : ", k)
	fmt.Println("Highest Palindrome:", result)
}
