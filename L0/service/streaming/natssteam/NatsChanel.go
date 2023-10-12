package natssteam

import (
	"errors"
	"github.com/nats-io/stan.go"
)

type Chanel struct {
	url         string
	clusterID   string
	clientID    string
	durableName stan.SubscriptionOption

	sc  stan.Conn
	sub stan.Subscription
}

func NewNatsChanel(url string, clusterID string, clientID string, durableName string) Chanel {
	return Chanel{
		url:         url,
		clusterID:   clusterID,
		clientID:    clientID,
		durableName: stan.DurableName(durableName),
	}
}

func (c *Chanel) Connect() error {
	sc, err := stan.Connect(c.clusterID, c.clientID, stan.NatsURL(c.url))
	if err != nil {
		return err
	}

	c.sc = sc
	return nil
}

func (c *Chanel) Subscribe(subject string, handler any) error {
	hnd, ok := handler.(func(msg *stan.Msg))
	if !ok {
		return errors.New("expected MsgHandler type of handler")
	}
	sub, err := c.sc.Subscribe(subject, hnd, c.durableName)
	if err != nil {
		return err
	}
	c.sub = sub

	return nil
}

func (c *Chanel) Unsubscribe() {
	c.sub.Unsubscribe()
}

func (c *Chanel) Close() {
	c.sc.Close()
}
