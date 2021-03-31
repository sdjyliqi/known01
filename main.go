package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"known01/conf"
	"known01/router"
	"math/rand"
	"time"
)

func init() {
	var confFile string
	flag.StringVar(&confFile, "c", "", "configuration file")
	flag.Parse()
	if confFile == "" {
		glog.Fatal("You must input path of the yml ....")
	}
	conf.Init(confFile, &conf.DefaultConfig)
	rand.Seed(time.Now().UnixNano())
}

func main() {
	defer glog.Flush()
	r := gin.Default()
	// register the `/metrics` route.
	router.InitRouter(r)
	r.Run(fmt.Sprintf("0.0.0.0:%d", conf.DefaultConfig.Port))
}
