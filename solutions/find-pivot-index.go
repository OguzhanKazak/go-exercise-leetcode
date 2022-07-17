package main

func pivotIndex(nums []int) int {
	leftSum := 0
	rightSum := 0
	for i := range nums { //pivot index
		for k := range nums {
			if i > k {
				leftSum += nums[k]
			}
			if i < k {
				rightSum += nums[k]
			}
		}

		if leftSum == rightSum {
			return i
		}
		leftSum = 0
		rightSum = 0
	}

	return -1

}

func pivotIndexBetterPerformance(nums []int) int {

	totalSum := 0
	leftSum := 0

	for i := range nums {
		totalSum += nums[i]
	}

	for i := range nums {
		leftSum += nums[i]
		if totalSum-leftSum == leftSum-nums[i] {
			return i
		}
	}

	return -1

}
