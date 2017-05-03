package archive

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"praetoriansentry/btce-plot/pkg/data"
	"sort"
)

func WriteToFile(ts data.TradeSet, fileName string) error {
	originalTs, err := loadTradeFile(fileName)
	if err != nil {
		return err
	}
	mergedData := mergeSets(ts, originalTs)
	return saveToFile(mergedData, fileName)
}

func saveToFile(ts data.TradeSet, fileName string) error {
	data, err := json.Marshal(ts)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fileName, data, os.ModePerm)
}

func mergeSets(ts1, ts2 data.TradeSet) data.TradeSet {
	if ts1 == nil && ts2 == nil {
		return nil
	}
	if ts1 == nil {
		return ts2
	}
	if ts2 == nil {
		return ts1
	}
	tidMap := make(map[int]data.Trade, 0)
	for _, v := range ts1 {
		tidMap[v.Tid] = v
	}
	for _, v := range ts2 {
		tidMap[v.Tid] = v
	}
	var trades data.TradeSet = make(data.TradeSet, 0)
	for _, t := range tidMap {
		trades = append(trades, t)

	}
	sort.Sort(trades)
	return trades
}

func loadTradeFile(fileName string) (data.TradeSet, error) {
	var ts data.TradeSet
	file, err := os.Open(fileName) // For read access.
	if err != nil {
		return ts, nil
	}

	defer file.Close()

	rawdata, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(rawdata, &ts)
	if err != nil {
		return nil, err
	}

	return ts, nil
}
