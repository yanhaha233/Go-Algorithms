package bubblesort

import (
	"reflect"
	"testing"
)

func TestArr(t *testing.T) {
	arr := []int{1, 10, 2, 4, 5, 9, 7, 6, 8, 3}
	got := BubbleSort(arr)
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
