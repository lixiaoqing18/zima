package kernel

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lixiaoqing18/zima/framework"
)

type ZimaGinService struct {
	container framework.Container
	engine    *gin.Engine
}

func NewZimaGinService(params ...any) (any, error) {
	c := params[0].(framework.Container)
	e := params[1].(*gin.Engine)
	return &ZimaGinService{
		container: c,
		engine:    e,
	}, nil
}

func (service *ZimaGinService) WebEngine() http.Handler {
	return service.engine
}
