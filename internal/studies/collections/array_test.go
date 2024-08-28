package collections_test

import (
	"testing"

	"github.com/RenatoValentim/algorithms-and-data-structures/internal/studies/collections"
	"github.com/stretchr/testify/assert"
)

func TestCollection(t *testing.T) {
	assert := assert.New(t)

	testCase := []struct {
		err      error
		expected *collections.Array
		name     string
		input    []int
		size     int
	}{
		{
			name:     "Should return an error for invalid array size ",
			input:    nil,
			size:     0,
			expected: nil,
			err:      collections.ErrIvalidArraySize,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			_, err := collections.NewArray(tc.input, tc.size)
			assert.Equal(err, tc.err)
		})
	}
}
