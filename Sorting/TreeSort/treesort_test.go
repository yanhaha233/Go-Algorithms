package TreeSort

import (
	"reflect"
	"testing"
)

func TestTreeSort(t *testing.T) {
	got := TreeSort([]int{2,3,1,5,4,7,6,9,8},&btree{nil})
	want := []int{1,2,3,4,5,6,7,8,9}

	if !reflect.DeepEqual(got,want){
		t.Errorf("got %v want %v",got,want)
	}
}
