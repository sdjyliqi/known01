package control

import (
	"github.com/sdjyliqi/feirars/models"
	"github.com/sdjyliqi/feirars/utils"
	"strings"
)

//GetNewsDetailItems ...获取客户端咨询弹窗相关接口
func (pc *pingbackCenter) GetNewsDetailItems(chn string, pageID, pageCount int, tsStart, tsEnd int64, eventKey string) ([]models.NewsDetailWeb, int64, error) {
	items, count, err := pc.newsDetail.GetItemsByPage(pc.db, chn, pageID, pageCount, tsStart, tsEnd, eventKey)
	if err != nil {
		return nil, 0, nil
	}
	webItems := make([]models.NewsDetailWeb, 0, len(items))
	for _, v := range items {
		wItem := pc.newsDetail.CovertWebItem(v)
		webItems = append(webItems, wItem)
	}
	return webItems, count, nil
}

//GetNewsDetailCols ...前端显示的表头
func (pc *pingbackCenter) GetNewsDetailCols() []map[string]string {
	return pc.newsDetail.Cols()
}

func (pc *pingbackCenter) GetNewsChannel(name string,eventKey string) ([]string, error) {
	item, err := pc.userBasic.GetUserBasic(pc.db, name)
	if err != nil {
		return nil, err
	}
	if item.Chn == "" {
		return pc.newsDetail.GetAllChannels(pc.db,eventKey)
	}

	chn_list := strings.Split(item.Chn, ",")
	return chn_list, nil
}

//GetNewsChart ...获取渠道弹窗的趋势图数据
func (pc *pingbackCenter) GetNewsChart(chn string, tsStart, tsEnd int64,eventKey string) (*utils.ChartDetail, error) {
	return pc.newsDetail.GetChartItems(pc.db, chn, tsStart, tsEnd,eventKey)
}
