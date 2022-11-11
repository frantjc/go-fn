package fn

// AnyArray is an array of any and has methods that
// mimic a subset of the JavaScript array functions.
type AnyArray[T any] []T

// Every tests whether all elements in the array pass the
// test implemented by the provided function. It returns a Boolean value.
func (a AnyArray[T]) Every(f func(T, int) bool) bool {
	return Every(a, f)
}

// Filter creates a new array with all elements that pass
// the test implemented by the provided function.
func (a AnyArray[T]) Filter(f func(T, int) bool) AnyArray[T] {
	return Filter(a, f)
}

// FindIndex returns the index of the first element in the array
// that satisfies the provided testing function. Otherwise, it returns -1,
// indicating that no element passed the test
func (a AnyArray[T]) FindIndex(f func(T, int) bool) int {
	return FindIndex(a, f)
}

// Find returns the first element in the provided array that satisfies
// the provided testing function. If no values satisfy the testing function,
// the zero value of T is returned.
func (a AnyArray[T]) Find(f func(T, int) bool) T {
	return Find(a, f)
}

// ForEach executes a provided function once
// for each array element
func (a AnyArray[T]) ForEach(f func(T, int)) {
	ForEach(a, f)
}

// Reverse creates a new array by reversing the given array.
// The first array element becomes the last, and the last array element becomes the first.
func (a AnyArray[T]) Reverse() AnyArray[T] {
	return Reverse(a)
}

// Slice returns a shallow copy of a portion
// of an array into a new array object selected
// from start to end (end not included) where
// start and end represent the index of items in that array.
// The original array will not be modified.
//
// If start<0, it is treated as distance from the end of the array.
// If end<=0, it is treated as distance from the end of the array.
func (a AnyArray[T]) Slice(start, end int) AnyArray[T] {
	return Slice(a, start, end)
}

// Some tests whether at least one element in the array passes
// the test implemented by the provided function.
// It returns true if, in the array, it finds an element
// for which the provided function returns true; otherwise
// it returns false. It doesn't modify the array.
func (a AnyArray[T]) Some(f func(T, int) bool) bool {
	return Some(a, f)
}
