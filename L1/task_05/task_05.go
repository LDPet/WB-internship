package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	// для страта необходимо 2 аргумента комндной строки
	// 1 - путь к проге
	// 2 - кол-во воркеров
	// 2 - время работы
	if len(os.Args) < 3 {
		println("Укажите количество воркеров и время работы")
		os.Exit(1)
	}

	// полчение int из string
	wCount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		println(err)
		os.Exit(1)
	}
	ttl, err := strconv.Atoi(os.Args[2])
	if err != nil {
		println(err)
		os.Exit(1)
	}
	if ttl < 0 {
		println("Время работы должно быть больше 0")
		os.Exit(1)
	}

	jobs := make(chan int, wCount)
	wg := sync.WaitGroup{}

	//запуск работяг
	for i := 0; i < wCount; i++ {
		wg.Add(1)
		go worker(i, &wg, jobs)
	}

	// создаём канал, в который через через ttl секунд придёт сигнвл
	quit := time.After(time.Duration(ttl) * time.Second)

	//выдача заданий
	for {
		select {
		// по приходу сигнала закрываем канал и завершаем программу
		case <-quit:
			// канал закрывается для оставновки работяг
			close(jobs)
			// дожидаемся работяг
			wg.Wait()
			return
		// пока не пришел сигнал создаём задания
		default:
			jobs <- rand.Int()
		}
	}
}

func worker(id int, wg *sync.WaitGroup, jobs <-chan int) {
	//сам работяга ожидает и обрабатывает задания до закрытия канала
	for num := range jobs {
		fmt.Printf("%d: %d \n", id, num)
	}
	wg.Done()
}
