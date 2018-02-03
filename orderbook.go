package hitbtc

type OneOrder struct {
	Price float64 `json:"price,string"`
	Size  float64 `json:"size,string"`
}

type OrderBook struct {
	Ask []OneOrder `json:"ask"`
	Bid []OneOrder `json:"bid"`
}
