package twoSums

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Num []int
	Target int
	output []int
}

var (
	scenarios = []TestCase{
		{
			Num:    []int{2, 7, 11, 15},
			Target: 9,
			output: []int{0, 1},
		},
		{
			Num:    []int{3, 2, 4},
			Target: 6,
			output: []int{1, 2},
		},
		{
			Num:    []int{3, 3},
			Target: 6,
			output: []int{0, 1},
		},
		{
			Num:    []int{3, 3, 3},
			Target: 6,
			output: []int{0, 1},
		},
	}
)

func Test_twoSum(t *testing.T) {
	for i, scenario := range scenarios {
		t.Run(fmt.Sprintf("scenario:%d", i + 1), func(t *testing.T) {
			result := twoSum(scenario.Num, scenario.Target)
			assert.Equal(t, scenario.output, result)
		})
	}
}

func Test_twoSumBrute(t *testing.T) {
	for i, scenario := range scenarios {
		t.Run(fmt.Sprintf("scenario:%d", i + 1), func(t *testing.T) {
			result := twoSumBrute(scenario.Num, scenario.Target)
			assert.Equal(t, scenario.output, result)
		})
	}
}
