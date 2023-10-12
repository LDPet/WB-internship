package natssteam

import (
	"github.com/nats-io/stan.go"
)

type Chanel struct {
	url       string
	clusterID string
	clientID  string

	sc stan.Conn
}

func NewNatsChanel(url string, clusterID string, clientID string) Chanel {
	return Chanel{
		url:       url,
		clusterID: clusterID,
		clientID:  clientID,
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

func (c *Chanel) Publish(subject string, data []byte) error {
	return c.sc.Publish(subject, data)
}

func (c *Chanel) Close() {
	c.sc.Close()
}
