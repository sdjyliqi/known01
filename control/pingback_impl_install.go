package control

import (
	"encoding/json"
	"github.com/sdjyliqi/feirars/models"
	"github.com/sdjyliqi/feirars/utils"
	"strings"
)

//获取安装统计数据
func (pc *pingbackCenter) GetInstallDetailItems(chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]models.InstallDetailWeb, int64, error) {
	items, count, err := pc.installDetail.GetItemsByPage(pc.db, chn, pageID, pageCount, tsStart, tsEnd)
	if err != nil {
		return nil, 0, nil
	}
	webItems := make([]models.InstallDetailWeb, 0, len(items))
	for _, v := range items {
		wItem := pc.installDetail.CovertWebItem(v)
		webItems = append(webItems, wItem)
	}
	return webItems, count, nil
}

//GetInstallDetailCols ...
func (pc *pingbackCenter) GetInstallDetailCols() []map[string]string {
	return pc.installDetail.Cols()
}

//GetInstallChannel ...获取安装统计项的所有渠道列表
func (pc *pingbackCenter) GetInstallChannel(name string) ([]string, error) {
	item, err := pc.userBasic.GetUserBasic(pc.db, name)
	if err != nil {
		return nil, err
	}
	if item.Chn == "" {
		return pc.installDetail.GetAllChannels(pc.db)
	}
	chn_list := strings.Split(item.Chn, ",")
	return chn_list, nil
}

//GetInstallNewsChart ...获取渠道统计安装的趋势图数据
func (pc *pingbackCenter) GetInstallChart(chn string, tsStart, tsEnd int64) (*utils.ChartDetail, error) {
	return pc.installDetail.GetChartItems(pc.db, chn, tsStart, tsEnd)
}

//	items, err := testActiveDetail.GetItemsForHistory(testutil.TestMysql, "all",day.Unix(),5)

//GetHistoryCalculator ...基于历史留存数据
func (pc *pingbackCenter) GetInstallHistoryCalculator(chn string, tsStart int64, days int) ([]*utils.HistoryDetail, error) {
	eventDayUserIDs := make([]string, 0)
	var historyItems []*utils.HistoryDetail
	items, err := pc.installDetail.GetItemsForHistory(pc.db, chn, tsStart, days)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, nil
	}
	if items[0].PreserveDetail.Detail != "" {
		err = json.Unmarshal([]byte(items[0].InstallDetail.Detail), &eventDayUserIDs)
		if err != nil {
			return nil, err
		}
	}
	for _, v := range items {
		var uvList []string
		if v.PreserveDetail.Detail != "" {
			json.Unmarshal([]byte(v.PreserveDetail.Detail), &uvList)
		}
		uvList = utils.SliceUnique(uvList)
		historyItems = append(historyItems, &utils.HistoryDetail{
			EventDay:  v.InstallDetail.EventDay.Format(utils.DayTime),
			Uv:        v.InstallDetail.Uv,
			HistoryUv: len(utils.TwoSliceIntersect(eventDayUserIDs, uvList)),
			UserIDs:   eventDayUserIDs,
			Detail:    uvList,
		})
	}
	return historyItems, nil
}
