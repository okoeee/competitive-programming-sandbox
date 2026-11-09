package leetcode

func findMaxConsecutiveOnes(nums []int) int {
	counter := 0
	result := 0

	for _, num := range nums {
		if num == 1 {
			counter++
			result = max(counter, result)
		} else if num == 0 {
			result = max(counter, result)
			counter = 0
		}

	}

	return result

}
