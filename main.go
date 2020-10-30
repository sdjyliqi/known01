package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/sdjyliqi/known01/router"
)

func main() {
	flag.Parse()
	glog.Flush()
	r := gin.Default()

	// register the `/metrics` route.
	router.InitRouter(r)
	r.Run(":8899")

}
