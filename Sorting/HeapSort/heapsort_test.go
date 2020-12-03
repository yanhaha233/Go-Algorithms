package heapsort

import (
	"reflect"
	"testing"
)

func TestArr(t *testing.T) {
	got := HeapSort([]int{2, 3, 5, 6, 1, 4, 8, 7})
	want := []int{1, 2, 3, 4, 5, 6, 7, 8}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
