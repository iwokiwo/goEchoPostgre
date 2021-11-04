package controllers

import (
	"net/http"

	"hotels/db"
	"hotels/models"

	"github.com/labstack/echo"
)

type result struct {
	ID          int     `json:"id"`
	Product_id  int     `json:"product_id"`
	Nama_produk *string `json:"nm_produk"`
}

func Gettrsdtls(c echo.Context) error {
	//var obj result
	var res models.Response
	var arrobj []result
	db := db.DbManager()

	//--------------------------Cara Pertama--------------------------------------
	//rows, err := db.Table("transaksi_dtls").Select("transaksi_dtls.id, Transaksi_dtls.product_id,products.nama_produk").Joins("left join products on Transaksi_dtls.product_id = products.id").Rows()
	//	for rows.Next() {
	//	err = rows.Scan(&obj.ID, &obj.Product_id, &obj.Nama_produk)
	//	if err != nil {
	//		return err
	//	}

	//	arrobj = append(arrobj, obj)
	//}

	db.Table("transaksi_dtls").Select("transaksi_dtls.id, Transaksi_dtls.product_id,products.nama_produk").Joins("left join products on Transaksi_dtls.product_id = products.id").Scan(&arrobj)
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj
	return c.JSON(http.StatusOK, res)
}
