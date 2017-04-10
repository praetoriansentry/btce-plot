package data

type Trade struct {
	Type string `json:"type"`
	Price float64 `json:"price"`
	Amount float64 `json:"amount"`
	Tid uint64 `json:"tid"`
	Timestamp uint64 `json:"timestamp"`
}

type TradeResponse map[string][]Trade
