package routers

import (
	"net/http"
	"squall/config"
	"squall/consts"

	"github.com/gin-gonic/gin"
)

type statusResp struct {
	Status  string            `json:"status"`
	Env     config.AppEnvType `json:"env"`
	Version string            `json:"version"`
}

func Roots(c *gin.Context) {
	status := statusResp{
		"live",
		config.Config.AppEnv,
		consts.AppVersion,
	}
	c.PureJSON(http.StatusOK, status)
}
