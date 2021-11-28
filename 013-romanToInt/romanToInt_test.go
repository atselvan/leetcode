package romanToInt

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Roman  string
	Output int
}

var (
	testCases = []TestCase{
		{
			Roman:  "III",
			Output: 3,
		},
		{
			Roman:  "IV",
			Output: 4,
		},
		{
			Roman:  "IX",
			Output: 9,
		},
		{
			Roman:  "LVIII",
			Output: 58,
		},
		{
			Roman:  "MCMXCIV",
			Output: 1994,
		},
		{
			Roman:  "DCXXI",
			Output: 621,
		},
	}
)

func Test_romanToInt(t *testing.T){
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("scenario:%d", i), func(t *testing.T) {
			output := romanToInt(testCase.Roman)
			assert.Equal(t, testCase.Output, output)
		})
	}
}

func Test_romanToIntBrute(t *testing.T){
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("scenario:%d", i), func(t *testing.T) {
			output := romanToIntBrute(testCase.Roman)
			assert.Equal(t, testCase.Output, output)
		})
	}
}