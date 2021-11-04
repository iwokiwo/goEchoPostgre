package controllers

import (
	"net/http"

	"hotels/db"
	"hotels/models"

	"hotels/helpers"

	"github.com/labstack/echo"
)

var jsonRegister map[string]interface{} = map[string]interface{}{}

func GetUsers(c echo.Context) error {
	var res models.Response
	db := db.DbManager()
	userss := []models.Users{}
	db.Find(&userss)

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = userss
	return c.JSON(http.StatusOK, res)
}

func RegisterUser(c echo.Context) error {
	var res models.Response
	db := db.DbManager()
	if err := c.Bind(&jsonRegister); err != nil {
	
		return c.JSON(http.StatusOK, err)
	}

	// Cek email apakah ada atau tidak
	userss := []models.Users{}
	if err := db.Where("email = ?", jsonRegister["email"]).First(&userss); err.RowsAffected == 1 {
		res.Status = 203
		res.Message = "Email Sudah Terdaftar"
		res.Data = jsonRegister["email"]
		return c.JSON(http.StatusOK,res )
	}
	
	hash, _ := helpers.HashPassword(jsonRegister["password"].(string))
	var register = models.Users{
		Name: jsonRegister["name"].(string),
		Email: jsonRegister["email"].(string),
		Password: hash,
		Phone : jsonRegister["phone"].(string),
	}
	//jika tidak ada simpan data
	result := db.Create(&register) 
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = result.RowsAffected  
	return c.JSON(http.StatusOK, res)	
			
}
