package main

import (
	"fmt"
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"known01/conf"
	"known01/router"
)

func main() {
	defer glog.Flush()
	r := gin.Default()
	// register the `/metrics` route.
	router.InitRouter(r)
	ginpprof.Wrapper(r)
	r.Run(fmt.Sprintf("0.0.0.0:%d", conf.DefaultConfig.Port))
}
