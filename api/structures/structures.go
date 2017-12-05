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
	SpecialPrice st.NullFloat64 `json:"specialPrice, omitempty"`
	SpecialFrom  st.NullTime    `json:"specialFrom, omitempty"`
	SpecialTo    st.NullTime    `json:"specialTo, omitempty"`
	Channel      string         `json:"channel"`
	UpdatedAt    st.NullTime    `json:"updatedAt, omitempty"`
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
