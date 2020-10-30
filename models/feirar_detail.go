package models

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/golang/glog"
	"github.com/sdjyliqi/feirars/utils"
	"strings"
	"time"
)

type FeirarDetail struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	EventDay   time.Time `json:"event_day" xorm:"not null comment('事件日期') DATETIME"`
	Channel    string    `json:"channel" xorm:"VARCHAR(64)"`
	EventKey   string    `json:"event_key" xorm:"comment('时间名称') VARCHAR(128)"`
	Pv         int       `json:"pv" xorm:"comment('PV用户数') INT(11)"`
	Uv         int       `json:"uv" xorm:"comment('UV用户数') INT(11)"`
	LastUpdate time.Time `json:"last_update" xorm:"not null comment('更新数据时间') DATETIME"`
	Detail     string    `json:"detail" xorm:"TEXT"`
}

func (t FeirarDetail) TableName() string {
	return "feirar_detail"
}

type FeirarDetailWeb struct {
	Id         int    `json:"id" `
	EventDay   string `json:"event_day" `
	Channel    string `json:"channel" `
	EventKey   string `json:"event_key"`
	Pv         string `json:"pv" `
	Uv         string `json:"uv" `
	LastUpdate string `json:"last_update"`
	Detail     string `json:"detail"`
}

func (t FeirarDetail) CovertWebItem(item *FeirarDetail) FeirarDetailWeb {
	webItem := FeirarDetailWeb{
		EventDay:   item.EventDay.Format(utils.DayTime),
		Channel:    item.Channel,
		EventKey:   item.EventKey,
		Pv:         fmt.Sprintf("%d", item.Pv),
		Uv:         fmt.Sprintf("%d", item.Uv),
		LastUpdate: item.LastUpdate.Format(utils.FullTime),
	}
	return webItem
}
func (t FeirarDetail) Cols() []map[string]string {
	var cols []map[string]string
	colEventDay := map[string]string{
		"name":  "日期",
		"key":   "event_day",
		"click": "0",
	}
	cols = append(cols, colEventDay)

	colClientChannel := map[string]string{
		"name":  "渠道",
		"key":   "channel",
		"click": "1",
	}
	cols = append(cols, colClientChannel)

	colEventKey := map[string]string{
		"name":  "api名称",
		"key":   "event_key",
		"click": "0",
	}
	cols = append(cols, colEventKey)

	colPv := map[string]string{
		"name":  "pv",
		"key":   "pv",
		"click": "0",
	}
	cols = append(cols, colPv)

	colUv := map[string]string{
		"name":  "uv",
		"key":   "uv",
		"click": "0",
	}
	cols = append(cols, colUv)
	colLastUpdate := map[string]string{
		"name":  "更新时间",
		"key":   "last_update",
		"click": "0",
	}
	cols = append(cols, colLastUpdate)
	return cols
}

//GetAllChannels ...获取所有渠道
func (t FeirarDetail) GetAllChannels(client *xorm.Engine, api string) ([]string, error) {
	var items []*FeirarDetail
	var channels []string
	err := client.Distinct("channel").OrderBy("channel").Find(&items)
	if err != nil {
		glog.Errorf("[mysql]Get the channel  from table %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	for _, v := range items {
		channels = append(channels, v.Channel)
	}
	return channels, nil
}
func (t FeirarDetail) GetItemsByPage(client *xorm.Engine, api, chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]*FeirarDetail, int64, error) {
	timeTS, timeTE := utils.ConvertToTime(tsStart), utils.ConvertToTime(tsEnd)
	var items []*FeirarDetail
	item := FeirarDetail{}
	session := client.Where("event_day>=?", timeTS).And("event_day<=?", timeTE)
	if chn != "" {
		chnList := utils.ChannelList(chn)
		session = session.In("channel", chnList)
	}
	if api != "" {
		session = session.Where("event_key =?", api)
	}
	session = session.Desc("event_day")
	if pageCount > 0 {
		session = session.Limit(pageCount, pageCount*(pageID-1))
	}
	err := session.Find(&items)
	if err != nil {
		glog.Errorf("[mysql]Get the items for from table %s failed,err:%+v", t.TableName(), err)
		return nil, 0, err
	}

	session = client.Where("event_day>=?", timeTS).And("event_day<=?", timeTE)
	if chn != "" {
		chnList := utils.ChannelList(chn)
		session = session.In("channel", chnList)
	}
	if api != "" {
		session = session.Where("event_key =?", api)
	}
	cnt, err := session.Count(item)
	if err != nil {
		glog.Errorf("[mysql]Get the count of items for from table %s failed,err:%+v", t.TableName(), err)
		return nil, 0, err
	}
	return items, cnt, nil
}

