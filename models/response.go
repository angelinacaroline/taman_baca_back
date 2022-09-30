package models

//untuk membuat format returnnya seperti apa di browser untuk memudahkan membaca API nya

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
