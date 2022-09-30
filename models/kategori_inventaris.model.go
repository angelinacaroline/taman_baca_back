package models

import (
	"fmt"
	"net/http"
	"taman_baca_back/db"
)

//Class untuk Tabel Kategori Inventaris
type Kategori_inventaris struct {
	Id_kategori_inventaris   string `json:"id_kategori_inventaris"`
	Nama_kategori_inventaris string `json:"nama_kategori_inventaris"`
}

//Select Data Kategori Inventaris
func DataKategoriInventaris() (Response, error) {
	var obj Kategori_inventaris
	var arrobj []Kategori_inventaris
	var res Response

	con, err := db.DbConnection()

	sqlStatement := "SELECT * FROM kategori_inventaris"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()
	//fmt.Println(con)

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(
			&obj.Id_kategori_inventaris,
			&obj.Nama_kategori_inventaris,
		)

		if err != nil {
			fmt.Println("Error : " + err.Error())
			return res, err
		}

		arrobj = append(arrobj, obj)
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}
