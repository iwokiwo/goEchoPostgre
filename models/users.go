package models

import (
	//"github.com/jinzhu/gorm"
)

type Users struct {
	//gorm.Model
	Id string `json:"id"`
	Name string `json:"names"`
	Email string `json:"email"`
	Password string `json:"password"`
	Phone string `json:"phone"`
	Active_status int `json:"active_status"`
	Is_subscribe int `json:"is_subscibe"`
	Is_monthly int `json:"is_monthly"`
	Is_agent int `json:"is_agent"`
	Is_modal_kerja int `json:"is_modal_kerja"`
	//Token string `json:"token"`
	//Password string
}
