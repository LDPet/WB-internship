package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

func main() {
	// для страта необходимо 2 аргумента комндной строки
	// 1 - путь к проге
	// 2 - кол-во воркеров
	if len(os.Args) < 2 {
		println("Укажите количество воркеров")
		os.Exit(1)
	}

	// полчение int из string
	wCount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		println(err)
		os.Exit(1)
	}

	jobs := make(chan int, wCount)
	wg := sync.WaitGroup{}

	//запуск работяг
	for i := 0; i < wCount; i++ {
		wg.Add(1)
		go worker(i, &wg, jobs)
	}

	// создаём канал и присоединяем его к сигналам
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

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
