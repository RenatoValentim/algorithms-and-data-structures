package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollection(t *testing.T) {
	assert := assert.New(t)

	testCase := []struct {
		err      error
		expected *array
		name     string
		input    []int
		size     int
	}{
		{
			name:     "Should return an error for invalid array size",
			input:    nil,
			size:     0,
			expected: nil,
			err:      ErrIvalidArraySize,
		},
		{
			name:  "Should create a new array with inputted size",
			input: nil,
			size:  10,
			expected: &array{
				Items: nil,
				Size:  10,
			},
			err: nil,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			a, err := NewArray(tc.input, tc.size)
			assert.Equal(err, tc.err)
			assert.Equal(a, tc.expected)
		})
	}
}
