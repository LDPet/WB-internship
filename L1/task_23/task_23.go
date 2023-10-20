package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Without order")
	fmt.Println(arr1)
	arr1 = delWithoutOrder(arr1, 5)
	fmt.Println(arr1)

	fmt.Println("With order")
	fmt.Println(arr2)
	arr2 = delWithOrder(arr2, 5)
	fmt.Println(arr2)

}

func delWithoutOrder(arr []int, i int) []int {
	arr[i] = arr[len(arr)-1]
	return arr[:len(arr)-1]
}

func delWithOrder(arr []int, i int) []int {
	return append(arr[:i], arr[i+1:]...)
}
