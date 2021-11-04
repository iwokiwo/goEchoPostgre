package routes

import (
	"hotels/controllers"
	//"net/http"
	"hotels/middleware"
	"github.com/labstack/echo"
	
)


func user(e *echo.Group) {
	grA := e.Group("/user")
	grA.GET("/", controllers.GetUsers,middleware.IsAuthenticated)
}

func registerUser(e *echo.Group){
	grA := e.Group("/user")
	grA.POST("/register", controllers.RegisterUser)
}

