func pivotIndex(nums []int) int {
	leftSum := 0
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}
	for i, num := range nums {
		if leftSum == totalSum-leftSum-num {
			return i
		}
		leftSum += num
	}
	return -1
}