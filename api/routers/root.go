package routers

import (
	"net/http"
	"squall/config"
	"squall/consts"
	"squall/generated"

	"github.com/gin-gonic/gin"
)

func (bh *BaseHandler) GetRoot(c *gin.Context) {
	status := generated.Root{
		string(config.Config.AppEnv),
		"live",
		consts.AppVersion,
	}
	c.PureJSON(http.StatusOK, status)
}
