package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushArray(t *testing.T) {
	assert := assert.New(t)

	testCase := []struct {
		pushErr  error
		initial  *array
		expected *array
		name     string
		toPush   int
	}{
		{
			name:   "Should push an element into the array",
			toPush: 5,
			initial: &array{
				Items: nil,
				Size:  10,
			},
			expected: &array{
				Items: []int{5},
				Size:  10,
			},
			pushErr: nil,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.initial.Push(tc.toPush)
			assert.Equal(err, tc.pushErr)
			assert.Equal(tc.expected, tc.initial)
		})
	}
}
