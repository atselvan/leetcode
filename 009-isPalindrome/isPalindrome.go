package isPalindrome

import (
	"strconv"
	"strings"
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
		number = number / 10
	}
	return x == reverse
}

func isPalindromeBrute(x int) bool {
	count := 0
	nums := strings.Split(strconv.Itoa(x), "")

	for i, num := range nums {
		if num == nums[len(nums)-(i+1)] {
			count++
		}
	}
	if count == len(nums) {
		return true
	}
	return false
}
