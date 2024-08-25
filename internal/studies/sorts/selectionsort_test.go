package sorts_test

import (
	"testing"

	"github.com/RenatoValentim/algorithms-and-data-structures/internal/studies/sorts"
	"github.com/stretchr/testify/assert"
)

func TestSelectionSort(t *testing.T) {
	assert := assert.New(t)

	testCases := []struct {
		name        string
		input       []int
		targetValue int
		expected    int
	}{
		{
			name:        "Should return a negative index if the array is nil",
			input:       nil,
			targetValue: 3,
			expected:    -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := sorts.SelectionSort(tc.input, tc.targetValue)
			assert.Equal(tc.expected, output)
		})
	}
}
