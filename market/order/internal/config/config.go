package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

// 这里的配置映射的是etc里面的文件
type Config struct {
	rest.RestConf
	// 去ETCD获取user的地址
	UserRpc zrpc.RpcClientConf
}
