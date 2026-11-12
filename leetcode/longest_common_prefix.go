package leetcode

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	result := ""

	minStrLen := len(strs[0])
	for _, str := range strs {
		minStrLen = min(minStrLen, len(str))
	}

	for i := range minStrLen {
		initStr := strs[0][i]
		isAllSame := true
		for _, str := range strs {
			isAllSame = isAllSame && (initStr == str[i])
		}

		if !isAllSame {
			return result
		}
		result += string(initStr)
	}

	return result

}
