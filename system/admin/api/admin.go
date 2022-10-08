package main

import (
	"flag"
	"fmt"
	"net/http"

	"erik-agile/common/errorx"
	"erik-agile/system/admin/api/internal/config"
	"erik-agile/system/admin/api/internal/handler"
	"erik-agile/system/admin/api/internal/svc"
	"erik-agile/system/admin/api/internal/svc/gorm"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/admin-api.yaml", "the config file")

func main() {
    flag.Parse()

    var c config.Config
    conf.MustLoad(*configFile, &c)
    logx.MustSetup(c.LogConf)
    server := rest.MustNewServer(c.RestConf)
    defer server.Stop()

    ctx := svc.NewServiceContext(c)
    db := gorm.NewGormdb(c)
    handler.RegisterHandlers(server, ctx, db)
    // 自定义错误
    httpx.SetErrorHandler(func(err error) (int, interface{}) {
        switch e := err.(type) {
        case *errorx.CodeError:
            return http.StatusOK, e.Datas()
        default:
            return http.StatusInternalServerError, nil
        }
    })
    fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
    server.Start()
}
