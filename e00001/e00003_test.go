package e00001

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Given an integer n, return a string array answert (1-indexed)
// where
// answer[i] == "FizzBuzz" if i is divisible by 3 and 5
// answert[i] == "Fizz" if i divisible by 3
// answert[i] == "Buzz" i i is divisible by 5
// answer[i] == i (as a string) if non of the above conditions are true

func getNewElement(index int) string {
	if index % 15 == 0 {
		return "FizzBuzz" 
	}
	if index % 3 == 0 {
		return "Fizz" 
	}
	if index % 5 == 0 {
		return "Buzz" 
	}
	return strconv.Itoa(index)
}

func fizzBuzz(n int) []string {
	out := make([]string, 0, n)
	for i := 0; i < n; i++ {
		out = append(out, getNewElement(i+1) )
	}
	return out
}

func Test_Suite(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		r := fizzBuzz(0)
		assert.Equal(t, 0, len(r))
	})

	t.Run("1", func(t *testing.T) {
		r := fizzBuzz(1)
		assert.Equal(t, 1, len(r))
		assert.Equal(t, []string{"1"}, r)
	})

	t.Run("2", func(t *testing.T) {
		r := fizzBuzz(2)
		assert.Equal(t, 2, len(r))
		assert.Equal(t, []string{"1", "2"}, r)
	})

	t.Run("3", func(t *testing.T) {
		r := fizzBuzz(3)
		assert.Equal(t, 3, len(r))
		assert.Equal(t, []string{"1", "2", "Fizz"}, r)
	})

	t.Run("5", func(t *testing.T) {
		r := fizzBuzz(5)
		assert.Equal(t, 5, len(r))
		assert.Equal(t, []string{"1", "2", "Fizz", "4", "Buzz"}, r)
	})
}

func Test_getNewElement(t *testing.T) {
	t.Run("can get FizzBuzz on 3*5", func(t *testing.T){
		actual := getNewElement(3*5)
		assert.Equal(t, "FizzBuzz", actual)
	})
}

