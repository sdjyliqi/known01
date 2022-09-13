/*
 author: lijingcan
 date: 2022/9/13 12:03
 desc: 接入prometheus监控，增加应用响应时间指标，标签添加在middleware/log中
*/

package handle

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
)

func PromHandler(handler http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}
var (
	//不同响应时间段请求量分布
	WebRequestSecondsBucket = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "Http_server_requests_seconds",
			Help: "Duration distribution of the same request",
			Buckets: []float64{0.1, 0.3, 0.6, 1.0, 3.0},
		},
		[]string{"application", "uri", "status"},
	)
)
