package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	// 去etcd获取user rpc的地址
	UserRpc zrpc.RpcClientConf
}
