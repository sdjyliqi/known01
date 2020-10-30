package models

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/golang/glog"
	"github.com/sdjyliqi/feirars/utils"
	"strings"
	"time"
)

type PreserveDetail struct {
	Id               int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	EventTime        time.Time `json:"event_time" xorm:"not null DATETIME"`
	Channel          string    `json:"channel" xorm:"VARCHAR(64)"`
	Uv               int       `json:"uv" xorm:"INT(11)"`
	NewUv            int       `json:"new_uv" xorm:"INT(11)"`
	PassedWeekActive int       `json:"passed_week_active" xorm:"INT(11)"`
	Day1Active       int       `json:"day1_active" xorm:"INT(11)"`
	Day2Active       int       `json:"day2_active" xorm:"INT(11)"`
	Day3Active       int       `json:"day3_active" xorm:"INT(11)"`
	Day4Active       int       `json:"day4_active" xorm:"INT(11)"`
	Day5Active       int       `json:"day5_active" xorm:"INT(11)"`
	Day6Active       int       `json:"day6_active" xorm:"INT(11)"`
	WeekActive       int       `json:"week_active" xorm:"INT(11)"`
	Detail           string    `json:"detail" xorm:"TEXT"`
	LastUpdate       time.Time `json:"last_update" xorm:"DATETIME"`
}
type PreserveDetailWeb struct {
	Id               string `json:"id" `
	EventTime        string `json:"event_time" `
	Channel          string `json:"channel" `
	Uv               string `json:"uv" `
	NewUv            string `json:"new_uv" `
	PassedWeekActive string `json:"passed_week_active"`
	Day1Active       string `json:"day1_active"`
	Day2Active       string `json:"day2_active" `
	Day3Active       string `json:"day3_active" `
	Day4Active       string `json:"day4_active"`
	Day5Active       string `json:"day5_active" `
	Day6Active       string `json:"day6_active"`
	WeekActive       string `json:"week_active" `
	Detail           string `json:"detail" `
	LastUpdate       string `json:"last_update" `
}

func (t PreserveDetail) TableName() string {
	return "preserve_detail"
}

func (t PreserveDetail) CovertWebItem(item *PreserveDetail) PreserveDetailWeb {
	webItem := PreserveDetailWeb{
		EventTime:        item.EventTime.Format(utils.DayTime),
		Channel:          item.Channel,
		Uv:               fmt.Sprintf("%d", item.Uv),
		NewUv:            fmt.Sprintf("%d", item.NewUv),
		PassedWeekActive: fmt.Sprintf("%d", item.PassedWeekActive),
		Day1Active:       fmt.Sprintf("%d", item.Day1Active),
		Day2Active:       fmt.Sprintf("%d", item.Day2Active),
		Day3Active:       fmt.Sprintf("%d", item.Day3Active),
		Day4Active:       fmt.Sprintf("%d", item.Day4Active),
		Day5Active:       fmt.Sprintf("%d", item.Day5Active),
		Day6Active:       fmt.Sprintf("%d", item.Day6Active),
		WeekActive:       fmt.Sprintf("%d", item.WeekActive),
		LastUpdate:       item.LastUpdate.Format(utils.FullTime),
	}
	return webItem
}

//Cols ...用户web显示使用
func (t PreserveDetail) Cols() []map[string]string {
	var cols []map[string]string
	colEventDay := map[string]string{
		"name": "日期",
		"key":  "event_time",
	}
	cols = append(cols, colEventDay)

	colClientChannel := map[string]string{
		"name":  "渠道",
		"key":   "channel",
		"click": "1",
	}
	cols = append(cols, colClientChannel)

	colPassedWeek := map[string]string{
		"name": "周活跃",
		"key":  "passed_week_active",
	}
	cols = append(cols, colPassedWeek)

	colUv := map[string]string{
		"name": "日活",
		"key":  "uv",
	}
	cols = append(cols, colUv)

	colNewUv := map[string]string{
		"name": "日新增",
		"key":  "new_uv",
	}
	cols = append(cols, colNewUv)
	//次日留存
	colDay1Active := map[string]string{
		"name": "1日留存",
		"key":  "day1_active",
	}
	cols = append(cols, colDay1Active)

	//二日留存
	colDay2Active := map[string]string{
		"name": "2日留存",
		"key":  "day2_active",
	}
	cols = append(cols, colDay2Active)

	//三日留存
	colDay3Active := map[string]string{
		"name": "3日留存",
		"key":  "day3_active",
	}
	cols = append(cols, colDay3Active)

	//四日留存
	colDay4Active := map[string]string{
		"name": "4日留存",
		"key":  "day4_active",
	}
	cols = append(cols, colDay4Active)

	//五日留存
	colDay5Active := map[string]string{
		"name": "5日留存",
		"key":  "day5_active",
	}
	cols = append(cols, colDay5Active)

	//六日留存
	colDay6Active := map[string]string{
		"name": "6日留存",
		"key":  "day6_active",
	}
	cols = append(cols, colDay6Active)

	//周留存
	colWeekActive := map[string]string{
		"name": "周留存",
		"key":  "week_active",
	}
	cols = append(cols, colWeekActive)

	colLastUpdate := map[string]string{
		"name": "更新时间",
		"key":  "last_update",
	}
	cols = append(cols, colLastUpdate)
	return cols
}

