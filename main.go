package main

import (
	"fmt"
	"github.com/bowillkin/proto"
	"github.com/bowillkin/proto/ipip"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/koding/multiconfig"
	"log"
	"net/http"
)

type defaultConfig struct {
	ServerAddress string `default:"localhost:80"`
	IpAddress     string `default:"ipip-grpc:80"`
}

type IpReq struct {
	Ip string `json:"ip" binding:"required"`
}

func main() {
	r := gin.Default()
	config := new(defaultConfig)
	m := multiconfig.New()
	m.MustLoad(config)
	dev := false
	if dev {
		config.IpAddress = "localhost:5000"
		config.ServerAddress = "localhost:8081"
	}
	ipipConn, err := proto.DefaultConn(config.IpAddress)
	if err != nil {
		log.Fatalf("new account client: %s", err)
	}
	ipipClient := ipip.NewIpipClient(ipipConn)
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "")
	})
	r.POST("/ip", func(c *gin.Context) {
		var (
			ipReq IpReq
		)
		if err := c.MustBindWith(&ipReq, binding.JSON); err != nil {
			c.JSON(http.StatusBadRequest, "bad")
		} else {
			ipRes, err := ipipClient.GetAreaDataByIp(c.Request.Context(), &ipip.GetAreaDataByIpReq{
				RemoteIp: ipReq.Ip,
			})
			if err != nil {
				log.Println(err)
				c.JSON(http.StatusBadGateway, "bad")
			} else {
				c.JSON(http.StatusOK, fmt.Sprintf("ok, ip data: %v", ipRes))
			}
		}
	})

	if err := r.Run(config.ServerAddress); err != nil {
		log.Fatalf("server run: %s", err)
	}
}
