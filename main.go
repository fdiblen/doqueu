package main

import (
	"errors"
	"net/http"

	"github.com/fdiblen/doqueu/controller"
	_ "github.com/fdiblen/doqueu/docs"
	"github.com/fdiblen/doqueu/httputil"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Docker Queue API
// @version         1.0
// layout 		    "BaseLayout"
// @description     This is a server to manage Docker containers.
// @termsOfService  http://doqueu.io/terms/

// @contact.name   API Support
// @contact.url    http://www.doqueu.io/support
// @contact.email  support@doqueu.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
// @description					doqueu API Authorization

// @securitydefinitions.oauth2.application  OAuth2Application
// @tokenUrl                                https://doqueu.io/oauth/token
// @scope.write                             Grants write access
// @scope.admin                             Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit  OAuth2Implicit
// @authorizationUrl                     https://doqueu.io/oauth/authorize
// @scope.write                          Grants write access
// @scope.admin                          Grants read and write access to administrative information

// @securitydefinitions.oauth2.password  OAuth2Password
// @tokenUrl                             https://doqueu.io/oauth/token
// @scope.read                           Grants read access
// @scope.write                          Grants write access
// @scope.admin                          Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode  OAuth2AccessCode
// @tokenUrl                               https://doqueu.io/oauth/token
// @authorizationUrl                       https://doqueu.io/oauth/authorize
// @scope.admin                            Grants read and write access to administrative information

func main() {
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	c := controller.NewController()

	v1 := server.Group("/api/v1")
	{
		accounts := v1.Group("/accounts")
		{
			accounts.GET(":id", c.ShowAccount)
			accounts.GET("", c.ListAccounts)
			accounts.POST("", c.AddAccount)
			accounts.DELETE(":id", c.DeleteAccount)
			accounts.PATCH(":id", c.UpdateAccount)
			accounts.POST(":id/images", c.UploadAccountImage)
		}
		admin := v1.Group("/admin")
		{
			admin.Use(auth())
			admin.POST("/auth", c.Auth)
		}
		containers := v1.Group("/containers")
		{
			containers.GET("", c.ListContainers)
			containers.GET(":id", c.ShowContainer)
			containers.POST("/run", c.RunContainer)
			containers.POST("/stop", c.StopContainer)
		}
	}
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.Run(":8080")
}

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.GetHeader("Authorization")) == 0 {
			httputil.NewError(c, http.StatusUnauthorized, errors.New("Authorization is required Header"))
			c.Abort()
		}
		c.Next()
	}
}
