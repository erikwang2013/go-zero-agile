package config

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
    rest.RestConf
    Mysql struct {
        DataSource  string
        TablePrefix string
    }

    CacheRedis cache.CacheConf
    Auth       struct {
        AccessSecret string
        AccessExpire int64
    }
    Permission struct {
        Role string
    }
    LogConf  logx.LogConf
    AdminRpc zrpc.RpcClientConf
}
