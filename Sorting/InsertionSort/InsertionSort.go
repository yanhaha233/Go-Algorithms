package insertionsort

/*
	InsertionSort   插入排序
*/

func InsertionSort(arr []int) []int {
	var i, j int
	for i = 1; i < len(arr); i++ {
		for j = 0; j < i; j++ {
			if arr[j] > arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
