package main

func simple(nums []int) int {
	const wCount = 8
	// создаём каныла для общения с работягами.
	//каналы буфферизированные, чтоб было меньше блокировок при записи и чтении
	jobs := make(chan int, wCount)
	results := make(chan int, wCount)
	res := 0

	//запуск работяг
	for i := 0; i < wCount; i++ {
		go worker(jobs, results)
	}

	//выдача заданий
	for _, num := range nums {
		jobs <- num
	}
	defer close(jobs)

	//получаем результат работы
	for i := 0; i < len(nums); i++ {
		res += <-results
	}
	defer close(results)

	return res
}

func worker(jobs <-chan int, results chan<- int) {
	//сам работяга ожидает и обрабатывает задания до закрытия канала
	for num := range jobs {
		results <- num * num
	}
}
