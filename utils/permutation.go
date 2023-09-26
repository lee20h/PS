package utils

func nextPermutation(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}

	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i == -1 {
		return false
	}

	j := n - 1
	for nums[j] <= nums[i] {
		j--
	}
	nums[i], nums[j] = nums[j], nums[i]

	reverse(nums, i+1, n-1)
	return true
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}