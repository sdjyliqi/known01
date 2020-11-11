package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"known01/router"
)

func main() {
	flag.Parse()
	glog.Flush()
	r := gin.Default()
	// register the `/metrics` route.
	router.InitRouter(r)
	r.Run("0.0.0.0:8080")
}
