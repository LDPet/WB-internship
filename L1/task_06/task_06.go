package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	st := make(chan struct{})
	go chanelStop(st)
	time.Sleep(3 * time.Second)
	// отправляем сигнал как знак завершения
	st <- struct{}{}
	time.Sleep(time.Second)

	go chanelClose(st)
	time.Sleep(3 * time.Second)
	// закрываем канал как знак завершения
	close(st)
	time.Sleep(time.Second)

	ctx, cancel := context.WithCancel(context.Background())
	go ctxDone(ctx)
	time.Sleep(3 * time.Second)
	// создаёет сигнал в Done канале
	cancel()
	time.Sleep(time.Second)

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	go ctxTimeout(ctx)
	time.Sleep(3 * time.Second)
	// создаёет сигнал в Done канале через время (обёртка для WithDeadline)
	time.Sleep(time.Second)

	fmt.Println("main done")
}

func chanelStop(st chan struct{}) {
	for {
		select {
		case <-st:
			fmt.Println("chanelStop received signal")
			return
		default:
			fmt.Println("chanelStop work")
			time.Sleep(time.Second)
		}
	}
}

func chanelClose(st chan struct{}) {
	for {
		select {
		case <-st:
			fmt.Println("chanelStop received signal")
			return
		default:
			fmt.Println("chanelStop work")
			time.Sleep(time.Second)
		}
	}

	//for {
	//	_, ok := <-st
	//	if ok {
	//		fmt.Println("chanelClose work")
	//		time.Sleep(time.Second)
	//	} else {
	//		break
	//	}
	//}
	//fmt.Println("chanelClose chanel closed")
}

func ctxDone(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("ctxDone received signal")
			return
		default:
			fmt.Println("ctxDone work")
			time.Sleep(time.Second)
		}
	}
}

func ctxTimeout(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("ctxTimeout received signal")
			return
		default:
			fmt.Println("ctxTimeout work")
			time.Sleep(time.Second)
		}
	}
}
