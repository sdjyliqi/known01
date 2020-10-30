package middleware

//import (
//	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
//	"github.com/oschwald/geoip2-golang"
//	"sync"
//)
//
//var Reader *geoip2.Reader
//
//var ip2util *ip2region.Ip2Region
//var loadIPLocOnce sync.Once

//
//func init() {
//	loadIPLocOnce.Do(func() {
//		var err error
//		ip2util, err = ip2region.New(conf.DefaultConfig.IPLOC)
//		if err != nil {
//			log.Fatalf("[init]Load GeoLite2-City.mmdb failed,err:%+v", err)
//		}
//	})
//}
//
////... RequestAddIPLoc  add request id into header
//func RequestAddIPLoc() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		ipInfo, err := ip2util.BtreeSearch(c.ClientIP())
//		if err == nil {
//			ipCity := ipInfo.City
//			ipProvince := ipInfo.Province
//			ipCityID := ipInfo.CityId
//			c.Request.Header.Add("IPLOC", fmt.Sprintf("%d", ipCityID))
//			c.Request.Header.Add("IPCITY", ipCity)
//			c.Request.Header.Add("IPPROVINCE", ipProvince)
//		}
//		c.Next()
//	}
//}
