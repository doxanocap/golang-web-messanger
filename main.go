package main

import (
	"fmt"
)

func main() {
	fmt.Println(secondsToRemoveOccurrences("0110101"))
	// 1011010s
}

func secondsToRemoveOccurrences(s string) int {
	c := 0
	for {
		if replace(s) == s {
			return c
		}
		s = replace(s)
		c++
	}
	return 0
}

func replace(s string) string {
	ans := ""
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && s[i] == '0' && s[i+1] == '1' {
			ans = ans + "10"
			i++
		} else {
			ans = ans + string(s[i])
		}
	}
	return ans
}
