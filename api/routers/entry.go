package routers

import (
	"net/http"
	"squall/generated"
	"squall/helper"
	"squall/models"

	"github.com/gin-gonic/gin"
)

func (bh *BaseHandler) GetEntries(c *gin.Context, params generated.GetEntriesParams) {
	size := 20
	if params.Size != nil {
		size = *params.Size
	}

	ids := []int{}
	if params.IgnoreIds != nil {
		ids = *params.IgnoreIds
	}
	var es []models.Entry
	err := getDB().Limit(size).Not(ids).Find(&es).Error
	if err != nil {
		helper.ErrorResponse(c, err)

		return
	}
	c.PureJSON(http.StatusOK, &es)
}
