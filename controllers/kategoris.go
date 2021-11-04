package controllers

import (
	"net/http"

	"hotels/db"
	"hotels/models"

	"github.com/labstack/echo"
)

func GetKategoris(c echo.Context) error {
	db := db.DbManager()
	kategori := []models.Kategoris{}
	db.First(&kategori)
	//db.Joins("Kategoris").Joins("Stores").Find(&transaksis)

	// spew.Dump(json.Marshal(users))
	// return c.JSON(http.StatusOK, users)
	return c.JSON(http.StatusOK, kategori)
}
