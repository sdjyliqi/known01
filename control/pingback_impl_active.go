package control

import (
	"github.com/sdjyliqi/feirars/models"
	"github.com/sdjyliqi/feirars/utils"
	"strings"
)

//GetActiveDetailItems ...获取客户的激活方式统计数据
func (pc *pingbackCenter) GetActiveDetailItems(chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]models.ActiveDetailWeb, int64, error) { //按照页面获取统计激活方式数据
	items, count, err := pc.activeDetail.GetItemsByPage(pc.db, chn, pageID, pageCount, tsStart, tsEnd)
	if err != nil {
		return nil, 0, nil
	}
	webItems := make([]models.ActiveDetailWeb, 0, len(items))
	for _, v := range items {
		wItem := pc.activeDetail.CovertWebItem(v)
		webItems = append(webItems, wItem)
	}
	return webItems, count, nil
}

func (pc *pingbackCenter) GetActiveDetailCols() []map[string]string {
	return pc.activeDetail.Cols()
}

func (pc *pingbackCenter) GetActiveChannel(name string) ([]string, error) {
	chn, err := pc.UserChn(name)
	if err != nil {
		return []string{}, nil
	}
	if chn == "" {
		return pc.activeDetail.GetAllChannels(pc.db)
	}
	chn_list := strings.Split(chn, ",")
	return chn_list, nil
}

func (pc *pingbackCenter) GetActiveChart(chn string, tsStart, tsEnd int64) (*utils.ChartDetail, error) {
	return pc.activeDetail.GetChartItems(pc.db, chn, tsStart, tsEnd)
}
