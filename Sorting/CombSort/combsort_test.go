package combsort

import (
	"reflect"
	"testing"
)

func TestArr(t *testing.T) {
	arr := []int{2, 3, 5, 1, 4}
	got := CombSort(arr)
	want := []int{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
