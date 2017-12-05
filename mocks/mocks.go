package mocks

import (
	"fmt"
	gen "github.com/pintobikez/price-service/api/structures"
)

// MOCK STRUCTURES DEFINITION
type (
	RepositoryMock struct {
		Iserror bool
	}
	PublisherMock struct {
		Iserror bool
	}
)

// MOCK Repository - START
func (c *RepositoryMock) Connect() error {
	return nil
}
func (c *RepositoryMock) Disconnect() {
	return
}

func (c *RepositoryMock) FindChannel(channel string) (*gen.Channel, error) {
	return nil, nil
}
func (c *RepositoryMock) FindProduct(id string) (*gen.Product, error) {
	return nil, nil
}
func (c *RepositoryMock) GetChannels() (map[string]int64, error) {
	return make(map[string]int64), nil
}
func (c *RepositoryMock) PutProduct(s *gen.Product) (int64, error) {
	return 0, nil
}

func (c *RepositoryMock) Health() error {
	if c.Iserror {
		return fmt.Errorf("Erro Health")
	}
	return nil
}

// MOCK Repository - END

// MOCK Publisher - START
func (c *PublisherMock) Connect() error {
	return nil
}
func (c *PublisherMock) Close() {
	return
}
func (c *PublisherMock) Publish(s *gen.Product) error {
	if s.ID == "SCD" {
		return fmt.Errorf("Erro")
	}
	return nil
}
func (c *PublisherMock) Health() error {
	if c.Iserror {
		return fmt.Errorf("Erro Health")
	}
	return nil
}

// MOCK Publisher - END
