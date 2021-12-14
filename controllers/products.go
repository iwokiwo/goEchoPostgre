package controllers

import (
	"net/http"
	
	"hotels/db"
	"hotels/models"

	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
	
)

func GetProducts(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	idUser := claims["id"].(string)
	//fmt.Println(claims)
	//fmt.Println(idUser)

	db := db.DbManager()
	transaksis := []models.Products{}
	//db.Preload("Kategoris").Preload("Stores").Find(&transaksis)
	db.Preload("Kategoris").Preload("Stores").Find(&transaksis, "products.Owner_id = ?",idUser)

	
	return c.JSON(http.StatusOK, transaksis)
}
