package containers

/*
Container - Контейнер чего-либо.
*/
type Container[T any] interface {
	Empty() bool
	Size() int
	Clear()
	Values() []T
}
