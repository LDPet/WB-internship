package main

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"service/cache/amcacher"
	"sync"
	"syscall"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	repository, err := openDB()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB opened")
	defer repository.Close()

	chanel, err := connectNats()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Nats connected")
	defer chanel.Close()
	//defer chanel.Unsubscribe()

	cacher := amcacher.NewAsyncMapCacher[string, []byte]()
	log.Println("Cache created")

	err = initCache(repository, cacher)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	wg := sync.WaitGroup{}

	ctx = context.WithValue(ctx, "chan", quit)
	ctx = context.WithValue(ctx, "wg", &wg)

	go subscribe(chanel, repository, cacher)
	wg.Add(1)
	go runServer(ctx, cacher)

	sig := <-quit
	quit <- sig
	wg.Wait()
	log.Println("Service stop")
}
