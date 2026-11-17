package leetcode

import (
	"slices"
	"strings"
)

func isAnagram(s string, t string) bool {
	ss := strings.Split(s, "")
	ts := strings.Split(t, "")

	slices.Sort(ss)
	slices.Sort(ts)

	return slices.Equal(ss, ts)
}
