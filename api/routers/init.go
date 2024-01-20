package routers

import "squall/models"

var getDB = models.GetDB // テスト時はこれをモックターゲットにする。DB呼び出しはすべてgetDBを使う
