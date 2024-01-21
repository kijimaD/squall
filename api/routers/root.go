package routers

import (
	"net/http"
	"squall/config"
	"squall/consts"
	"squall/generated"

	"github.com/gin-gonic/gin"
)

func Roots(c *gin.Context) {
	status := generated.GetRootResp{
		string(config.Config.AppEnv),
		"live",
		consts.AppVersion,
	}
	c.PureJSON(http.StatusOK, status)
}
