package analysis

import (
	"github.com/montanaflynn/stats"
	"time"
	"praetoriansentry/btce-plot/pkg/data"
	"sort"
)

func BuildIndicators(tr data.TradeSet, bucketSeconds uint64) []data.Indicator {
	buckets := make(data.TradeBuckets, 0)
	for _, trade := range tr {
		_, ok := buckets[trade.Timestamp / bucketSeconds]
		if !ok {
			buckets[trade.Timestamp / bucketSeconds] = make(data.TradeSet, 0)
		}
		buckets[trade.Timestamp / bucketSeconds] = append(buckets[trade.Timestamp / bucketSeconds], trade)
	}

	indicators := make([]data.Indicator, 0)
	for _, bucket := range buckets {
		
		if len(bucket) < 1 {
			continue
		}

		// sort the bucket
		sort.Sort(bucket)

		prices := make([]float64, 0)
		amounts := make([]float64, 0)
		times := make([]uint64, 0)
		for _, trade := range bucket {
			prices = append(prices, trade.Price)
			amounts = append(amounts, trade.Amount)
			times = append(times, trade.Timestamp)
		}
		t := time.Unix(int64(times[0]), 0)
		h, _ := stats.Max(prices)
		l, _ := stats.Min(prices)
		v, _ := stats.Sum(amounts)
		ind := data.Indicator{
			Date: t.Format("2006-01-02T15:04:05"),
			Timestamp: times[0],
			Open: prices[0],
			High: h,
			Low: l,
			Close: prices[len(prices) - 1],
			Volume: v,
		}
		indicators = append(indicators, ind)
	}
	sort.Sort(data.IndicatorSet(indicators))
	return indicators
}
