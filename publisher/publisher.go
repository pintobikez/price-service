package publisher

import gen "github.com/pintobikez/price-service/api/structures"

type PubSub interface {
	Connect() error
	Close()
	Publish(s *gen.Product) error
	Health() error
}
