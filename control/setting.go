package control

import (
	"github.com/go-xorm/xorm"
	"github.com/sdjyliqi/feirars/conf"
	"github.com/sdjyliqi/feirars/models"
	"github.com/sdjyliqi/feirars/utils"
)

type SettingCenter interface {
	ClientUpdate(request *utils.UpdateArgs) (string, error)
}

type settingCenter struct {
	db          *xorm.Engine
	cfg         *conf.FeirarConfig
	newsSetting models.NewsSetting
}

func CreateSettingCenter(cfg *conf.FeirarConfig) SettingCenter {
	utils.InitMySQL(cfg.DBMysql, false)
	return &settingCenter{
		db:          utils.GetMysqlClient(),
		cfg:         cfg,
		newsSetting: models.NewsSetting{},
	}

}
