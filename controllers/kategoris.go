package controllers

import (
	"net/http"

	"hotels/db"
	"hotels/models"

	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
)

func GetKategoris(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	idUser := claims["id"].(string)

	db := db.DbManager()
	kategori := []models.Kategoris{}
	db.Find(&kategori, "user_id = ?",idUser)
	//db.Joins("Kategoris").Joins("Stores").Find(&transaksis)

	// spew.Dump(json.Marshal(users))
	// return c.JSON(http.StatusOK, users)
	return c.JSON(http.StatusOK, kategori)
}

func AddKategoris(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	idUser := claims["id"].(string)

	var res models.Response
	db := db.DbManager()
	if err := c.Bind(&jsonRegister); err != nil {
	
		return c.JSON(http.StatusOK, err)
	}

	var addkategori = models.Kategoris{
		User_id: idUser,
		Nama_kategori: jsonRegister["nama_kategori"].(string),

	}
	//jika tidak ada simpan data
	result := db.Create(&addkategori) 
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = result.RowsAffected 
	return c.JSON(http.StatusOK, res)

}
