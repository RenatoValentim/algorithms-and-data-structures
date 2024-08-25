package sorts_test

import (
	"testing"

	"github.com/RenatoValentim/algorithms-and-data-structures/internal/studies/sorts"
	"github.com/stretchr/testify/assert"
)

func TestSelectionSort(t *testing.T) {
	assert := assert.New(t)

	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Should return a negative index if the array is nil",
			input:    nil,
			expected: []int{},
		},
		{
			name:     "Should return a negative index if array is empty",
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
			output := sorts.SelectionSort(tc.input)
			assert.Equal(tc.expected, output)
		})
	}
}