func (t FeirarDetail) GetChartItems(client *xorm.Engine, api, chn string, tsStart, tsEnd int64) (*utils.ChartDetail, error) {
	chartXvalue := []string{}
	chartXDic := map[string]bool{}
	timeTS, timeTE := utils.ConvertToTime(tsStart), utils.ConvertToTime(tsEnd)
	var items []*FeirarDetail
	session := client.Where("event_day>=?", timeTS).And("event_day<=?", timeTE)
	if chn != "" {
		chnList := utils.ChannelList(chn)
		session = session.In("channel", chnList)
	}
	if api != "" {
		session = session.Where("event_key =?", api)
	}
	err := session.OrderBy("event_day,channel").
		Find(&items)
	if err != nil {
		glog.Errorf("[mysql]Get the items for from table %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}

	chartPVValue := map[string]utils.ChartLineSeries{}
	chartUVValue := map[string]utils.ChartLineSeries{}
	for _, v := range items {
		//时间正序计算x轴的日期
		xValue := v.EventDay.Format(utils.DayTime)
		_, ok := chartXDic[xValue]
		if !ok {
			chartXDic[xValue] = true
			chartXvalue = append(chartXvalue, xValue)
		}
		//计算pv chart数据
		idx := fmt.Sprintf("%s%s%s", v.Channel, utils.SepChar, v.EventKey)
		val, ok := chartPVValue[idx]
		//pv chart
		if ok {
			val.Data = append(val.Data, float64(v.Pv))
			val.EventTime = append(val.EventTime, xValue)
			chartPVValue[idx] = val
		} else {
			chartPVValue[idx] = utils.ChartLineSeries{
				Data:      []float64{float64(v.Pv)},
				EventTime: []string{xValue},
			}
		}
		//计算UV chart
		val, ok = chartUVValue[idx]
		//pv chart
		if ok {
			val.Data = append(val.Data, float64(v.Pv))
			val.EventTime = append(val.EventTime, xValue)
			chartUVValue[idx] = val
		} else {
			chartUVValue[idx] = utils.ChartLineSeries{
				Data:      []float64{float64(v.Pv)},
				EventTime: []string{xValue},
			}
		}
	}

	var chartYlines []utils.ChartSeriesYValue
	for k, v := range chartPVValue {
		infos := strings.Split(k, utils.SepChar)
		lineTitle := fmt.Sprintf("渠道%s接口%sPV趋势图", infos[0], infos[1])
		chartYLine := utils.ChartSeriesYValue{
			Name:      lineTitle,
			ChartType: "line",
			Data:      v.Data,
			EventTime: v.EventTime,
		}
		chartYlines = append(chartYlines, chartYLine)
	}
	for k, v := range chartUVValue {
		infos := strings.Split(k, utils.SepChar)
		//chan_
		lineTitle := fmt.Sprintf("渠道%s接口%sUV趋势图", infos[0], infos[1])
		chartYLine := utils.ChartSeriesYValue{
			Name:      lineTitle,
			ChartType: "line",
			Data:      v.Data,
			EventTime: v.EventTime,
		}
		chartYlines = append(chartYlines, chartYLine)
	}
	chartItems := &utils.ChartDetail{
		XAxis:  chartXvalue,
		Series: chartYlines,
	}
	return utils.ChartItemsMend(chartItems), err
}
