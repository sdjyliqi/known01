package control

import (
	"encoding/json"
	"github.com/sdjyliqi/feirars/models"
	"github.com/sdjyliqi/feirars/utils"
	"strings"
)

//GetPreserveDetailItems ...获取留存统计相关数据
func (pc *pingbackCenter) GetPreserveDetailItems(chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]models.PreserveDetailWeb, int64, error) {
	items, count, err := pc.preserveDetail.GetItemsByPage(pc.db, chn, pageID, pageCount, tsStart, tsEnd)
	if err != nil {
		return nil, 0, nil
	}
	webItems := make([]models.PreserveDetailWeb, 0, len(items))
	for _, v := range items {
		wItem := pc.preserveDetail.CovertWebItem(v)
		webItems = append(webItems, wItem)
	}
	return webItems, count, nil
}

//GetPreserveDetailCols ...前端显示的表头
func (pc *pingbackCenter) GetPreserveDetailCols() []map[string]string {
	return pc.preserveDetail.Cols()
}

//GetPreserveChannel ...获取留存的所有渠道
func (pc *pingbackCenter) GetPreserveChannel(name string) ([]string, error) {
	item, err := pc.userBasic.GetUserBasic(pc.db, name)
	if err != nil {
		return nil, err
	}
	if item.Chn == "" {
		return pc.preserveDetail.GetAllChannels(pc.db)
	}
	chn_list := strings.Split(item.Chn, ",")
	return chn_list, nil
}

//GetPreserveChart...基于渠道的留存统计图
func (pc *pingbackCenter) GetPreserveChart(chn string, tsStart, tsEnd int64) (*utils.ChartDetail, error) {
	return pc.preserveDetail.GetChartItems(pc.db, chn, tsStart, tsEnd)
}

func (pc *pingbackCenter) GetPreserveHistoryCalculator(chn string, tsStart int64, days int) ([]*utils.HistoryDetail, error) {
	eventDayUserIDs := make([]string, 0)
	var historyItems []*utils.HistoryDetail
	items, err := pc.preserveDetail.GetItemsForHistory(pc.db, chn, tsStart, days)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, nil
	}
	if items[0].Detail != "" {
		err = json.Unmarshal([]byte(items[0].Detail), &eventDayUserIDs)
		if err != nil {
			return nil, err
		}
	}
	for _, v := range items {
		var uvList []string
		if v.Detail != "" {
			json.Unmarshal([]byte(v.Detail), &uvList)
		}
		historyItems = append(historyItems, &utils.HistoryDetail{
			EventDay:  v.EventTime.Format(utils.DayTime),
			Uv:        v.Uv,
			HistoryUv: len(utils.TwoSliceIntersect(eventDayUserIDs, uvList)),
			UserIDs:   eventDayUserIDs,
			Detail:    uvList,
		})
	}
	return historyItems, nil
}
