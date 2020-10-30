package handle

import (
	"github.com/sdjyliqi/feirars/conf"
	"github.com/sdjyliqi/feirars/control"
	"sync"
)

var onceControl sync.Once
var PingbackCenter control.PingbackCenter
var UCenter control.UserCenter
var SettingCenter control.SettingCenter //负责升级或者弹窗下发

func init() {
	onceControl.Do(func() {
		PingbackCenter = control.CreatePingbackCenter(&conf.DefaultConfig)
		UCenter = control.CreateUserCenter(&conf.DefaultConfig)
		SettingCenter = control.CreateSettingCenter(&conf.DefaultConfig)
	})
}
