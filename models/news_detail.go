package models

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/golang/glog"
	"github.com/sdjyliqi/feirars/utils"
	"strings"
	"time"
)

type NewsDetail struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	EventDay   time.Time `json:"event_day" xorm:"not null comment('事件日期') DATETIME"`
	Channel    string    `json:"channel" xorm:"VARCHAR(64)"`
	Pv         int       `json:"pv" xorm:"comment('PV用户数') INT(11)"`
	Uv         int       `json:"uv" xorm:"comment('UV用户数') INT(11)"`
	LastUpdate time.Time `json:"last_update" xorm:"not null comment('更新数据时间') DATETIME"`
	Detail     string    `json:"detail" xorm:"TEXT"`
	ShowPv     int       `json:"show_pv" xorm:"comment('show PV用户数') INT(11)"`
	ShowUv     int       `json:"show_uv" xorm:"comment('show UV用户数') INT(11)"`
	ClickPv    int       `json:"click_pv" xorm:"comment('click PV用户数') INT(11)"`
	ClickUv    int       `json:"click_uv" xorm:"comment('click UV用户数') INT(11)"`
}

type NewsDetailWeb struct {
	Id         int    `json:"id" `
	EventDay   string `json:"event_day" `
	Channel    string `json:"channel" `
	ShowPv     string `json:"show_pv"`
	ShowUv     string `json:"show_uv" `
	ClickPv    string `json:"click_pv" `
	ClickUv    string `json:"click_uv" `
	LastUpdate string `json:"last_update" `
}

func (t NewsDetail) TableName() string {
	return "news_detail"
}

func (t NewsDetail) CovertWebItem(item *NewsDetail) NewsDetailWeb {
	webItem := NewsDetailWeb{
		EventDay:   item.EventDay.Format(utils.DayTime),
		Channel:    item.Channel,
		ShowPv:     fmt.Sprintf("%d", item.ShowPv),
		ShowUv:     fmt.Sprintf("%d", item.ShowUv),
		ClickPv:    fmt.Sprintf("%d", item.ClickPv),
		ClickUv:    fmt.Sprintf("%d", item.ClickUv),
		LastUpdate: item.LastUpdate.Format(utils.FullTime),
	}
	return webItem
}

//Cols ...用户web显示使用
func (t NewsDetail) Cols() []map[string]string {
	var cols []map[string]string
	colEventDay := map[string]string{
		"name": "日期",
		"key":  "event_day",
	}
	cols = append(cols, colEventDay)

	colClientChannel := map[string]string{
		"name":  "渠道",
		"key":   "channel",
		"click": "1",
	}
	cols = append(cols, colClientChannel)

	colShowPv := map[string]string{
		"name":  "show pv",
		"key":   "show_pv",
		"click": "0",
	}
	cols = append(cols, colShowPv)

	colShowUv := map[string]string{
		"name":  "show uv",
		"key":   "show_uv",
		"click": "0",
	}
	cols = append(cols, colShowUv)

	colClickPv := map[string]string{
		"name":  "click pv",
		"key":   "click_pv",
		"click": "0",
	}
	cols = append(cols, colClickPv)

	colClickUv := map[string]string{
		"name":  "click uv",
		"key":   "click_uv",
		"click": "0",
	}
	cols = append(cols, colClickUv)

	colLastUpdate := map[string]string{
		"name":  "更新时间",
		"key":   "last_update",
		"click": "0",
	}
	cols = append(cols, colLastUpdate)
	return cols
}

