package e00001

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 1342. Number of Steps to Reduca a Numer to Zero
// Given an integer num, return the number of steps to reduce it to Zero
// In one step, if the current number is even, you have to devide it by 2,
// otherwise, you have to subtrct 1 from it

func numberOfSteps(num int) int {
	steps := 0
	for steps = 0; num != 0; steps++ {
		if num%2 == 0 {
			num = num / 2
		} else {
			num = num - 1
		}
	}
	return steps
}

func Test_NumberOfSteps(t *testing.T) {
	assert.Equal(t, 0, numberOfSteps(0))
	assert.Equal(t, 1, numberOfSteps(1))
	assert.Equal(t, 2, numberOfSteps(2))
	assert.Equal(t, 3, numberOfSteps(3))
	assert.Equal(t, 4, numberOfSteps(8))
	assert.Equal(t, 6, numberOfSteps(14))
}
