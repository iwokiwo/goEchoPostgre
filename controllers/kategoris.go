package controllers

import (
	"encoding/json"
	"net/http"
	//"strconv"
	_ "encoding/json"
	"fmt"
	"hotels/db"
	"hotels/models"


	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type Multi struct {
	User_id           int    `json:"user_ids"`
	Nama_kategori string `json:"nm_kategori"`
}
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
	Bayu []string `json:"bayu"`
}



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

func DeleteKategoris(c echo.Context) error {
	var res models.Response
	db := db.DbManager()
	
	if err := c.Bind(&jsonRegister); err != nil {
	
		return c.JSON(http.StatusOK, err)
	}
	var yes uint=uint(jsonRegister["id"].(float64))
	var deleteKategori = models.Kategoris{
		ID :yes,
	}
	result := 	db.Delete(&deleteKategori)
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = result.RowsAffected 
	return c.JSON(http.StatusOK, res)


}

func AddKategoriMulti(c echo.Context) error {


	var res models.Response
	db := db.DbManager()

	m := make(map[string]interface{})
	//var obj map[string]interface{}
	var arr []interface{}
	if err := c.Bind(&m); err != nil {
		return err
	}
	strs := m["bayu"].([]interface{})
	//arr = append(arr, strs)

	//tambahkan objek user_id di dalam json bayu
	//obj["user_id"] = idUser
	//strs = append(strs, obj)
	jsonString, _ := json.Marshal(&m)
	jsonStringBayu, _ := json.Marshal(strs)

	fmt.Println(m)
	fmt.Println(strs)
	fmt.Println(string(jsonString))
	fmt.Println(string(jsonStringBayu))

	var data []Multi

	json.Unmarshal(jsonStringBayu, &data)
	fmt.Println(data)


	var modeldata []models.Kategoris
	json.Unmarshal(jsonStringBayu,&modeldata)
	fmt.Println(modeldata)

	//snafing
	json.Unmarshal([]byte("{}"), &m)
	m["id"]="1"
	arr = append(arr, m)
	fmt.Println("snefing objek",arr)

	//var users = []models.Kategoris{{User_id: "1",Nama_kategori: "jinzhu1"},{User_id: "1",Nama_kategori: "jinzhu1"}}
	//var users = []models.Kategoris{jsonStringBayu}
	//fmt.Println("user",users)
	for _, datakategori := range modeldata {
		db.Create(&datakategori)
	}

	res.Status = 200
	res.Message = "Success"
	res.Data = m
	return c.JSON(http.StatusOK, res)

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
	
	fmt.Println(&jsonRegister)
	var yes uint=uint(jsonRegister["id"].(float64))
		//y :=i
	fmt.Println(yes)
	if yes != 0 {
		var addkategori = models.Kategoris{
			ID :yes,
			User_id: idUser,
			Nama_kategori: jsonRegister["nama_kategori"].(string),
	
		}
		result := db.Save(&addkategori)
		res.Status = http.StatusOK
		res.Message = "Success"
		res.Data = result.RowsAffected 
		return c.JSON(http.StatusOK, res)

	} else {
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



}
