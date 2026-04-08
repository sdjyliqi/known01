package main

import (
	"flag"
	"fmt"
	"github.com/cleey/glogrotate"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"known01/conf"
	"known01/handle"
	"known01/router"
	"known01/utils"
	"math/rand"
	"net/http"
	"time"
)

//logCut...完成日志切割,默认值为1小时切割一次，只需设置保留的时间即可。
func logCut() {
	glogrotate.Start(glogrotate.RotateOption{
		Remain: time.Hour * 24 * 7,
	})
}

func init() {
	var ymlPath string
	flag.StringVar(&ymlPath, "c", "", "configuration file")
	flag.Parse()
	if ymlPath == "" {
		glog.Fatal("You must input path of the yml ....")
	}
	//初始化配置，覆盖原来的默认配置参数苏
	conf.InitConfig(ymlPath, &conf.DefaultConfig)
	//检查配置项的合法性，如果任何一项为空，立即fatal掉
	if conf.DefaultConfig.DBMysql == "" || conf.DefaultConfig.Port == 0 || conf.DefaultConfig.WordDic == "" {
		fmt.Println("The content of yml is invalid.")
		glog.Errorln("The content of yml is invalid.")
		glog.Fatal("The content of yml is invalid.")
	}
	rand.Seed(time.Now().UnixNano())
	utils.InitMySQL(conf.DefaultConfig.DBMysql, true) //建立MySQL连接
	utils.InitSegDic()                                //初始化分词词表
	handle.InitBrain()                                //初始化
	logCut()
	// 初始化监控指标
	prometheus.MustRegister(
		handle.WebRequestSecondsBucket,
	)
}

func main() {
	// 1. 在独立的 Goroutine 中启动 Prometheus metrics 服务
	go func() {
		promMux := http.NewServeMux()
		promMux.Handle("/metrics", promhttp.Handler())
		promAddr := fmt.Sprintf("0.0.0.0:%d", conf.DefaultConfig.PromPort)

		fmt.Printf("Prometheus metrics server is running on %s\n", promAddr)
		glog.Infof("Prometheus metrics server is running on %s", promAddr)

		if err := http.ListenAndServe(promAddr, promMux); err != nil {
			glog.Fatalf("Failed to start Prometheus metrics server: %v", err)
		}
	}()

	// 2. 启动主应用服务
	r := gin.Default()
	// gin.SetMode(gin.ReleaseMode)

	router.InitRouter(r)

	appAddr := fmt.Sprintf("0.0.0.0:%d", conf.DefaultConfig.Port)
	fmt.Printf("Application server is running on %s\n", appAddr)

	r.Run(appAddr)
}
