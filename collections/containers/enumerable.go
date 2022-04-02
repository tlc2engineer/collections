package collections

type EnumerableWithIndex[T any] interface {
	// Each calls the given function once for each element, passing that element's index and value.
	Each(func(index int, value T))

	// Map invokes the given function once for each element and returns a
	// container containing the values returned by the given function.
	// TODO would appreciate help on how to enforce this in containers (don't want to type assert when chaining)
	// Map(func(index int, value interface{}) interface{}) Container

	// Select returns a new container containing all elements for which the given function returns a true value.
	// TODO need help on how to enforce this in containers (don't want to type assert when chaining)
	// Select(func(index int, value interface{}) bool) Container

	// Any passes each element of the container to the given function and
	// returns true if the function ever returns true for any element.
	Any(func(index int, value T) bool) bool

	// All passes each element of the container to the given function and
	// returns true if the function returns true for all elements.
	All(func(index int, value T) bool) bool

	// Find passes each element of the container to the given function and returns
	// the first (index,value) for which the function is true or -1,nil otherwise
	// if no element matches the criteria.
	Find(func(index int, value T) bool) (int, T)
}
