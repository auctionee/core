package server

import (
	"net/http"

	"github.com/auctionee/core/internal"
	"github.com/auctionee/core/internal/perm"
	"github.com/auctionee/core/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func NewServer(cfg *internal.Config) http.Server {
	return http.Server{
		Addr:    cfg.Port,
		Handler: setupRouter(),
	}
}
func setupRouter() *gin.Engine {
	r := gin.New()
	engine := r.Group("/engine/")
	getter := engine.Group("/get/")
	r.Use(gin.Logger(), gin.Recovery())
	engine.Use(func(c *gin.Context) {
		perm.Permissions(c)
	})
	engine.POST("start/", func(c *gin.Context) {
		handlers.StartAuction(c)
	})
	getter.GET("auctions/", func(c *gin.Context) {})
	getter.GET("goods/mine/", func(c *gin.Context) {})
	getter.GET("goods/all/", func(c *gin.Context) {})

	return r
}
