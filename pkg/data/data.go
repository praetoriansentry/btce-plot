package data

type Trade struct {
	Type string `json:"type"`
	Price float64 `json:"price"`
	Amount float64 `json:"amount"`
	Tid uint64 `json:"tid"`
	Timestamp uint64 `json:"timestamp"`
}

type Indicator struct {
	Date string
	Open float64
	High float64
	Low float64
	Close float64
	Volume float64
}

type TradeResponse map[string]TradeSet
type TradeBuckets map[uint64]TradeSet
type TradeIndicators map[uint64][]Indicator

type TradeSet []Trade

func (t TradeSet) Len() int {return len(t)}
func (t TradeSet) Swap(i, j int)  { t[i], t[j] = t[j], t[i] }
func (t TradeSet) Less(i, j int) bool {return t[i].Timestamp < t[j].Timestamp }
