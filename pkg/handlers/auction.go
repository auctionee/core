package handlers

import (
	"net/http"

	"github.com/auctionee/core/internal/core"

	"github.com/auctionee/core/pkg/models"
	"github.com/gin-gonic/gin"
)

func Process(c *gin.Context) {
	request := models.NewRequest()
	err := request.Unmarshall(c.Request.Body)
	if err != nil {
		c.AbortWithError(
			http.StatusBadRequest,
			err,
		)
	}
	if err = core.Execute(request); err != nil {
		c.AbortWithError(
			http.StatusBadRequest,
			err,
		)
	}
	c.Status(http.StatusOK)
}
