package leetcode

import "strings"

func lengthOfLastWord(s string) int {

	list := strings.Split(strings.TrimSpace(s), " ")

	return len(list[len(list)-1])

}
