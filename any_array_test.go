package fn_test

import (
	"testing"

	"github.com/frantjc/go-fn"
)

func TestAnyEveryTrue(t *testing.T) {
	var (
		a = fn.AnyArray[int]([]int{1, 2, 3, 4})
		f = func(a, _ int) bool {
			return a > 0
		}
		expected = true
		actual   = a.Every(f)
	)
	if expected != actual {
		t.Error("actual", actual, "does not equal expected", expected)
		t.FailNow()
	}
}

func TestAnyEveryFalse(t *testing.T) {
	var (
		a = fn.AnyArray[int]([]int{0, 1, 2, 3})
		f = func(a, _ int) bool {
			return a > 0
		}
		expected = false
		actual   = a.Every(f)
	)
	if expected != actual {
		t.Error("actual", actual, "does not equal expected", expected)
		t.FailNow()
	}
}

func TestAnyFilter(t *testing.T) {
	var (
		a = fn.AnyArray[int]([]int{0, 1, 2, 3})
		f = func(a, _ int) bool {
			return a > 0
		}
		expected = fn.AnyArray[int]([]int{1, 2, 3})
		actual   = a.Filter(f)
	)
	for i := range expected {
		if expected[i] != actual[i] {
			t.Error("actual", actual, "does not equal expected", expected)
			t.FailNow()
		}
	}
}

func TestAnyFindIndex(t *testing.T) {
	var (
		a = fn.AnyArray[int]([]int{1, 2, 3, 4})
		f = func(a, _ int) bool {
			return a == 2
		}
		expected = 1
		actual   = a.FindIndex(f)
	)
	if expected != actual {
		t.Error("actual", actual, "does not equal expected", expected)
		t.FailNow()
	}
}

func TestAnyFind(t *testing.T) {
	var (
		a = fn.AnyArray[int]([]int{1, 2, 3, 4})
		f = func(a, _ int) bool {
			return a > 0
		}
		expected = 1
		actual   = a.Find(f)
	)
	if expected != actual {
		t.Error("actual", actual, "does not equal expected", expected)
		t.FailNow()
	}
}

func TestAnyReverse(t *testing.T) {
	var (
		a        = fn.AnyArray[int]([]int{1, 2, 3, 4, 3, 2})
		expected = fn.AnyArray[int]([]int{2, 3, 4, 3, 2, 1})
		actual   = a.Reverse()
	)
	for i := range expected {
		if expected[i] != actual[i] {
			t.Error("actual", actual, "does not equal expected", expected)
			t.FailNow()
		}
	}
}

func TestAnySliceFromStart(t *testing.T) {
	var (
		a        = fn.AnyArray[int]([]int{0, 1, 2, 3})
		expected = fn.AnyArray[int]([]int{0, 1})
		actual   = a.Slice(0, 2)
	)
	for i := range expected {
		if expected[i] != actual[i] {
			t.Error("actual", actual, "does not equal expected", expected)
			t.FailNow()
		}
	}
}

func TestAnySliceFromEnd(t *testing.T) {
	var (
		a        = fn.AnyArray[int]([]int{0, 1, 2, 3})
		expected = fn.AnyArray[int]([]int{2})
		actual   = a.Slice(-2, -1)
	)
	for i := range expected {
		if expected[i] != actual[i] {
			t.Error("actual", actual, "does not equal expected", expected)
			t.FailNow()
		}
	}
}

func TestAnySomeTrue(t *testing.T) {
	var (
		a = fn.AnyArray[int]([]int{1, 2, 3, 4})
		f = func(a, _ int) bool {
			return a == 4
		}
		expected = true
		actual   = a.Some(f)
	)
	if expected != actual {
		t.Error("actual", actual, "does not equal expected", expected)
		t.FailNow()
	}
}

func TestAnySomeFalse(t *testing.T) {
	var (
		a = fn.AnyArray[int]([]int{1, 2, 3, 4})
		f = func(a, _ int) bool {
			return a == 5
		}
		expected = false
		actual   = a.Some(f)
	)
	if expected != actual {
		t.Error("actual", actual, "does not equal expected", expected)
		t.FailNow()
	}
}
