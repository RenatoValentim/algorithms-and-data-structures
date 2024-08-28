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
		Items: items,
		Size:  size,
	}, nil
}

func (a *array) Push(i int) {
	a.Items = append(a.Items, i)
}
