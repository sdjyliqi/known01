package models

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/golang/glog"
	"github.com/sdjyliqi/feirars/utils"
	"strings"
	"time"
)

type UninstallDetail struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	EventDay   time.Time `json:"event_day" xorm:"not null comment('事件日期') DATETIME"`
	Channel    string    `json:"channel" xorm:"VARCHAR(64)"`
	Pv         int       `json:"pv" xorm:"comment('PV用户数') INT(11)"`
	Uv         int       `json:"uv" xorm:"comment('UV用户数') INT(11)"`
	LastUpdate time.Time `json:"last_update" xorm:"not null comment('更新数据时间') DATETIME"`
	Detail     string    `json:"detail" xorm:"TEXT"`
}

type UninstallDetailWeb struct {
	Id         string `json:"id" `
	EventDay   string `json:"event_day" `
	Channel    string `json:"channel" `
	Pv         string `json:"pv" `
	Uv         string `json:"uv" `
	LastUpdate string `json:"last_update" `
	Detail     string `json:"detail" `
}

func (t UninstallDetail) TableName() string {
	return "uninstall_detail"
}

func (t UninstallDetail) CovertWebItem(item *UninstallDetail) UninstallDetailWeb {
	webItem := UninstallDetailWeb{
		EventDay:   item.EventDay.Format(utils.DayTime),
		Channel:    item.Channel,
		Pv:         fmt.Sprintf("%d", item.Pv),
		Uv:         fmt.Sprintf("%d", item.Uv),
		LastUpdate: item.LastUpdate.Format(utils.FullTime),
	}
	return webItem
}

//Cols ...用户web显示使用
func (t UninstallDetail) Cols() []map[string]string {
	var cols []map[string]string
	colEventDay := map[string]string{
		"name":  "日期",
		"key":   "event_day",
		"click": "0",
	}
	cols = append(cols, colEventDay)

	colChannel := map[string]string{
		"name":  "渠道",
		"key":   "channel",
		"click": "1",
	}
	cols = append(cols, colChannel)

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
func (t UninstallDetail) GetAllChannels(client *xorm.Engine) ([]string, error) {
	var items []*UninstallDetail
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

func (t UninstallDetail) GetItemsByPage(client *xorm.Engine, chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]*UninstallDetail, int64, error) {
	var items []*UninstallDetail
	item := &UninstallDetail{}
	timeTS, timeTE := utils.ConvertToTime(tsStart), utils.ConvertToTime(tsEnd)
	//按照页获取
	session := client.
		Where("event_day>=?", timeTS).And("event_day<=?", timeTE).
		Desc("event_day")

	if chn != "" {
		chnList := utils.ChannelList(chn)
		session = session.In("channel", chnList)
	}
	if pageCount > 0 {
		session = session.Limit(pageCount, pageCount*(pageID-1))
	}
	err := session.Find(&items)
	if err != nil {
		glog.Errorf("[mysql]Get the items for from table %s failed,err:%+v", t.TableName(), err)
		return nil, 0, err
	}
	//获取总条数
	session = client.Where("event_day>=?", timeTS).And("event_day<=?", timeTE)
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

func (t UninstallDetail) GetChartItems(client *xorm.Engine, chn string, tsStart, tsEnd int64) (*utils.ChartDetail, error) {
	chartXvalue := []string{}
	chartXDic := map[string]bool{}
	timeTS, timeTE := utils.ConvertToTime(tsStart), utils.ConvertToTime(tsEnd)
	var items []*UninstallDetail
	session := client.Where("event_day>=?", timeTS).And("event_day<=?", timeTE)
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
		idx := fmt.Sprintf("%s%s%s", v.Channel, utils.SepChar, "-")
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
			val.Data = append(val.Data, float64(v.Uv))
			val.EventTime = append(val.EventTime, xValue)
			chartUVValue[idx] = val
		} else {
			chartUVValue[idx] = utils.ChartLineSeries{
				Data:      []float64{float64(v.Uv)},
				EventTime: []string{xValue},
			}
		}
	}

	var chartYlines []utils.ChartSeriesYValue
	for k, v := range chartPVValue {
		infos := strings.Split(k, utils.SepChar)
		lineTitle := fmt.Sprintf("渠道%sPV趋势图", infos[0])
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
		lineTitle := fmt.Sprintf("渠道%sUV趋势图", infos[0])
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
