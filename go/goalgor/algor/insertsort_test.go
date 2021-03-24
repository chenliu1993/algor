package algor

import (
	"testing"
)

func TestInsertSort(t *testing.T) {
	src := []int{1, 4, 3, 5, 6, 2, 7, 8}
	InsertSort(src)
	if isSameSlice(src, src) {
		t.Error(("Wrong!!"))
	}
}

func BenchmarkInsertSort(b *testing.B) {
	src := []int{1, 4, 3, 5, 6, 2, 7, 8}
	InsertSort((src))
}

func ExampleInsertTest() {
	src := []int{1, 4, 3, 5, 6, 2, 7, 8}
	InsertSort(src)
}
