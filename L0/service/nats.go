package main

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"os"
	"service/cache"
	"service/db"
	"service/model"
	"service/model/representation"
	"service/streaming"
	"service/streaming/natssteam"
	"sync"
)

func connectNats() (streaming.Chanel, error) {
	url := "nats://" + os.Getenv("NATS_HOST_IN") + ":" + os.Getenv("NATS_PORT")
	clusterID := os.Getenv("NATS_CLUSTER_ID")
	clientID := "service"
	durableName := os.Getenv("DURABLE_NAME")

	chanel := natssteam.NewNatsChanel(url, clusterID, clientID, durableName)
	err := chanel.Connect()
	if err != nil {
		return nil, fmt.Errorf("nats connecting error: %s", err)
	}

	return &chanel, nil
}

func subscribe(chanel streaming.Chanel, repository db.OrderRepository, cacher cache.Cacher[string, []byte]) {
	subj := os.Getenv("NATS_SUBJECT")

	handler := func(msg *stan.Msg) {
		log.Println("Message received")

		orderMsg := model.Order{}
		err := json.Unmarshal(msg.Data, &orderMsg)
		if err != nil {
			log.Printf("Error: %s\n", err)
		} else {
			log.Printf("Message with uid: %s parsed\n", orderMsg.OrderUid)
		}
		order := representation.Order{Uid: orderMsg.OrderUid, Body: msg.Data}

		wg := &sync.WaitGroup{}

		wg.Add(1)
		writeOrderInDb(wg, repository, &order)
		wg.Add(1)
		writeOrderInCache(wg, cacher, &order)

		wg.Wait()
	}

	err := chanel.Subscribe(subj, handler)
	if err != nil {
		log.Fatal("subscribe error: ", err)
	}
}
