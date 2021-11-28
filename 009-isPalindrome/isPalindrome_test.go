package isPalindrome

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Num    int
	Output bool
}

var (
	testCases = []TestCase{
		{
			Num:    121,
			Output: true,
		},
		{
			Num:    -121,
			Output: false,
		},
		{
			Num:    123,
			Output: false,
		},
		{
			Num:    10,
			Output: false,
		},
		{
			Num:    -101,
			Output: false,
		},
		{
			Num:    12321,
			Output: true,
		},
		{
			Num:    123454321,
			Output: true,
		},
	}
)

func Test_isPalindrome(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("scenario:%d", i), func(t *testing.T) {
			output := isPalindrome(testCase.Num)
			assert.Equal(t, output, testCase.Output)
		})
	}
}

func Test_isPalindromeBrute(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("scenario:%d", i), func(t *testing.T) {
			output := isPalindromeBrute(testCase.Num)
			assert.Equal(t, output, testCase.Output)
		})
	}
}
