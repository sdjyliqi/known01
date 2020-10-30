package utils

//var redisOnce sync.Once
//var redisClients *redis.Client
//
////parseRedisAddr ...解析数据库地址结构 redis://db:password@host:port
//func parseRedisAddr(addr string) (host string, password string, db int) {
//	db = 0
//	u, err := url.Parse(addr)
//	if err != nil {
//		host = addr
//	} else {
//		host = u.Host
//		db64, _ := strconv.ParseInt(u.User.Username(), 0, 32)
//
//		db = int(db64)
//		password, _ = u.User.Password()
//	}
//	glog.V(4).Infof("parse redis URI. addr = %s, host = %s, db = %d", safeRedisAddr(addr), host, db)
//	return
//}
//
//func safeRedisAddr(addr string) string {
//	return regexp.MustCompile(`://([^:]+):(.*)@`).ReplaceAllString(addr, "://$1:****@")
//}
//
//func createOneRedisClient(addr string) {
//	host, password, db := parseRedisAddr(addr)
//	redisClients = redis.NewClient(&redis.Options{
//		Addr:     host,
//		Password: password,
//		DB:       db,
//	})
//}
//
////InitRedisClients 初始化一次redis连接，后续直接可以使用了
//func InitRedisClients(addr string) {
//	redisOnce.Do(func() {
//		glog.V(4).Infof("[init] Init redis client for the address %s.", addr)
//		createOneRedisClient(addr)
//	})
//}
//
////GetRedisClient  ... 返回redis.Client类型的客户端连接
//func GetRedisClient() *redis.Client {
//	return redisClients
//}
