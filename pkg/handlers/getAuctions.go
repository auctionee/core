package handlers

import (
	"net/http"

	"github.com/auctionee/core/internal/core/sync"

	"github.com/gin-gonic/gin"
)

func ListCurrentAuctions(c *gin.Context) {
	list := sync.AuctionsList{List: sync.LocalDB}
	c.JSON(http.StatusOK, list)
}