//GetAllChannels ...获取所有渠道
func (t PreserveDetail) GetAllChannels(client *xorm.Engine) ([]string, error) {
	var items []*PreserveDetail
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

func (t PreserveDetail) GetItemsByPage(client *xorm.Engine, chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]*PreserveDetail, int64, error) {
	var items []*PreserveDetail
	item := &PreserveDetail{}
	timeTS, timeTE := utils.ConvertToTime(tsStart), utils.ConvertToTime(tsEnd)
	session := client.Where("event_time>=?", timeTS).And("event_time<=?", timeTE)
	if chn != "" {
		chnList := utils.ChannelList(chn)
		session = session.In("channel", chnList)
	}

	session = session.Desc("event_time")
	if pageCount > 0 {
		session = session.Limit(pageCount, pageCount*(pageID-1))
	}
	err := session.Find(&items)
	if err != nil {
		glog.Errorf("[mysql]Get the items for from table %s failed,err:%+v", t.TableName(), err)
		return nil, 0, err
	}
	session = client.Where("event_time>=?", timeTS).And("event_time<=?", timeTE)
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

func (t PreserveDetail) GetChartItems(client *xorm.Engine, chn string, tsStart, tsEnd int64) (*utils.ChartDetail, error) {
	chartXvalue := []string{}
	chartXDic := map[string]bool{}
	timeTS, timeTE := utils.ConvertToTime(tsStart), utils.ConvertToTime(tsEnd)
	var items []*PreserveDetail
	session := client.Where("event_time>=?", timeTS).And("event_time<=?", timeTE)
	if chn != "" {
		chnList := utils.ChannelList(chn)
		session = session.In("channel", chnList)
	}
	err := session.OrderBy("event_time,channel").
		Find(&items)
	if err != nil {
		glog.Errorf("[mysql]Get the items for from table %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}

	chartUVValue := map[string]utils.ChartLineSeries{}
	chartNewUVValue := map[string]utils.ChartLineSeries{}
	for _, v := range items {
		//时间正序计算x轴的日期
		xValue := v.EventTime.Format(utils.DayTime)
		_, ok := chartXDic[xValue]
		if !ok {
			chartXDic[xValue] = true
			chartXvalue = append(chartXvalue, xValue)
		}
		//计算UV chart数据
		idx := fmt.Sprintf("%s%s%s", v.Channel, utils.SepChar, "-")
		val, ok := chartUVValue[idx]
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
		//计算新增UV chart
		val, ok = chartNewUVValue[idx]
		if ok {
			val.Data = append(val.Data, float64(v.NewUv))
			val.EventTime = append(val.EventTime, xValue)
			chartNewUVValue[idx] = val
		} else {
			chartNewUVValue[idx] = utils.ChartLineSeries{
				Data:      []float64{float64(v.NewUv)},
				EventTime: []string{xValue},
			}
		}
	}

	var chartYlines []utils.ChartSeriesYValue
	for k, v := range chartUVValue {
		infos := strings.Split(k, utils.SepChar)
		lineTitle := fmt.Sprintf("%s渠道UV趋势图", infos[0])
		chartYLine := utils.ChartSeriesYValue{
			Name:      lineTitle,
			ChartType: "line",
			Data:      v.Data,
			EventTime: v.EventTime,
		}
		chartYlines = append(chartYlines, chartYLine)
	}

	for k, v := range chartNewUVValue {
		infos := strings.Split(k, utils.SepChar)
		//chan_
		lineTitle := fmt.Sprintf("%s渠道新增UV趋势图", infos[0])
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

//GetItemsHistory
func (t PreserveDetail) GetItemsForHistory(client *xorm.Engine, chn string, tsStart int64, days int) ([]*PreserveDetail, error) {
	var items []*PreserveDetail
	timeTS := utils.ConvertToTime(tsStart)
	session := client.Where("event_time>=?", timeTS).And("channel =?", chn)

	session = session.OrderBy("event_time")
	if days >= 0 {
		session = session.Limit(days, 0)
	}
	err := session.Find(&items)
	if err != nil {
		glog.Errorf("[mysql]Get the items for from table %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}
