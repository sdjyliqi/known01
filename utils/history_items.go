package utils

var HistoryCalculatorCols = []map[string]string{
	map[string]string{
		"name": "日期",
		"key":  "event_day",
	},
	map[string]string{
		"name": "当日安装人数",
		"key":  "uv",
	},

	map[string]string{
		"name": "当日活跃人数",
		"key":  "history_uv",
	},
}

type HistoryDetail struct {
	EventDay  string   `json:"event_day" `
	Uv        int      `json:"uv"`
	HistoryUv int      `json:"history_uv"`
	UserIDs   []string `json:"-"`
	Detail    []string `json:"-"`
}
