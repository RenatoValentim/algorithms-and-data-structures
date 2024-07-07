package searches

type binarySearch struct {
	items []int
}

func NewBinarySearch(items []int) *binarySearch {
	return &binarySearch{
		items: items,
	}
}

func (bs *binarySearch) Search(item int) int {
	return -1
}
