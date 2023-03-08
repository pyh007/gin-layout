package routers

import (
	"github.com/gin-gonic/gin"
	controllerV1 "github.com/wannanbigpig/gin-layout/internal/controller/v1"
	controllerV1test "github.com/wannanbigpig/gin-layout/internal/controller/v1/mysql"
)

//111
func setApiRoute(r *gin.Engine) {
	// version 1
	v1 := r.Group("/api/v1")
	{
		auth := controllerV1.NewAuthController()
		v1.POST("/login", auth.Login)
		v1.POST("/register", auth.Register)
		demo := controllerV1.NewDemoController()
		v1.GET("/hello-world", demo.HelloWorld)

		mysqlTest := controllerV1test.NewMysqlTestController()
		v1.POST("/trans", mysqlTest.Transaction)
	}
}
