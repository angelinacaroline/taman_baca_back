package models

import (
	"fmt"
	"net/http"
	"taman_baca_back/db"
)

//Class untuk Tabel Pengarang
type Pengarang struct {
	Id_pengarang   string `json:"id_pengarang"`
	Nama_pengarang string `json:"nama_pengarang"`
}

//Select Data Pengarang
func DataPengarang() (Response, error) {
	var obj Pengarang
	var arrobj []Pengarang
	var res Response

	con, err := db.DbConnection()

	sqlStatement := "SELECT * FROM pengarang"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()
	//fmt.Println(con)

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(
			&obj.Id_pengarang,
			&obj.Nama_pengarang,
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
