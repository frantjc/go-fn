package main

import (
	"fmt"

	"github.com/frantjc/go-fn"
)

func main() {
	var (
		_ fn.AnyArray[any]
	)

	fn.Map([]int{1, 2, 3}, func(e, i int) int {
		return e + 1
	}).Filter(func(e, i int) bool {
		return e > 2
	}).ForEach(func(e, i int) {
		fmt.Printf("[%d] = %d\n", i, e)
	})
	// [0] = 3
	// [1] = 4
}
