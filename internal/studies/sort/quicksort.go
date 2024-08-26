package sort

import (
	"math/rand"
	"time"
)

func QuickSort(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	if len(arr) == 1 {
		return arr
	}

	rand.NewSource(time.Now().UnixNano())
	randomIndex := rand.Intn(len(arr))

	pivot := arr[randomIndex]
	left := []int{}
	right := []int{}
	center := []int{}

	for _, value := range arr {
		if value < pivot {
			left = append(left, value)
		}

		if value == pivot {
			center = append(center, value)
		}

		if value > pivot {
			right = append(right, value)
		}
	}

	return append(append(QuickSort(left), center...), QuickSort(right)...)
}
