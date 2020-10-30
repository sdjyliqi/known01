package control

import (
	"github.com/go-xorm/xorm"
	"github.com/sdjyliqi/feirars/conf"
	"github.com/sdjyliqi/feirars/models"
	"github.com/sdjyliqi/feirars/utils"
)

type UserCenter interface {
	Login(userID, passport string) error //用户登录
	Logout(userID string) error          //用户登录
	UserAuthChn(userID, requestChn string) string
	UserChn(userID string) ([]string, error)
}

type userCenter struct {
	db        *xorm.Engine
	cfg       *conf.FeirarConfig
	UserBasic models.UserBasic
}

func CreateUserCenter(cfg *conf.FeirarConfig) UserCenter {
	utils.InitMySQL(cfg.DBMysql, false)
	return &userCenter{
		db:        utils.GetMysqlClient(),
		cfg:       cfg,
		UserBasic: models.UserBasic{},
	}
}
