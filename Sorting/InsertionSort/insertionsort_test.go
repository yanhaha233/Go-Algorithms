package insertionsort

import (
	"reflect"
	"testing"
)

func TestArr(t *testing.T) {
	got := InsertionSort([]int{2, 5, 3, 1, 6, 4})
	want := []int{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
