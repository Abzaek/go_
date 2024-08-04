package Str

import "strings"

func IsPalindrome() bool {

	lst := WordCounter(false)

	l := 0
	r := len(lst) - 1

	for l < r {
		if !strings.EqualFold(lst[l], lst[r]) {
			return false
		}
		l += 1
		r -= 1
	}
	return true
}
