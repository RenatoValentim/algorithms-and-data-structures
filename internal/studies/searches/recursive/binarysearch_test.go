package recursive_test

import (
	"testing"

	"github.com/RenatoValentim/algorithms-and-data-structures/internal/studies/searches/recursive"
	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	assert := assert.New(t)

	testCases := []struct {
		name        string
		input       []int
		targetValue int
		expected    int
	}{
		{
			name:        "Should return a negative index if array is nil",
			input:       nil,
			targetValue: 3,
			expected:    -1,
		},
		{
			name:        "Should return a negative index if array is empty",
			input:       []int{},
			targetValue: 3,
			expected:    -1,
		},
		{
			name:        "Should return a negative index if value not in array",
			input:       []int{0, 2},
			targetValue: 3,
			expected:    -1,
		},
		{
			name:        "Should return the index for the searched value when the guess is equals for the value",
			input:       []int{0, 2, 3, 4, 5},
			targetValue: 3,
			expected:    2,
		},
		{
			name:        "Should return the index for the searched value when the value is in the array but smaller than the guess",
			input:       []int{0, 2, 3, 4, 5},
			targetValue: 2,
			expected:    1,
		},
		{
			name:        "Should return the index for the searched value when the value is in the array but greater than the guess",
			input:       []int{0, 2, 3, 4, 5},
			targetValue: 4,
			expected:    3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := recursive.BinarySearch(tc.input, tc.targetValue)
			assert.Equal(tc.expected, output)
		})
	}
}
