package collections

type Array struct {
	Items []int
	Size  int
}

func NewArray(items []int, size int) (*Array, error) {
	if size <= 0 {
		return nil, ErrIvalidArraySize
	}
	return &Array{
		Items: []int{},
		Size:  size,
	}, nil
}
