package mysql

import (
	"github.com/gin-gonic/gin"
	"github.com/wannanbigpig/gin-layout/internal/controller"
	"github.com/wannanbigpig/gin-layout/internal/service"
	"github.com/wannanbigpig/gin-layout/internal/validator"
	"github.com/wannanbigpig/gin-layout/internal/validator/form"
)

type TestMysqlController struct {
	controller.Api
}

func NewMysqlTestController() *TestMysqlController {
	return &TestMysqlController{}
}

// Transaction 测试gorm事务
func (api *TestMysqlController) Transaction(c *gin.Context) {

	registerForm := form.RegisterForm()
	if err := validator.CheckPostParams(c, &registerForm); err != nil {
		return
	}
	// 实际业务调用
	err := service.NewAuthService().Register(registerForm.UserName, registerForm.PassWord)
	// 根据业务返回值判断业务成功 OR 失败
	if err != nil {
		api.Err(c, err)
		return
	}

	api.Success(c)
}
