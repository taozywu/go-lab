package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"kratos-demo/internal/server/http"
	"kratos-demo/internal/service"

	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
)

func main() {
	// 命令行解析
	flag.Parse()
	// 加载../configs/所有toml
	if err := paladin.Init(); err != nil {
		panic(err)
	}
	// 初始化log
	log.Init(nil) // debug flag: log.dir={path}
	// 延迟关闭log
	defer log.Close()
	// 输出log info
	log.Info("kratos-demo start")
	// 创建一个新的service
	svc := service.New()
	// 将service带入http server中
	httpSrv := http.New(svc)
	// 初始一个channel，1个长度，信号类型
	c := make(chan os.Signal, 1)
	// 将如下几个信号输入到c中
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		// 读取c
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		// 退出、程序自己正常退出、程序终止
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
			// shutdown
			if err := httpSrv.Shutdown(ctx); err != nil {
				log.Error("httpSrv.Shutdown error(%v)", err)
			}
			log.Info("kratos-demo exit")
			// 业务逻辑清理
			svc.Close()
			// 如果在超时之前慢操作完成，则释放资源
			cancel()
			// 睡眠1s
			time.Sleep(time.Second)
			return
		// 挂起
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
