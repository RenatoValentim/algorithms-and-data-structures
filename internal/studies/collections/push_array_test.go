package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushArray(t *testing.T) {
	assert := assert.New(t)

	testCase := []struct {
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
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			tc.initial.Push(tc.toPush)
			assert.Equal(tc.expected, tc.initial)
		})
	}
}
