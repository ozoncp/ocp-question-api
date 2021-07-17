package utils

func contains(source []int, needle int) bool {
	for _, value := range source {
		if value == needle {
			return true
		}
	}

	return false
}

func Filter(source []int) []int {
	hardcodeSlice := []int{1, 2, 3}

	result := make([]int, 0)
	for _, v := range source {
		if !contains(hardcodeSlice, v) {
			result = append(result, v)
		}
	}

	return result
}
