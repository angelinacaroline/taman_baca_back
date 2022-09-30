package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"taman_baca_back/models"
)

func DataPengarang(c echo.Context) error { //Nama Function harus diawali huruf besar untuk bisa dipanggil di folder yang lain
	result, err := models.DataPengarang()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
