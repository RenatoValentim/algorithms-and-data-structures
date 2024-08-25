package sorts

func SelectionSort(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	aux := 0
	for range arr {
		for currentIndex := range arr {
			if currentIndex < len(arr)-1 {
				nextIndex := currentIndex + 1
				if arr[currentIndex] > arr[nextIndex] {
					aux = arr[currentIndex]
					arr[currentIndex] = arr[nextIndex]
					arr[nextIndex] = aux
				}
			}
		}
	}
	return arr
}
