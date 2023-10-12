package main

import (
	"client/steraming/natssteam"
	"encoding/json"
	"github.com/go-faker/faker/v4"
	"github.com/joho/godotenv"
	"log"
	"math/rand"
	"os"
	"reflect"
	"time"
)

type FakeOrder struct {
	OrderUid    string `json:"order_uid" faker:"uuid_digit"`
	TrackNumber string `json:"track_number" faker:"len=14"`
	Entry       string `json:"entry" faker:"oneof: WBIL"`
	Delivery    struct {
		Name    string `json:"name" faker:"name"`
		Phone   string `json:"phone" faker:"e_164_phone_number"`
		Zip     string `json:"zip" faker:"uuid_digit"`
		City    string `json:"city" faker:"len=10"`
		Address string `json:"address" faker:"len=10"`
		Region  string `json:"region" faker:"len=10"`
		Email   string `json:"email" faker:"email"`
	} `json:"delivery"`
	Payment struct {
		Transaction  string `json:"transaction" faker:"-"`
		RequestId    string `json:"request_id" faker:"len=10"`
		Currency     string `json:"currency" faker:"currency"`
		Provider     string `json:"provider" faker:"oneof: wbpay, sbp, cash, card"`
		Amount       int    `json:"amount" faker:"boundary_start=0, boundary_end=100"`
		PaymentDt    int    `json:"payment_dt"`
		Bank         string `json:"bank" faker:"len=10"`
		DeliveryCost int    `json:"delivery_cost" faker:"boundary_start=0, boundary_end=10000"`
		GoodsTotal   int    `json:"goods_total" faker:"boundary_start=0, boundary_end=1000"`
		CustomFee    int    `json:"custom_fee" faker:"boundary_start=0, boundary_end=100"`
	} `json:"payment"`
	Items []struct {
		ChrtId      int    `json:"chrt_id" faker:"boundary_start=0, boundary_end=10000"`
		TrackNumber string `json:"track_number" faker:"-"`
		Price       int    `json:"price" faker:"boundary_start=0, boundary_end=10000"`
		Rid         string `json:"rid" faker:"uuid_hyphenated"`
		Name        string `json:"name" faker:"len=10"`
		Sale        int    `json:"sale" faker:"boundary_start=0, boundary_end=100"`
		Size        string `json:"size" faker:"oneof: small, medium, large"`
		TotalPrice  int    `json:"total_price" faker:"boundary_start=0, boundary_end=10000"`
		NmId        int    `json:"nm_id" faker:"boundary_start=0, boundary_end=10000"`
		Brand       string `json:"brand" faker:"len=10"`
		Status      int    `json:"status" faker:"boundary_start=0, boundary_end=1000"`
	} `json:"items"`
	Locale            string `json:"locale" faker:"oneof: ru, en, us, uk, ch"`
	InternalSignature string `json:"internal_signature" faker:"password"`
	CustomerId        string `json:"customer_id" faker:"uuid_digit"`
	DeliveryService   string `json:"delivery_service" faker:"oneof: meest, cells, courier"`
	Shardkey          string `json:"shardkey" faker:"uuid_digit"`
	SmId              int    `json:"sm_id" faker:"boundary_start=5, boundary_end=10000"`
	DateCreated       string `json:"date_created"  faker:"timestampMy"`
	OofShard          string `json:"oof_shard" faker:"len=10"`
}

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	url := "nats://" + os.Getenv("NATS_HOST_IN") + ":" + os.Getenv("NATS_PORT")
	clusterID := os.Getenv("NATS_CLUSTER_ID")
	clientID := "client"
	subj := os.Getenv("NATS_SUBJECT")

	chanel := natssteam.NewNatsChanel(url, clusterID, clientID)
	err = chanel.Connect()
	if err != nil {
		log.Fatal("Error while connecting ", err)
	}
	defer chanel.Close()

	err = faker.AddProvider("timestampMy", func(v reflect.Value) (interface{}, error) {
		return time.Now().Format(time.RFC3339), nil
	})
	if err != nil {
		log.Fatal("add faker provider error ", err)
	}

	num := rand.Intn(10) + 1
	for i := 0; i < num; i++ {
		order, err := genOrder()
		if err != nil {
			log.Fatal("genOrder error ", err)
		}

		data, err := json.Marshal(*order)
		if err != nil {
			log.Fatal("Marshal error ", err)
		}
		err = chanel.Publish(subj, data)
		if err != nil {
			log.Fatal("Publish error", err)
		}
	}
}

func genOrder() (*FakeOrder, error) {
	res := FakeOrder{}
	err := faker.FakeData(&res)
	if err != nil {
		return nil, err
	}

	res.Payment.Transaction = res.OrderUid
	for i := range res.Items {
		res.Items[i].TrackNumber = res.TrackNumber
	}

	return &res, nil
}
