package routers

import (
	"squall/generated"
	"squall/models"
)

var getDB = models.GetDB // テスト時はこれをモックターゲットにする。DB呼び出しはすべてgetDBを使う

type BaseHandler struct{}

var _ generated.ServerInterface = &BaseHandler{}
