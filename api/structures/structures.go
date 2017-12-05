package structures

import "time"

type Product struct {
	ID     string           `json:"id"`
	Prices []*ProductPrices `json:"prices"`
}

type Channel struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type ProductPrices struct {
	Price        float64   `json:"price"`
	SpecialPrice float64   `json:"specialPrice"`
	SpecialFrom  time.Time `json:"specialFrom"`
	SpecialTo    time.Time `json:"specialTo"`
	Channel      string    `json:"channel"`
	UpdatedAt    time.Time `json:"updatedAt"`
	ChannelID    int64
}

type HealthStatus struct {
	Pub  *HealthStatusDetail `json:"publisher"`
	Repo *HealthStatusDetail `json:"repository"`
}

type HealthStatusDetail struct {
	Status string `json:"status"`
	Detail string `json:"detail,omitempty"`
}
