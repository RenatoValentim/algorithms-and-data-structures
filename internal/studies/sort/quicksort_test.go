package sort_test

import (
	"testing"

	"github.com/RenatoValentim/algorithms-and-data-structures/internal/studies/sort"
	"github.com/stretchr/testify/assert"
)

func TestQuicksort(t *testing.T) {
	assert := assert.New(t)

	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Should return an empty array is nil",
			input:    nil,
			expected: []int{},
		},
		{
			name:     "Should return an empty array is empty",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Should sort the array values",
			input:    []int{0, 4, 3, 7, 9, 4},
			expected: []int{0, 3, 4, 4, 7, 9},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := sort.QuickSort(tc.input)
			assert.Equal(tc.expected, output)
		})
	}
}
