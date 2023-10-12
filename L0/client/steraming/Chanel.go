package steraming

type Chanel interface {
	Connect() error
	Publish(subject string, data []byte) error
	Close()
}
