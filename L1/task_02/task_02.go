package main

import "sync"

func main() {
	nums := []int{2, 3, 5, 7, 11, 13} // входные данные
	// для синхронизации "потомков", т.к. при завершении основной горутины (для main) завершаются и дочернии.
	// похоже на симафор
	wg := &sync.WaitGroup{}

	for _, num := range nums {
		wg.Add(1) // инкремент счётчика wg
		go printNum(wg, num)
	}

	wg.Wait() // ждём обнуления счётчика wg
}

func printNum(wg *sync.WaitGroup, num int) {
	println(num * num)
	wg.Done() // декремент счётчика wg (wg.Add(-1))
}
