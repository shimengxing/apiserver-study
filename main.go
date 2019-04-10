package main

import (
	"apiserver-study/config"
	"errors"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
	"time"

	"apiserver-study/router"
	"github.com/gin-gonic/gin"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver配置文件路径")
)

func main() {
	pflag.Parse()

	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine.
	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	router.Load(
		g,
		middlewares...,
	)

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("路由器未响应，或者启动超时", err)
		}
		log.Info("路由器部署成功")
	}()

	log.Infof("开始监听端口: %s", viper.GetString("addr"))
	log.Infof(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("url") + viper.GetString("addr") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Info("等待路由器，1秒后重试")
		time.Sleep(time.Second)
	}
	return errors.New("不能连接路由器")
}
