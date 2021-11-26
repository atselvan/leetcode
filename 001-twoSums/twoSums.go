package twoSums

func twoSum(nums []int, target int) []int {
	var output []int
	numMap := make(map[int]int)
	for i, num := range nums {
		if val, ok := numMap[num]; ok {
			return append(output, val, i)
		} else {
			numMap[target - num] = i
		}
	}
	return output
}

func twoSumBrute(nums []int, target int) []int {
	var output []int
	for i, n1 := range nums {
		result := target - n1
		for j, n2 := range nums {
			if result == n2 && i != j {
				return append(output, i, j)
			}
		}
	}
	return output
}
