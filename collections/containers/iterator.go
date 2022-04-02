package collections

// IteratorWithIndex is stateful iterator for ordered containers whose values can be fetched by an index.
type IteratorWithIndex[T any] interface {
	// Next moves the iterator to the next element and returns true if there was a next element in the container.
	// If Next() returns true, then next element's index and value can be retrieved by Index() and Value().
	// If Next() was called for the first time, then it will point the iterator to the first element if it exists.
	// Modifies the state of the iterator.
	Next() bool

	// Value returns the current element's value.
	// Does not modify the state of the iterator.
	Value() T

	// Index returns the current element's index.
	// Does not modify the state of the iterator.
	Index() int

	// Begin resets the iterator to its initial state (one-before-first)
	// Call Next() to fetch the first element if any.
	Begin()

	// First moves the iterator to the first element and returns true if there was a first element in the container.
	// If First() returns true, then first element's index and value can be retrieved by Index() and Value().
	// Modifies the state of the iterator.
	First() bool
}

// IteratorWithKey is a stateful iterator for ordered containers whose elements are key value pairs.
type IteratorWithKey[K, V any] interface {
	// Next moves the iterator to the next element and returns true if there was a next element in the container.
	// If Next() returns true, then next element's key and value can be retrieved by Key() and Value().
	// If Next() was called for the first time, then it will point the iterator to the first element if it exists.
	// Modifies the state of the iterator.
	Next() bool

	// Value returns the current element's value.
	// Does not modify the state of the iterator.
	Value() V

	// Key returns the current element's key.
	// Does not modify the state of the iterator.
	Key() K

	// Begin resets the iterator to its initial state (one-before-first)
	// Call Next() to fetch the first element if any.
	Begin()

	// First moves the iterator to the first element and returns true if there was a first element in the container.
	// If First() returns true, then first element's key and value can be retrieved by Key() and Value().
	// Modifies the state of the iterator.
	First() bool
}

type ReverseIteratorWithIndex[T any] interface {
	// Prev moves the iterator to the previous element and returns true if there was a previous element in the container.
	// If Prev() returns true, then previous element's index and value can be retrieved by Index() and Value().
	// Modifies the state of the iterator.
	Prev() bool

	// End moves the iterator past the last element (one-past-the-end).
	// Call Prev() to fetch the last element if any.
	End()

	// Last moves the iterator to the last element and returns true if there was a last element in the container.
	// If Last() returns true, then last element's index and value can be retrieved by Index() and Value().
	// Modifies the state of the iterator.
	Last() bool

	IteratorWithIndex[T]
}

type ReverseIteratorWithKey[K, V any] interface {
	// Prev moves the iterator to the previous element and returns true if there was a previous element in the container.
	// If Prev() returns true, then previous element's key and value can be retrieved by Key() and Value().
	// Modifies the state of the iterator.
	Prev() bool

	// End moves the iterator past the last element (one-past-the-end).
	// Call Prev() to fetch the last element if any.
	End()

	// Last moves the iterator to the last element and returns true if there was a last element in the container.
	// If Last() returns true, then last element's key and value can be retrieved by Key() and Value().
	// Modifies the state of the iterator.
	Last() bool

	IteratorWithKey[K, V]
}
