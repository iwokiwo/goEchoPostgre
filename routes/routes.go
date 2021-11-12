package routes

import (
	"hotels/controllers"
	"net/http"
	//"hotels/middleware"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	// }))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"authorization", "Content-Type"},
		AllowCredentials: true,
		AllowMethods:     []string{echo.OPTIONS, echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Helfffoo")
	})
	v1 := e.Group("/api/v1")
	produk(v1)

	user(v1)
	registerUser(v1)

	kategori(v1)
	transaksidtls(v1)
	GenerateHashPassword(v1)
	login(v1)

	return e
}
func produk(e *echo.Group) {
	grA := e.Group("/produk")
	grA.GET("/", controllers.GetProducts,middleware.IsAuthenticated)
}

// func user(e *echo.Group) {
// 	grA := e.Group("/user")
// 	grA.GET("/", controllers.GetUsers,middleware.IsAuthenticated)
// }

func GenerateHashPassword(e *echo.Group) {
	grA := e.Group("/login")
	grA.GET("/buatPass/:password", controllers.GenerateHashPassword)
}

func login(e *echo.Group) {
	grA := e.Group("/login")
	grA.POST("/", controllers.CheckLogin)
}


func kategori(e *echo.Group) {
	grA := e.Group("/kategori")
	grA.GET("/", controllers.GetKategoris)
}
func transaksidtls(e *echo.Group) {
	grA := e.Group("/transaksi")
	grA.GET("/transaksidtls", controllers.Gettrsdtls)
}
