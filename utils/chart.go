package utils

import "fmt"

type ChartSeriesYValue struct {
	Name      string    `json:"name" `
	ChartType string    `json:"chart_type" `
	Data      []float64 `json:"data" `
	EventTime []string  `json:"event_time" `
}
type ChartDetail struct {
	XAxis  []string            `json:"x_axis"`
	Series []ChartSeriesYValue `json:"series"`
}

type ChartLineSeries struct {
	Data      []float64 `json:"data" `
	EventTime []string  `json:"event_time" `
}

func SliceInsertValueAtPosition(s []float64, index int, value float64) []float64 {
	rear := s[index:]
	newSlice := []float64{}
	newSlice = append(newSlice, s[:index]...)
	newSlice = append(newSlice, value)
	newSlice = append(newSlice, rear...)
	return newSlice
}

//根据日期修补数据，保证x轴上的数据和line的Y轴数据个事一致
func ChartItemsMend(value *ChartDetail) *ChartDetail {
	if value == nil {
		return nil
	}
	if value.XAxis == nil || len(value.XAxis) == 0 {
		return value
	}

	if value.Series == nil || len(value.Series) == 0 {
		return value
	}

	xDic := map[string]bool{}
	for k, v := range value.Series {
		for _, vv := range v.EventTime {
			idx := fmt.Sprintf("%d_%v", k, vv)
			xDic[idx] = true
		}
	}

	for k, _ := range value.Series {
		for kk, vv := range value.XAxis {
			idx := fmt.Sprintf("%d_%v", k, vv)
			_, ok := xDic[idx]
			if !ok {
				newSlice := SliceInsertValueAtPosition(value.Series[k].Data, kk, 0.0)
				value.Series[k].Data = newSlice
			}
		}
	}
	return value
}
