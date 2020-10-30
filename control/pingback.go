package control

import (
	"errors"
	"github.com/go-xorm/xorm"
	"github.com/sdjyliqi/feirars/conf"
	"github.com/sdjyliqi/feirars/models"
	"github.com/sdjyliqi/feirars/utils"
	"strings"
)

type PingbackCenter interface {
	GetActiveDetailItems(chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]models.ActiveDetailWeb, int64, error) //获取客户的激活方式统计数据
	GetActiveDetailCols() []map[string]string
	GetActiveChart(chn string, tsStart, tsEnd int64) (*utils.ChartDetail, error)
	GetActiveChannel(name string) ([]string, error)

	GetFeirarDetailItems(chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]models.FeirarDetailWeb, int64, error) //获取feirar请求统计数据
	GetFeirarDetailCols() []map[string]string
	GetFeirarChart(chn string, tsStart, tsEnd int64) (*utils.ChartDetail, error)
	GetFeirarChannel(name string) ([]string, error)

	//统计feirar news接口相关数据
	GetFeirarNewsDetailItems(chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]models.FeirarDetailWeb, int64, error) //获取feirar请求统计数据
	GetFeirarNewsDetailCols() []map[string]string
	GetFeirarNewsChannel(name string) ([]string, error)
	GetFeirarNewsChart(chn string, tsStart, tsEnd int64) (*utils.ChartDetail, error)

	//统计feirar update接口相关数据
	GetFeirarUpdateDetailItems(chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]models.FeirarDetailWeb, int64, error) //获取feirar请求统计数据
	GetFeirarUpdateDetailCols() []map[string]string
	GetFeirarUpdateChannel(name string) ([]string, error)
	GetFeirarUpdateChart(chn string, tsStart, tsEnd int64) (*utils.ChartDetail, error)

	GetInstallDetailItems(chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]models.InstallDetailWeb, int64, error) //获取安装统计数据
	GetInstallDetailCols() []map[string]string
	GetInstallChannel(name string) ([]string, error)
	GetInstallChart(chn string, tsStart, tsEnd int64) (*utils.ChartDetail, error)
	GetInstallHistoryCalculator(chn string, tsStart int64, days int) ([]*utils.HistoryDetail, error)

	GetUninstallDetailItems(chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]models.UninstallDetailWeb, int64, error) //获取卸载统计数据
	GetUninstallDetailCols() []map[string]string
	GetUninstallChannel(name string) ([]string, error)
	GetUninstallChart(chn string, tsStart, tsEnd int64) (*utils.ChartDetail, error)

	//信息点击相关统计
	GetNewsDetailItems(chn string, pageID, pageCount int, tsStart, tsEnd int64, eventKey string) ([]models.NewsDetailWeb, int64, error) //获取客户端弹窗统计数据
	GetNewsDetailCols() []map[string]string
	GetNewsChannel(name, eventKey string) ([]string, error)
	GetNewsChart(chn string, tsStart, tsEnd int64, eventKey string) (*utils.ChartDetail, error)
	//留存相关接口
	GetPreserveDetailItems(chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]models.PreserveDetailWeb, int64, error) //获取留存统计数据
	GetPreserveDetailCols() []map[string]string
	GetPreserveChannel(name string) ([]string, error)
	GetPreserveChart(chn string, tsStart, tsEnd int64) (*utils.ChartDetail, error)
	GetPreserveHistoryCalculator(chn string, tsStart int64, days int) ([]*utils.HistoryDetail, error)

	//信息点击相关统计
	GetUdtrstDetailItems(chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]models.UdtrstDetailWeb, int64, error) //获取客户端弹窗统计数据
	GetUdtrstDetailCols() []map[string]string
	GetUdtrstChannel(name string) ([]string, error)
	GetUdtrstChart(chn string, tsStart, tsEnd int64) (*utils.ChartDetail, error)
}

type pingbackCenter struct {
	db              *xorm.Engine
	cfg             *conf.FeirarConfig
	activeDetail    models.ActiveDetail
	feirarDetail    models.FeirarDetail
	installDetail   models.InstallDetail
	uninstallDetail models.UninstallDetail
	newsDetail      models.NewsDetail
	preserveDetail  models.PreserveDetail
	userBasic       models.UserBasic
	udtrstDetail    models.UdtrstDetail
}

func CreatePingbackCenter(cfg *conf.FeirarConfig) PingbackCenter {
	utils.InitMySQL(cfg.DBMysql, true)
	return &pingbackCenter{
		db:              utils.GetMysqlClient(),
		cfg:             cfg,
		activeDetail:    models.ActiveDetail{},
		feirarDetail:    models.FeirarDetail{},
		installDetail:   models.InstallDetail{},
		uninstallDetail: models.UninstallDetail{},
		newsDetail:      models.NewsDetail{},
		preserveDetail:  models.PreserveDetail{},
		userBasic:       models.UserBasic{},
		udtrstDetail:    models.UdtrstDetail{},
	}
}

func (uc *pingbackCenter) UserChn(userID string) (string, error) {
	item, err := uc.userBasic.GetUserBasic(uc.db, userID)
	if err != nil {
		return "", err
	}
	if item == nil {
		return "", errors.New("must-auth")
	}
	return item.Chn, nil
}

func (uc *pingbackCenter) UserAuthChn(userID, requestChn string) string {
	item, err := uc.userBasic.GetUserBasic(uc.db, userID)
	if err != nil {
		return "XXX"
	}
	//如果该用户的授权渠道为空，证明全部放开
	if item.Chn == "" {
		return requestChn
	}
	//如果用户请求的渠道为空，直接返回已经授权的渠道
	if requestChn == "" {
		return item.Chn
	}

	chnMap := map[string]bool{}
	chnListInDB := strings.Split(item.Chn, ",")
	for _, v := range chnListInDB {
		chnMap[v] = true
	}
	authChnList := []string{}
	chnListRequest := strings.Split(requestChn, ",")
	for _, v := range chnListRequest {
		_, ok := chnMap[v]
		if ok {
			authChnList = append(authChnList, v)
		}
	}
	if len(authChnList) == 0 {
		return "XXX"
	}
	return strings.Join(authChnList, ",")
}
