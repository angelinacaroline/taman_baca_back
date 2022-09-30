package models

import (
	"fmt"
	"net/http"
	"taman_baca_back/db"
)

//Class untuk Tabel Kategori Buku
type Kategori_buku struct {
	Id_kategori_buku string `json:"id_kategori_buku"`
	Nama_kategori    string `json:"nama_kategori"`
}

//Select Data Kategori Buku
func DataKategoriBuku() (Response, error) {
	var obj Kategori_buku
	var arrobj []Kategori_buku
	var res Response

	con, err := db.DbConnection()

	sqlStatement := "SELECT * FROM kategori_buku"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()
	//fmt.Println(con)

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(
			&obj.Id_kategori_buku,
			&obj.Nama_kategori,
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
