package fn

// ComparableArray is an array of comparable and has methods that
// mimic a subset of the JavaScript array functions that are
// unique to comparables.
type ComparableArray[T comparable] AnyArray[T]

// Includes determines whether an array includes a certain value
// among its entries, returning true or false as appropriate.
func (a ComparableArray[T]) Includes(b T) bool {
	return Includes(a, b)
}

// IndexOf returns the first index at which a given element
// can be found in the array, or -1 if it is not present.
func (a ComparableArray[T]) IndexOf(b T) int {
	return IndexOf(a, b)
}

// LastIndexOf returns the last index at which a given element
// can be found in the array, or -1 if it is not present.
// The array is searched backwards, starting at fromIndex.
func (a ComparableArray[T]) LastIndexOf(b T, from int) int {
	return LastIndexOf(a, b, from)
}

// Unique creates a new array with only unique elements from the given array.
func (a ComparableArray[T]) Unique() ComparableArray[T] {
	return Unique(a)
}

// Filter creates a new array with all elements that pass
// the test implemented by the provided function.
func (a ComparableArray[T]) Filter(f func(T, int) bool) ComparableArray[T] {
	return ComparableArray[T](Filter(a, f))
}

// Reverse creates a new array by reversing the given array.
// The first array element becomes the last, and the last array element becomes the first.
func (a ComparableArray[T]) Reverse() ComparableArray[T] {
	return ComparableArray[T](Reverse(a))
}

// Slice returns a shallow copy of a portion
// of an array into a new array object selected
// from start to end (end not included) where
// start and end represent the index of items in that array.
// The original array will not be modified.
//
// If start<0, it is treated as distance from the end of the array.
// If end<=0, it is treated as distance from the end of the array.
func (a ComparableArray[T]) Slice(start, end int) ComparableArray[T] {
	return ComparableArray[T](Slice(a, start, end))
}
