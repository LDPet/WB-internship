package streaming

type Chanel interface {
	Connect() error
	Subscribe(subject string, handler any) error
	Unsubscribe()
	Close()
}
