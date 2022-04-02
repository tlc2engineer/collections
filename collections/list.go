package collections

type List[T any] interface {
	Get(index int) (error, T)
	Add(values ...T)
	Contains(values ...T) bool
	AddList(list List[T])
	Insert(index int, values ...T) error
	Remove(index int) error
	SubList(start, end int) (error, List[T])
	Set(index int, value T) error
	//Container[T]
}
