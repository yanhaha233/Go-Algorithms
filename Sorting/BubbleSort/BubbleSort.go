package bubblesort

/*
	BubbleSort 冒泡排序
*/

func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		tmp := 0
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				tmp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
	return arr
}
