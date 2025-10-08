package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGenerateRandomElements tests the random data generation function
func TestGenerateRandomElements(t *testing.T) {
	testCases := []struct {
		name string
		size int
	}{
		{"zero size", 0},
		{"small size", 10},
		{"medium size", 1000},
		{"large size", 100000},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := generateRandomElements(tc.size)

			assert.Len(t, res, tc.size, "Should return slice of correct length")

			if tc.size > 0 {
				for _, val := range res {
					assert.GreaterOrEqual(t, val, 0, "Element should be positive")
				}
			}
		})
	}
}

// TestMaximum tests the sequential maximum finding function
func TestMaximum(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{"empty slice", []int{}, 0},
		{"single element", []int{42}, 42},
		{"max at beginning", []int{100, 1, 2, 3}, 100},
		{"max at end", []int{1, 2, 3, 100}, 100},
		{"max in middle", []int{1, 100, 2, 3}, 100},
		{"all same", []int{5, 5, 5, 5}, 5},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := maximum(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

// TestMaxChunks tests the parallel maximum finding function
func TestMaxChunks(t *testing.T) {
	withTailElements := make([]int, 10)
	for i := range withTailElements {
		withTailElements[i] = i
	}
	withTailElements[9] = 100

	allSameElements := make([]int, 100)
	for i := range allSameElements {
		allSameElements[i] = 42
	}

	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{"empty slice", []int{}, 0},
		{"single element", []int{42}, 42},
		{"less elements than chunks", []int{1, 5, 3}, 5},
		{"exact number of chunks", []int{1, 2, 3, 4, 5, 6, 7, 8}, 8},
		{"with tail elements", withTailElements, 100},
		{"all elements same", allSameElements, 42},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := maxChunks(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}
