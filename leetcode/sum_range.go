package leetcode

type NumArray struct {
	prefix []int
}

func Constructor(nums []int) NumArray {
	pre := make([]int, 0, 0)

	acc := 0
	for _, num := range nums {
		pre = append(pre, acc+num)
		acc += num
	}

	return NumArray{prefix: append([]int{0}, pre...)}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefix[right+1] - this.prefix[left]
}
