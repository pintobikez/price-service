package structures

import st "github.com/pintobikez/price-service/repository/structures"

type Product struct {
	ID     string           `json:"id"`
	Prices []*ProductPrices `json:"prices"`
}

type Channel struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type ProductPrices struct {
	Price        float64        `json:"price"`
	SpecialPrice st.NullFloat64 `json:"specialPrice"`
	SpecialFrom  st.NullTime    `json:"specialFrom"`
	SpecialTo    st.NullTime    `json:"specialTo"`
	Channel      string         `json:"channel"`
	UpdatedAt    st.NullTime    `json:"updatedAt"`
	ChannelID    int64          `json:"-"`
}

type HealthStatus struct {
	Pub  *HealthStatusDetail `json:"publisher"`
	Repo *HealthStatusDetail `json:"repository"`
}

type HealthStatusDetail struct {
	Status string `json:"status"`
	Detail string `json:"detail,omitempty"`
}
