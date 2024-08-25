package recursive

func BinarySearch(array []int, value int) int {
	if len(array) == 0 {
		return -1
	}
	minIndex := 0
	maxIndex := len(array) - 1
	middleIndex := (minIndex + maxIndex) / 2
	guess := array[middleIndex]
	if guess == value {
		return middleIndex
	}
	if guess > value {
		return BinarySearch(array[:middleIndex], value)
	} else {
		index := BinarySearch(array[middleIndex+1:], value)
		if index == -1 {
			return index
		}
		return middleIndex + 1 + index
	}
}
