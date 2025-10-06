package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	t.Run("zero size", func(t *testing.T) {
		result := generateRandomElements(0)
		assert.Empty(t, result, "Should return empty slice for size 0")
	})

	t.Run("small size", func(t *testing.T) {
		size := 10
		result := generateRandomElements(size)

		assert.Len(t, result, size, "Should  return slice of correct length")

		for _, val := range result {
			assert.GreaterOrEqual(t, val, 0, "Element should be >= 0")
			assert.Less(t, val, 1000, "Element should be < 1000")
		}
	})

	t.Run("determinism", func(t *testing.T) {
		size := 100
		result1 := generateRandomElements(size)
		result2 := generateRandomElements(size)

		assert.Equal(t, result1, result2, "Same seed should produce same results")
	})

	t.Run("different sizes", func(t *testing.T) {
		sizes := []int{1, 100, 1000}
		for _, size := range sizes {
			t.Run(string(rune(size)), func(t *testing.T) {
				res := generateRandomElements(size)
				assert.Len(t, res, size)
			})
		}
	})
}

func TestMaximum(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		res := maximum([]int{})
		assert.Equal(t, 0, res, "Should return 0 for empty slice")
	})

	t.Run("single element", func(t *testing.T) {
		res := maximum([]int{42})
		assert.Equal(t, 42, res)
	})

	t.Run("max at different positions", func(t *testing.T) {
		testCases := []struct {
			name     string
			input    []int
			expected int
		}{
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
	})

	t.Run("with negative numbers", func(t *testing.T) {
		res := maximum([]int{-5, -1, -10, 0, 5})
		assert.Equal(t, 5, res)
	})
}

func TestMaxChunks(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		res := maxChunks([]int{})
		assert.Equal(t, 0, res, "Should return 0 for empty slice")
	})

	t.Run("single element", func(t *testing.T) {
		res := maxChunks([]int{42})
		assert.Equal(t, 42, res)
	})

	t.Run("less elements than chunks", func(t *testing.T) {
		res := maxChunks([]int{1, 5, 3})
		assert.Equal(t, 5, res)
	})

	t.Run("exact number of chunks", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8}
		res := maxChunks(input)
		assert.Equal(t, 8, res)
	})

	t.Run("with ending", func(t *testing.T) {
		input := make([]int, 10)
		for i := range input {
			input[i] = i
		}
		input[9] = 100

		res := maxChunks(input)
		assert.Equal(t, 100, res)
	})
}
