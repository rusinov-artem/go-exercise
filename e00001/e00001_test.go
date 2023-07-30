package e00001

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 1480. Running Sum of 1d Array

// Given an array nums. We define a running sum of an array as
// runningSum[i] = sum(nums[0]…nums[i]).
// Return the running sum of nums.

func RunningSum(inp []int64) []int64 {
	var sum int64 = 0
	runningSum := make([]int64, len(inp))
	for i, v := range inp {
		sum = sum + v
		runningSum[i] = sum
	}
	return runningSum
}

func RunningSumImproved(inp []int64) []int64 {
	for i := 1; i < len(inp); i++ {
		inp[i] = inp[i-1] + inp[i]
	}
	return inp
}

func Test_RunningSum(t *testing.T) {
	algos := map[string]func([]int64) []int64{
		"RunningSum":         RunningSum,
		"RunningSumImproved": RunningSumImproved,
	}

	tests := []struct {
		input    []int64
		expected []int64
	}{
		{
			input:    []int64{1, 1, 1, 1, 1},
			expected: []int64{1, 2, 3, 4, 5},
		},
	}

	for algName, alg := range algos {
		for _, testCase := range tests {
			assertSlicesEq(t, testCase.expected, alg(testCase.input), algName)
		}
	}
}

func assertSlicesEq(t *testing.T, exp, act []int64, msg ...interface{}) {
	assert.True(t, reflect.DeepEqual(exp, act), msg)
}
