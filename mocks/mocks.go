package mocks

import (
	"fmt"
	gen "github.com/pintobikez/price-service/api/structures"
)

// MOCK STRUCTURES DEFINITION
type (
	RepositoryMock struct {
		IsError   bool
		IsErrorCh bool
	}
	PublisherMock struct {
		IsError bool
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
	if id == "SC" {
		return new(gen.Product), nil
	}
	if id == "SCD" {
		a := new(gen.Product)
		a.ID = id
		return a, nil
	}
	if id == "SCA" {
		return nil, fmt.Errorf("Error")
	}
	return nil, nil
}
func (c *RepositoryMock) GetChannels() (map[string]int64, error) {
	ret := make(map[string]int64)

	if c.IsErrorCh {
		return ret, fmt.Errorf("Error")
	}

	ret["loja1"] = 1

	return ret, nil
}
func (c *RepositoryMock) PutProduct(s *gen.Product) (int64, error) {
	if c.IsError {
		return 0, fmt.Errorf("Error")
	}
	if s.ID == "SCA" || s.ID == "SCD" {
		return 1, nil
	}
	return 0, nil
}

func (c *RepositoryMock) Health() error {
	if c.IsError {
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
	if c.IsError {
		return fmt.Errorf("Erro Health")
	}
	return nil
}

// MOCK Publisher - END
