package solutions

func RunningSum(nums []int) []int {
	count := 0

	resultSlice := make([]int, len(nums))

	for i := range nums {
		count += nums[i]
		resultSlice[i] = count
	}

	return resultSlice

}
