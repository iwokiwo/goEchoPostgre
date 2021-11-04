package controllers

import (
	"net/http"

	"hotels/db"
	"hotels/models"

	"github.com/labstack/echo"
)

func GetProducts(c echo.Context) error {
	db := db.DbManager()
	transaksis := []models.Products{}
	db.Preload("Kategoris").Preload("Stores").Find(&transaksis)
	//db.Joins("Kategoris").Joins("Stores").Find(&transaksis)

	// spew.Dump(json.Marshal(users))
	// return c.JSON(http.StatusOK, users)
	return c.JSON(http.StatusOK, transaksis)
}
