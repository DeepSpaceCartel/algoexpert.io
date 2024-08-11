// This file is initialized with a code version of this
// question's sample test case. Feel free to add, edit,
// or remove test cases in this file as you see fit!

package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	expected := []int{-1, 11}
	output := TwoNumberSum([]int{3, 5, -4, 8, 11, 1, -1, 6}, 10)
	require.ElementsMatch(t, expected, output)
}

func BenchmarkCase1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwoNumberSum([]int{3, 5, -4, 8, 11, 1, -1, 6}, 10)
	}
}
