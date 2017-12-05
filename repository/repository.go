package repository

import gen "github.com/pintobikez/price-service/api/structures"

type Repository interface {
	Connect() error
	Disconnect()
	FindProduct(id string) (*gen.Product, error)
	GetChannels() (map[string]int64, error)
	PutProduct(s *gen.Product) (int64, error)
	Health() error
}
