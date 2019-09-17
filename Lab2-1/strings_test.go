package Lab2_1

import "testing"

func Test_reverse(t *testing.T) {
	//noinspection SpellCheckingInspection
	reverseMap := map[string]string {
		"": "",
		"a": "a",
		"ab": "ba",
		"abc": "cba",
		"abcd": "dcba",
		"Göre": "eröG",
		"Göre": "eröG",
	}
	for forward, reverse := range reverseMap {
		if forward != Reverse(reverse) {
			t.Errorf("Expected <%s> to be the reverse of <%s>", forward, reverse)
		}
	}
}

func Test_palindrome(t *testing.T) {
	palindromes := []string {
		"I",
		"bb",
		"civic",
		"Ada",
		"race car",
		"Is it crazy how saying sentences backwards creates backwards sentences saying how crazy it is?",
		"Go hang a salami, I'm a lasagna hog.",
		"दामाद",
		"लाल",
		"कनक",
		"रामाला भाला मारा", // Probably not possible
		"rör",
		"rör",
	}
	for _, palindrome := range palindromes {
		if !IsPalindrome(palindrome) {
			t.Errorf("Expected the following text to be a palindrome: \n%s", palindrome)
		}
		if IsPalindrome(palindrome + "abc") {
			t.Errorf("Expected the following text to not be a palindrome: \n%s", palindrome + "abc")
		}
	}
}
