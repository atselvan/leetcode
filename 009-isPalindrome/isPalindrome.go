package isPalindrome

import (
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	number := x
	reverse := 0
	for number != 0 {
		reverse = reverse * 10
		reverse = reverse + number%10
		number /= 10
	}
	return x == reverse
}

func isPalindromeBrute(x int) bool {
	nums := strconv.Itoa(x)
	for i, _ := range nums {
		if i < len(nums)/2 {
			if nums[i] != nums[len(nums)-(i+1)] {
				return false
			}
		}
	}
	return true
}