//GetAllChannels ...获取所有渠道
func (t NewsDetail) GetAllChannels(client *xorm.Engine, eventKey string) ([]string, error) {
	var items []*NewsDetail
	var channels []string
	err := client.Distinct("channel").And(fmt.Sprintf("event_type ='%s'", eventKey)).OrderBy("channel").Find(&items)
	if err != nil {
		glog.Errorf("[mysql]Get the channel  from table %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	for _, v := range items {
		channels = append(channels, v.Channel)
	}
	return channels, nil
}

func (t NewsDetail) GetItemsByPage(client *xorm.Engine, chn string, pageID, pageCount int, tsStart, tsEnd int64, eventKey string) ([]*NewsDetail, int64, error) {
	var items []*NewsDetail
	item := &NewsDetail{}
	timeTS, timeTE := utils.ConvertToTime(tsStart), utils.ConvertToTime(tsEnd)
	session := client.Where("event_day>=?", timeTS).And("event_day<=?", timeTE).And(fmt.Sprintf("event_type ='%s'", eventKey))

	if chn != "" {
		chnList := utils.ChannelList(chn)
		session = session.In("channel", chnList)
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

	session = client.Where("event_day>=?", timeTS).And("event_day<=?", timeTE).And(fmt.Sprintf("event_type ='%s'", eventKey))
	if chn != "" {
		chnList := utils.ChannelList(chn)
		session = session.In("channel", chnList)
	}
	cnt, err := session.Count(item)
	if err != nil {
		glog.Errorf("[mysql]Get the count of items for from table %s failed,err:%+v", t.TableName(), err)
		return nil, 0, err
	}
	return items, cnt, nil
}

func (t NewsDetail) GetChartItems(client *xorm.Engine, chn string, tsStart, tsEnd int64, eventKey string) (*utils.ChartDetail, error) {
	chartXvalue := make([]string, 0)
	chartXDic := map[string]bool{}
	timeTS, timeTE := utils.ConvertToTime(tsStart), utils.ConvertToTime(tsEnd)
	var items []*NewsDetail
	session := client.Where("event_day>=?", timeTS).And("event_day<=?", timeTE).And(fmt.Sprintf("event_type ='%s'", eventKey))
	if chn != "" {
		chnList := utils.ChannelList(chn)
		session = session.In("channel", chnList)
	}
	err := session.OrderBy("event_day,channel").
		Find(&items)
	if err != nil {
		glog.Errorf("[mysql]Get the items for from table %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}

	chartShowPVValue := map[string]utils.ChartLineSeries{}
	chartShowUVValue := map[string]utils.ChartLineSeries{}
	chartClickPVValue := map[string]utils.ChartLineSeries{}
	chartClickUVValue := map[string]utils.ChartLineSeries{}
	for _, v := range items {
		//时间正序计算x轴的日期
		xValue := v.EventDay.Format(utils.DayTime)
		_, ok := chartXDic[xValue]
		if !ok {
			chartXDic[xValue] = true
			chartXvalue = append(chartXvalue, xValue)
		}

		idx := fmt.Sprintf("%s%s%s", v.Channel, utils.SepChar, "-")
		//计算show PV chart数据
		val, ok := chartShowPVValue[idx]
		if ok {
			val.Data = append(val.Data, float64(v.ShowPv))
			val.EventTime = append(val.EventTime, xValue)
			chartShowPVValue[idx] = val
		} else {
			chartShowPVValue[idx] = utils.ChartLineSeries{
				Data:      []float64{float64(v.ShowPv)},
				EventTime: []string{xValue},
			}
		}
		//计算show UV chart
		val, ok = chartShowUVValue[idx]
		if ok {
			val.Data = append(val.Data, float64(v.ShowUv))
			val.EventTime = append(val.EventTime, xValue)
			chartShowUVValue[idx] = val
		} else {
			chartShowUVValue[idx] = utils.ChartLineSeries{
				Data:      []float64{float64(v.ShowUv)},
				EventTime: []string{xValue},
			}
		}

		//计算click PV chart
		val, ok = chartClickPVValue[idx]
		if ok {
			val.Data = append(val.Data, float64(v.ClickPv))
			val.EventTime = append(val.EventTime, xValue)
			chartClickPVValue[idx] = val
		} else {
			chartClickPVValue[idx] = utils.ChartLineSeries{
				Data:      []float64{float64(v.ClickPv)},
				EventTime: []string{xValue},
			}
		}
		//计算click PV chart
		val, ok = chartClickUVValue[idx]
		if ok {
			val.Data = append(val.Data, float64(v.ClickUv))
			val.EventTime = append(val.EventTime, xValue)
			chartClickUVValue[idx] = val
		} else {
			chartClickUVValue[idx] = utils.ChartLineSeries{
				Data:      []float64{float64(v.ClickUv)},
				EventTime: []string{xValue},
			}
		}

	}
	var chartYlines []utils.ChartSeriesYValue
	//添加第一条线
	for k, v := range chartShowPVValue {
		infos := strings.Split(k, utils.SepChar)
		lineTitle := fmt.Sprintf("%s渠道show-PV趋势图", infos[0])
		chartYLine := utils.ChartSeriesYValue{
			Name:      lineTitle,
			ChartType: "line",
			Data:      v.Data,
			EventTime: v.EventTime,
		}
		chartYlines = append(chartYlines, chartYLine)
	}
	//添加第二条线
	for k, v := range chartShowUVValue {
		infos := strings.Split(k, utils.SepChar)
		//chan_
		lineTitle := fmt.Sprintf("%s渠道show-UV趋势图", infos[0])
		chartYLine := utils.ChartSeriesYValue{
			Name:      lineTitle,
			ChartType: "line",
			Data:      v.Data,
			EventTime: v.EventTime,
		}
		chartYlines = append(chartYlines, chartYLine)
	}

	//添加第三条线
	for k, v := range chartClickPVValue {
		infos := strings.Split(k, utils.SepChar)
		//chan_
		lineTitle := fmt.Sprintf("%s渠道click-UV趋势图", infos[0])
		chartYLine := utils.ChartSeriesYValue{
			Name:      lineTitle,
			ChartType: "line",
			Data:      v.Data,
			EventTime: v.EventTime,
		}
		chartYlines = append(chartYlines, chartYLine)
	}
	//添加第四条线
	for k, v := range chartClickUVValue {
		infos := strings.Split(k, utils.SepChar)
		//chan_
		lineTitle := fmt.Sprintf("%s渠道click-UV趋势图", infos[0])
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
