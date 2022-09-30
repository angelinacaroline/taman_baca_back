package models

import (
	"fmt"
	"net/http"
	"taman_baca_back/db"
)

//Class untuk Tabel Kategori Kelompok
type Kategori_kelompok struct {
	Id_kategori_kelompok string `json:"id_kategori_kelompok"`
	Nama_kategori        string `json:"nama_kategori"`
}

//Select Data Kategori Kelompok
func DataKategoriKelompok() (Response, error) {
	var obj Kategori_kelompok
	var arrobj []Kategori_kelompok
	var res Response

	con, err := db.DbConnection()

	sqlStatement := "SELECT * FROM kategori_kelompok"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()
	//fmt.Println(con)

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(
			&obj.Id_kategori_kelompok,
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
