package searches

func BinarySearch(items []int, item int) int {
	if items == nil || len(items) == 0 {
		return -1
	}
	minIndex := 0
	maxIndex := len(items) - 1
	middleIndex := (minIndex + maxIndex) / 2
	if items[middleIndex] == item {
		return middleIndex
	}
	return -1
}
