package leetcode

func containsDuplicate(nums []int) bool {
	set := make(map[int]struct{})
	for _, num := range nums {
		set[num] = struct{}{}
	}

	return !(len(nums) == len(set))
}
