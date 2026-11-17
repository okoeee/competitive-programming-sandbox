package leetcode

func twoSum(nums []int, target int) []int {

	seen := make(map[int]int)

	for i, num := range nums {
		need := target - num

		if j, ok := seen[need]; ok {
			return []int{i, j}
		}

		seen[num] = i
	}

	return []int{}

}
