package leetcode

func moveZeroes(nums []int) {

	write := 0

	for _, num := range nums {
		if num != 0 {
			nums[write] = num
			write++
		}
	}

	for i := write; i < len(nums); i++ {
		nums[i] = 0
	}

}
