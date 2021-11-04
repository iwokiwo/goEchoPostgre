package models

type Response struct {
	Status  int         `json:"Code"`
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
}