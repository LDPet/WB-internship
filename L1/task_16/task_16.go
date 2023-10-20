package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	l := 10000
	arr := make([]int, 0, l)
	for i := 0; i < l; i++ {
		arr = append(arr, rand.Intn(l*10))
	}

	//fmt.Println(arr)
	start := time.Now()
	qsort(arr, func(i int, i2 int) bool {
		return i < i2
	})
	duration := time.Since(start)
	//fmt.Println(arr)
	fmt.Println(sort.SliceIsSorted(arr, func(i int, i2 int) bool {
		return i < i2
	}))
	fmt.Println(duration)
}

func qsort[T any](arr []T, less func(T, T) bool) {
	_qsort(arr, 0, len(arr), less)
}

// быстрая сортировка работает по парадигме "разделяй и властвуй" (раздел, властвование, объединение)
// худший случай O(n^2)
// средний случай O(n*log(n))
func _qsort[T any](arr []T, l int, r int, less func(T, T) bool) {
	// условие выхода из рекурсии
	// конда на входе массив из 1 элемента (по умолчанию отсортирован)
	if l < r {
		// разделяем
		q := partition(arr, l, r, less)
		// властвуем
		_qsort(arr, l, q, less)
		_qsort(arr, q+1, r, less)
		// явного объединение нет, т.к. и разделеения по факту не было
		// не делили, а ограничили области индексами
	}
}

func partition[T any](arr []T, l int, r int, less func(T, T) bool) int {
	// рандомизация алгоритма
	// уменьшает вероятности худшего случаю
	// при случайном распределении элементов по массиву немного замедляет
	// если приходят данные, которые ближе к упорядоченным или уже упрядочены, значительно ускоряет работу
	ri := rand.Intn(r-l) + l
	arr[r-1], arr[ri] = arr[ri], arr[r-1]

	//устанавливаем значение барьера
	x := arr[r-1]
	// индекс вставки следующего элемента меньшего барьера
	i := l

	for j := l; j < r; j++ {
		// меньше элементы групппируем в начале
		if less(arr[j], x) {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	// ставим барьер на место
	arr[i], arr[r-1] = arr[r-1], arr[i]

	// возврат индекса барьера
	return i
}
