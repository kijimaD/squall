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
	// Where(&models.Entry{IsDone: false}) で検索すると、falseはゼロ値なので検索から無視される...
	// なので文字列で条件を指定する
	err := getDB().Debug().Limit(size).Where("is_done = ?", false).Not(ids).Find(&es).Error
	if err != nil {
		helper.ErrorResponse(c, err)

		return
	}
	c.PureJSON(http.StatusOK, &es)
}

func (bh *BaseHandler) PostDoneEntry(c *gin.Context, entryId generated.EntryIdParam) {
	entry := models.Entry{ID: helper.GetPtr(uint(entryId))}
	err := getDB().First(&entry).Error
	if err != nil {
		helper.ErrorResponse(c, err)

		return
	}

	err = getDB().Model(&entry).Updates(models.Entry{
		IsDone: true,
	}).Error
	if err != nil {
		helper.ErrorResponse(c, err)

		return
	}
	c.PureJSON(http.StatusOK, &entry)
}
