package main

import "fmt"

func main() {
	s := "abcabcbb"

	result := lengthOfLongestSubstring(s)

	fmt.Println(result)
}

func lengthOfLongestSubstring(s string) int {

	var left, maxSize int
	window := make(map[byte]bool)

	for right := 0; right < len(s); right++ {
		// expanding loop
		for window[s[right]] {
			//? shrinking loop
			delete(window, s[left])
			left++

		}

		window[s[right]] = true
		if right-left+1 > maxSize {
			maxSize = right - left + 1

		}
	}
	return maxSize
}
