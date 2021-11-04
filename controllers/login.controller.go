package controllers

import (
	"net/http"
	"fmt"
	"time"

	"hotels/db"
	"hotels/models"

	"hotels/helpers"

	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
	
)

type LoginToken struct {
	Id string `json:"id"`
	Name string `json:"names"`
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token"`
	
}

var json map[string]interface{} = map[string]interface{}{}

func CheckLogin(c echo.Context) error {
	// username := c.FormValue("username")
	// password := c.FormValue("password")

	// fmt.Println(username)
	// fmt.Println(password)
	if err := c.Bind(&json); err != nil {
		return c.JSON(http.StatusOK, err)
	}

	var res models.Response
	db := db.DbManager()
	userss := []models.Users{}
	
	// Cek email apakah ada atau tidak
	if err := db.Where("email = ?", json["username"]).First(&userss); err.RowsAffected == 0 {
		res.Status = 202
		res.Message = "Email Tidak Ditemukan"
		res.Data = err.RowsAffected
		return c.JSON(http.StatusOK,res )
	}
	//LoginToken := &userss[0]
	getPassword := LoginToken{
		Id: userss[0].Id,
		Name: userss[0].Name,
		Email: userss[0].Email,
		Password: userss[0].Password,
		Token: "",
	}

	fmt.Println(getPassword.Password)	

	//cocokan email
	match, err := helpers.CheckPasswordHash(json["password"].(string), getPassword.Password)
	if !match {
		res.Status = 201
		res.Message = "Password salah"
		res.Data = err
		return c.JSON(http.StatusOK, res)
		
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = json["username"]
	claims["level"] = "application"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {

		res.Status = http.StatusOK
		res.Message = "Error"
		res.Data = err.Error()
		return c.JSON(http.StatusOK, res)
	 }
	 getPassword.Token = t
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = getPassword
	return c.JSON(http.StatusOK, res)
	
}
func CheckLogin2(c echo.Context) error {
	var res models.Response

	if err := c.Bind(&json); err != nil {
		return err
	}

	//username := c.FormValue("username")
	//password := c.FormValue("password")
	//bayu := c.QueryParam("username")
	//fmt.Println(bayu)
	//fmt.Println(password)
	fmt.Println(json["password"])

//	getPassword := json[0]

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = json
	return c.JSON(http.StatusOK, res)
}


func GenerateHashPassword(c echo.Context) error {
	//var res models.Response
	password := c.Param("password")

	hash, _ := helpers.HashPassword(password)

	return c.JSON(http.StatusOK, hash)
}
