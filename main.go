package main

import "fmt"

func main() {
	fmt.Println(IsPalindrome("rcar"))
}

func IsPalindrome(str string) bool {
	runes := []rune(str)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}
	return true
}
