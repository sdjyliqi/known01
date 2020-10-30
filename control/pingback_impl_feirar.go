package control

import (
	"github.com/sdjyliqi/feirars/models"
	"github.com/sdjyliqi/feirars/utils"
	"strings"
)

//GetFeirarDetailItems ...获取feirar 接口统计数据
func (pc *pingbackCenter) GetFeirarDetailItems(chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]models.FeirarDetailWeb, int64, error) {
	items, count, err := pc.feirarDetail.GetItemsByPage(pc.db, "", chn, pageID, pageCount, tsStart, tsEnd)
	if err != nil {
		return nil, 0, nil
	}
	webItems := make([]models.FeirarDetailWeb, 0, len(items))
	for _, v := range items {
		wItem := pc.feirarDetail.CovertWebItem(v)
		webItems = append(webItems, wItem)
	}
	return webItems, count, nil
}

//GetFeirarDetailCols ...前端显示的表头
func (pc *pingbackCenter) GetFeirarDetailCols() []map[string]string {
	return pc.feirarDetail.Cols()
}

func (pc *pingbackCenter) GetFeirarChart(chn string, tsStart, tsEnd int64) (*utils.ChartDetail, error) {
	return pc.feirarDetail.GetChartItems(pc.db, "", chn, tsStart, tsEnd)
}
func (pc *pingbackCenter) GetFeirarChannel(name string) ([]string, error) {
	item, err := pc.userBasic.GetUserBasic(pc.db, name)
	if err != nil {
		return nil, err
	}
	if item.Chn == "" {
		return pc.feirarDetail.GetAllChannels(pc.db, "")
	}
	chn_list := strings.Split(item.Chn, ",")
	return chn_list, nil
}

//GetFeirarDetailItems ...获取feirar 接口统计数据
func (pc *pingbackCenter) GetFeirarNewsDetailItems(chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]models.FeirarDetailWeb, int64, error) {
	items, count, err := pc.feirarDetail.GetItemsByPage(pc.db, "/api/FeiRarNews", chn, pageID, pageCount, tsStart, tsEnd)
	if err != nil {
		return nil, 0, nil
	}
	webItems := make([]models.FeirarDetailWeb, 0, len(items))
	for _, v := range items {
		wItem := pc.feirarDetail.CovertWebItem(v)
		webItems = append(webItems, wItem)
	}
	return webItems, count, nil
}

//GetFeirarDetailCols ...前端显示的表头
func (pc *pingbackCenter) GetFeirarNewsDetailCols() []map[string]string {
	return pc.feirarDetail.Cols()
}

func (pc *pingbackCenter) GetFeirarNewsChart(chn string, tsStart, tsEnd int64) (*utils.ChartDetail, error) {
	return pc.feirarDetail.GetChartItems(pc.db, "/api/FeiRarNews", chn, tsStart, tsEnd)
}
func (pc *pingbackCenter) GetFeirarNewsChannel(name string) ([]string, error) {
	item, err := pc.userBasic.GetUserBasic(pc.db, name)
	if err != nil {
		return nil, err
	}
	if item.Chn == "" {
		return pc.feirarDetail.GetAllChannels(pc.db, "/api/FeiRarNews")
	}
	chn_list := strings.Split(item.Chn, ",")
	return chn_list, nil
}

//GetFeirarUpdateDetailItems ...获取feirar中news 接口统计数据
func (pc *pingbackCenter) GetFeirarUpdateDetailItems(chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]models.FeirarDetailWeb, int64, error) {
	items, count, err := pc.feirarDetail.GetItemsByPage(pc.db, "/api/update", chn, pageID, pageCount, tsStart, tsEnd)
	if err != nil {
		return nil, 0, nil
	}
	webItems := make([]models.FeirarDetailWeb, 0, len(items))
	for _, v := range items {
		wItem := pc.feirarDetail.CovertWebItem(v)
		webItems = append(webItems, wItem)
	}
	return webItems, count, nil
}

//GetFeirarDetailCols ...前端显示的表头
func (pc *pingbackCenter) GetFeirarUpdateDetailCols() []map[string]string {
	return pc.feirarDetail.Cols()
}

func (pc *pingbackCenter) GetFeirarUpdateChart(chn string, tsStart, tsEnd int64) (*utils.ChartDetail, error) {
	return pc.feirarDetail.GetChartItems(pc.db, "/api/update", chn, tsStart, tsEnd)
}
func (pc *pingbackCenter) GetFeirarUpdateChannel(name string) ([]string, error) {
	item, err := pc.userBasic.GetUserBasic(pc.db, name)
	if err != nil {
		return nil, err
	}
	if item.Chn == "" {
		return pc.feirarDetail.GetAllChannels(pc.db, "/api/update")
	}
	chn_list := strings.Split(item.Chn, ",")
	return chn_list, nil
}
