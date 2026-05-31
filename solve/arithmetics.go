package solve

func Sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

func Difference(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	total := nums[0]

	for _, num := range nums[1:] {
		total -= num
	}

	return total
}
