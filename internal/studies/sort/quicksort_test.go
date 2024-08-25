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
			name:     "Should return a negative index if the array is nil",
			input:    nil,
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := sort.Quicksort(tc.input)
			assert.Equal(tc.expected, output)
		})
	}
}
