package main

import (
	"fmt"
	"log"
	"service/cache"
	"service/db"
	"service/model/representation"
	"sync"
)

func initCache(repository db.OrderRepository, cacher cache.Cacher[string, []byte]) error {
	orders, err := repository.GetOrders()
	if err != nil {
		return fmt.Errorf("GetOrders error: %s", err)
	}

	for _, order := range orders {
		cacher.Set(order.Uid, order.Body)
	}

	return nil
}

func writeOrderInCache(wg *sync.WaitGroup, cacher cache.Cacher[string, []byte], order *representation.Order) {
	go func() {
		cacher.Set(order.Uid, order.Body)

		log.Printf("Message with uid: %s recorded in cache\n", (*order).Uid)

		wg.Done()
	}()
}
