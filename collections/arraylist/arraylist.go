package arraylist

type List[T any] struct {
	elements []T
	size     int
}

const (
	growthFactor = float32(2.0)  // growth by 100%
	shrinkFactor = float32(0.25) // shrink when size is 25% of capacity (0 means never shrink)
)

// New instantiates a new list and adds the passed values, if any, to the list
func New[T any](values ...T) *List[T] {
	list := &List[T]{}
	// if len(values) > 0 {
	// 	list.Add(values...)
	// }

	return list
}
