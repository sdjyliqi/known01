package control

import (
	"github.com/sdjyliqi/feirars/models"
	"github.com/sdjyliqi/feirars/utils"
	"strings"
)

//GetUdtrstDetailItems ...获取客户端咨询弹窗相关接口
func (pc *pingbackCenter) GetUdtrstDetailItems(chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]models.UdtrstDetailWeb, int64, error) {
	items, count, err := pc.udtrstDetail.GetItemsByPage(pc.db, chn, pageID, pageCount, tsStart, tsEnd)
	if err != nil {
		return nil, 0, nil
	}
	webItems := make([]models.UdtrstDetailWeb, 0, len(items))
	for _, v := range items {
		wItem := pc.udtrstDetail.CovertWebItem(v)
		webItems = append(webItems, wItem)
	}
	return webItems, count, nil
}

//GetUdtrstDetailCols ...前端显示的表头
func (pc *pingbackCenter) GetUdtrstDetailCols() []map[string]string {
	return pc.udtrstDetail.Cols()
}

func (pc *pingbackCenter) GetUdtrstChannel(name string) ([]string, error) {
	item, err := pc.userBasic.GetUserBasic(pc.db, name)
	if err != nil {
		return nil, err
	}
	if item.Chn == "" {
		return pc.udtrstDetail.GetAllChannels(pc.db)
	}
	chnList := strings.Split(item.Chn, ",")
	return chnList, nil
}

//GetUdtrstChart ...获取渠道弹窗的趋势图数据
func (pc *pingbackCenter) GetUdtrstChart(chn string, tsStart, tsEnd int64) (*utils.ChartDetail, error) {
	return pc.udtrstDetail.GetChartItems(pc.db, chn, tsStart, tsEnd)
}
