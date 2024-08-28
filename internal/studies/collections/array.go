package collections

type array struct {
	Items []int
	Size  int
}

func NewArray(items []int, size int) (*array, error) {
	if size <= 0 {
		return nil, ErrIvalidArraySize
	}
	return &array{
		Items: []int{},
		Size:  size,
	}, nil
}
