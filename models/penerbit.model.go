package models

import (
	"fmt"
	"net/http"
	"taman_baca_back/db"
)

//Class untuk Tabel Penerbit
type Penerbit struct {
	Id_penerbit   string `json:"id_penerbit"`
	Nama_penerbit string `json:"nama_penerbit"`
	No_telepon    string `json:"no_telepon"`
	Email         string `json:"email"`
}

//Select Data Penerbit
func DataPenerbit() (Response, error) {
	var obj Penerbit
	var arrobj []Penerbit
	var res Response

	con, err := db.DbConnection()

	sqlStatement := "SELECT * FROM penerbit"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()
	//fmt.Println(con)

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(
			&obj.Id_penerbit,
			&obj.Nama_penerbit,
			&obj.No_telepon,
			&obj.Email,
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
