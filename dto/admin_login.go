package dto

import (
	"github.com/e421083458/go_gateway_demo/public"
	"github.com/gin-gonic/gin"
)

type AdminLoginInput struct {
	Username string `json:"username" form:"username" comment:"姓名" example:"姓名" validate:"required"`
	Password string `json:"password" form:"password" comment:"密码" example:"123456" validate:"required"`
}

func (params *AdminLoginInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}
