package route

import (
	loginModel "builder/restful-api-gogin/application/controllers/auth/login"
	registerModel "builder/restful-api-gogin/application/controllers/auth/register"
	handlerLogin "builder/restful-api-gogin/application/handlers/auth/login"
	handlerRegister "builder/restful-api-gogin/application/handlers/auth/register"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler Auth
	*/
	LoginRepository := loginModel.NewRepositoryLogin(db)
	loginService := loginModel.NewServiceLogin(LoginRepository)
	loginHandler := handlerLogin.NewHandlerLogin(loginService)

	registerRepository := registerModel.NewRepositoryRegister(db)
	registerService := registerModel.NewServiceRegister(registerRepository)
	registerHandler := handlerRegister.NewHandlerRegister(registerService)

	/**
	@description All Auth Route
	*/
	groupRoute := route.Group("/api/v1/login")
	groupRoute.POST("/register", registerHandler.RegisterHandler)
	groupRoute.POST("/login", loginHandler.LoginHandler)

}
