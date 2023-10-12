package main

import (
	"fmt"
	"log"
	"os"
	"service/db"
	"service/db/PgRepository"
	"service/model/representation"
	"sync"
)

func openDB() (db.OrderRepository, error) {
	dataSourceName := "host=" + os.Getenv("PG_HOST") + " port=" + os.Getenv("PG_PORT") +
		" user=" + os.Getenv("PG_USER") + " password=" + os.Getenv("PG_PASSWORD") +
		" dbname=" + os.Getenv("PG_DB") + " sslmode=disable"

	orderRepository := PgRepository.NewPgOrderRepository()
	err := orderRepository.Open(dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("nats connecting error: %s", err)
	}

	return &orderRepository, nil
}

func writeOrderInDb(wg *sync.WaitGroup, repository db.OrderRepository, order *representation.Order) {
	go func() {
		err := repository.CreateOrder(*order)
		if err != nil {
			log.Println("Error: %s", err)
		} else {
			log.Printf("Message with uid: %s recorded in bd\n", (*order).Uid)
		}

		wg.Done()
	}()
}
