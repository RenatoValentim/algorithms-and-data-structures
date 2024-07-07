package searches_test

import (
	"testing"

	"github.com/RenatoValentim/algorithms-and-data-structures/internal/studies/searches"
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
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			bs := searches.NewBinarySearch(tc.input)
			output := bs.Search(tc.targetValue)
			assert.Equal(tc.expected, output)
		})
	}
}
