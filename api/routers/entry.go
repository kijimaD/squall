package routers

import (
	"net/http"
	"squall/helper"
	"squall/models"

	"github.com/gin-gonic/gin"
)

func (bh *BaseHandler) GetEntries(c *gin.Context) {
	var es []models.Entry
	err := getDB().Find(&es).Error
	if err != nil {
		helper.ErrorResponse(c, err)

		return
	}
	c.PureJSON(http.StatusOK, &es)
}
