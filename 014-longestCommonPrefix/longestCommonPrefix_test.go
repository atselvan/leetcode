package longestCommonPrefix

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Strs   []string
	Output string
}

var (
	testCases = []TestCase{
		{
			Strs:   []string{"flower", "flow", "flight"},
			Output: "fl",
		},
		{
			Strs:   []string{"dog", "racecar", "car"},
			Output: "",
		},
		{
			Strs:   []string{"a"},
			Output: "a",
		},
		{
			Strs:   []string{"test"},
			Output: "test",
		},
		{
			Strs:   []string{"", "racecar", "car"},
			Output: "",
		},
	}
)

func Test_longestCommonPrefix(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("scenario:%d", i), func(t *testing.T) {
			output := longestCommonPrefix(testCase.Strs)
			assert.Equal(t, testCase.Output, output)
		})
	}
}

func Test_longestCommonPrefixBrute(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("scenario:%d", i), func(t *testing.T) {
			output := longestCommonPrefixBrute(testCase.Strs)
			assert.Equal(t, testCase.Output, output)
		})
	}
}
