package mergesort

import (
	"reflect"
	"testing"
)

func TestArr(t *testing.T) {
	got := MergeSort([]int{2, 3, 1, 5, 8, 7, 4, 6, 9})
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
