package searches

func BinarySearch(array []int, value int) int {
	if len(array) == 0 {
		return -1
	}
	minIndex := 0
	maxIndex := len(array) - 1
	for minIndex <= maxIndex {
		middleIndex := (minIndex + maxIndex) / 2
		guess := array[middleIndex]
		if guess == value {
			return middleIndex
		}
		if guess > value {
			maxIndex = middleIndex - 1
		} else {
			minIndex = middleIndex + 1
		}
	}
	return -1
}
