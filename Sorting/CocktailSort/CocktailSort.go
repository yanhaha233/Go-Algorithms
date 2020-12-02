package cocktailsort

/*
	CocktailSort   鸡尾酒排序->双向冒泡排序.
*/

func CocktailSort(arr []int) (arrs []int) {
	tmp := 0

	for i := 0; i < len(arr)/2; i++ {
		left := 0
		right := len(arr) - 1
		for left < right {
			if arr[left] > arr[left+1] {
				tmp = arr[left+1]
				arr[left+1] = arr[left]
				arr[left] = tmp
			}
			left++

			if arr[right-1] > arr[right] {
				tmp = arr[right]
				arr[right] = arr[right-1]
				arr[right-1] = tmp
			}
			right--
		}
	}
	return arr
}
