package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"taman_baca_back/controllers"
)

func Init() *echo.Echo { //Pakai bintang karena untuk membuat function dari class
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "welcome here")
	})

	//Method API : POST dan GET
	//Postman untuk test API

	//CRUD KATEGORI KELOMPOK:
	e.GET("/kategori_kelompok", controllers.DataKategoriKelompok) //View Data

	//CRUD KELOMPOK:
	e.GET("/kelompok", controllers.DataKelompok)                      //View Data
	e.POST("/kelompok", controllers.SimpanDataKelompok)               //Add Data
	e.PUT("/kelompok", controllers.UbahDataKelompok)                  //Edit Data
	e.DELETE("/kelompok/:id_kelompok", controllers.HapusDataKelompok) //Delete Data

	//CRUD KATEGORI INVENTARIS:
	e.GET("/kategori_inventaris", controllers.DataKategoriInventaris) //View Data

	//CRUD INVENTARIS:
	e.GET("/inventaris", controllers.DataInventaris) //View Data

	//CRUD PENGARANG:
	e.GET("/pengarang", controllers.DataPengarang) //View Data

	//CRUD PENERBIT:
	e.GET("/penerbit", controllers.DataPenerbit) //View Data

	//CRUD PENERBIT:
	e.GET("/kategori_buku", controllers.DataKategoriBuku) //View Data

	return e
}
