package cocktailsort

import (
	"reflect"
	"testing"
)

func TestArr(t *testing.T) {
	arr := []int{2, 9, 5, 6, 1, 3, 4, 8, 7, 10}
	got := CocktailSort(arr)
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
