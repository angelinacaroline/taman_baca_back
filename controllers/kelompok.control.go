package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"taman_baca_back/models"
)

func DataKelompok(c echo.Context) error { //Nama Function harus diawali huruf besar untuk bisa dipanggil di folder yang lain
	result, err := models.DataKelompok()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func SimpanDataKelompok(c echo.Context) error {
	id_kelompok := c.FormValue("id_kelompok")
	nama_kelompok := c.FormValue("nama_kelompok")
	id_kategori_kelompok := c.FormValue("id_kategori_kelompok")
	penanggungjawab := c.FormValue("penanggungjawab")
	no_telepon := c.FormValue("no_telepon")
	alamat := c.FormValue("alamat")
	tanggal_terdaftar := c.FormValue("tanggal_terdaftar")

	result, err := models.SimpanDataKelompok(id_kelompok, nama_kelompok, id_kategori_kelompok, penanggungjawab, no_telepon, alamat, tanggal_terdaftar)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func UbahDataKelompok(c echo.Context) error {
	id_kelompok := c.FormValue("id_kelompok")
	nama_kelompok := c.FormValue("nama_kelompok")
	id_kategori_kelompok := c.FormValue("id_kategori_kelompok")
	penanggungjawab := c.FormValue("penanggungjawab")
	no_telepon := c.FormValue("no_telepon")
	alamat := c.FormValue("alamat")

	result, err := models.UbahDataKelompok(id_kelompok, nama_kelompok, id_kategori_kelompok, penanggungjawab, no_telepon, alamat)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func HapusDataKelompok(c echo.Context) error {
	id_kelompok := c.Param("id_kelompok")

	result, err := models.HapusDataKelompok(id_kelompok)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func DataInventaris(c echo.Context) error { //Nama Function harus diawali huruf besar untuk bisa dipanggil di folder yang lain
	result, err := models.DataInventaris()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
