package main

func main() {
	arr := []int{-3, -1, 0, 0, 1, 3, 5, 8, 10}

	i, ok := bSearch(arr, func(i int) int {
		return i - 6
	})
	println(i, ok)
	i, ok = bSearch(arr, func(i int) int {
		return i
	})
	println(i, ok)
}

// только отсортированный массив
func bSearch[T any](arr []T, cmp func(T) int) (int, bool) {
	// границы
	l, r := 0, len(arr)-1
	for l <= r {
		// чуть усложнено для избегаания переполнения
		mid := l + (r-l)/2
		//результат сравнения с целью
		val := cmp(arr[mid])
		// нашли
		if val == 0 {
			return mid, true
		}
		// если < 0 то arr[mid] < target => надо искать в правой половине => перемщаем l за середину
		// иначе аналогично
		if cmp(arr[mid]) < 0 {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	// не нашли
	return 0, false
}
