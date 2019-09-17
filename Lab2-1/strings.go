package Lab2_1

func Reverse(s string) string {
	// Create rune array from s
	// Loop over runes and swap them
	// Add exception for combining characters
	// Return new string based on reversed rune array
	return s
}

func IsPalindrome(candidate string) bool {
	// Convert to lowercase
	// Remove punctuation and whitespace
	reverse := Reverse(candidate)
	return reverse == candidate
}
