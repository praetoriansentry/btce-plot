package data

type Trade struct {
	Type      string  `json:"type"`
	Price     float64 `json:"price"`
	Amount    float64 `json:"amount"`
	Tid       int     `json:"tid"`
	Timestamp int     `json:"timestamp"`
}

type Indicator struct {
	Date      string
	Open      float64
	High      float64
	Low       float64
	Close     float64
	Volume    float64
	Timestamp int
}

type TradeResponse map[string]TradeSet
type TradeBuckets map[int]TradeSet
type TradeIndicators map[int][]Indicator

type TradeSet []Trade

func (t TradeSet) Len() int           { return len(t) }
func (t TradeSet) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t TradeSet) Less(i, j int) bool { return t[i].Timestamp < t[j].Timestamp }

type IndicatorSet []Indicator

func (is IndicatorSet) Len() int           { return len(is) }
func (is IndicatorSet) Swap(i, j int)      { is[i], is[j] = is[j], is[i] }
func (is IndicatorSet) Less(i, j int) bool { return is[i].Timestamp < is[j].Timestamp }
