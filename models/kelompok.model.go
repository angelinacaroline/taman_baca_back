package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"taman_baca_back/db"
	"time"
)

// Class untuk Tabel Kelompok
type Kelompok struct {
	Id_kelompok          string `json:"id_kelompok"`
	Nama_kelompok        string `json:"nama_kelompok"`
	Id_kategori_kelompok string `json:"id_kategori_kelompok"`
	Penanggungjawab      string `json:"penanggungjawab"`
	No_telepon           string `json:"no_telepon"`
	Alamat               string `json:"alamat"`
	Tanggal_terdaftar    string `json:"tanggal_terdaftar"`
	Nama_kategori        string `json:"nama_kategori"`
}

// Select Data Kelompok
func DataKelompok() (Response, error) {
	var obj Kelompok
	var arrobj []Kelompok
	var res Response

	con, err := db.DbConnection()

	sqlStatement := "SELECT kelompok.id_kelompok, kelompok.nama_kelompok, kelompok.id_kategori_kelompok, penanggungjawab, kelompok.no_telepon, kelompok.alamat, kelompok.tanggal_terdaftar, kategori_kelompok.nama_kategori_kelompok FROM kelompok JOIN kategori_kelompok ON kelompok.id_kategori_kelompok = kategori_kelompok.id_kategori_kelompok ORDER BY id_kelompok ASC"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()
	//fmt.Println(con)

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(
			&obj.Id_kelompok,
			&obj.Nama_kelompok,
			&obj.Id_kategori_kelompok,
			&obj.Penanggungjawab,
			&obj.No_telepon,
			&obj.Alamat,
			&obj.Tanggal_terdaftar,
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

// Insert Data Kelompok
func SimpanDataKelompok(id_kelompok string, nama_kelompok string, id_kategori_kelompok string, penanggungjawab string, no_telepon string, alamat string,
	tanggal_terdaftar string) (Response, error) {
	var res Response

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	con, err := db.DbConnection()

	if err != nil {
		return res, err
	}
	defer cancelfunc()

	query := "INSERT INTO kelompok VALUES(?,?,?,?,?,?,?)"

	stmt, err := con.PrepareContext(ctx, query)
	if err != nil {
		return res, err
	}

	result, err := stmt.ExecContext(ctx, id_kelompok, nama_kelompok, id_kategori_kelompok, penanggungjawab, no_telepon, alamat, tanggal_terdaftar)
	if err != nil {
		return res, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return res, err
	}
	log.Printf("%d kelompok added ", rows)
	fmt.Println(result)

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]string{"getid": id_kelompok}

	return res, nil
}

// Update Data Kelompok
func UbahDataKelompok(id_kelompok string, nama_kelompok string, id_kategori_kelompok string, penanggungjawab string, no_telepon string, alamat string) (Response, error) {
	var res Response

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	con, err := db.DbConnection()

	if err != nil {
		return res, err
	}
	defer cancelfunc()

	query := "UPDATE kelompok SET nama_kelompok = ? , id_kategori_kelompok = ?, penanggungjawab = ? , no_telepon = ? , alamat = ? WHERE id_kelompok = ?"
	stmt, err := con.PrepareContext(ctx, query)

	if err != nil {
		return res, err
	}

	result, err := stmt.ExecContext(ctx, nama_kelompok, id_kategori_kelompok, penanggungjawab, no_telepon, alamat, id_kelompok)
	if err != nil {
		return res, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return res, err
	}
	log.Printf("%d kelompok updated ", rows)
	fmt.Println(result)

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows": rows,
	}

	return res, nil
}

// Delete Data Kelompok
func HapusDataKelompok(id_kelompok string) (Response, error) {
	var res Response

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	con, err := db.DbConnection()

	if err != nil {
		return res, err
	}
	defer cancelfunc()

	query := "DELETE FROM kelompok WHERE id_kelompok = ?"
	stmt, err := con.PrepareContext(ctx, query)

	if err != nil {
		return res, err
	}

	result, err := stmt.ExecContext(ctx, id_kelompok)
	if err != nil {
		return res, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return res, err
	}
	log.Printf("%d kelompok deleted ", rows)
	fmt.Println(result)

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows": rows,
	}

	return res, nil
}

////Insert Data Mitra
//func SimpanDataMitra(tmp Mitra) (Response, error) {
//	var res Response
//
//	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
//	con, err := db.DbConnection()
//
//	if err != nil {
//		return res, err
//	}
//	defer cancelfunc()
//	query := "INSERT INTO mitra Values(?,?,?,?,?,?)"
//
//	stmt, err := con.PrepareContext(ctx, query)
//	if err != nil {
//		return res, err
//	}
//	result, err := stmt.ExecContext(ctx, tmp.Id_mitra, tmp.Nama_mitra, tmp.Penanggungjawab, tmp.No_telepon, tmp.Tanggal_terdaftar)
//	if err != nil {
//		return res, err
//	}
//	rows, err := result.RowsAffected()
//	if err != nil {
//		log.Printf("Error %s when finding rows affected", err)
//		return res, err
//	}
//	log.Printf("%d mitra added ", rows)
//	fmt.Println(result)
//
//	res.Status = http.StatusOK
//	res.Message = "Success"
//	res.Data = map[string]string{"getid": tmp.Id_mitra}
//
//	return res, nil
//}

// Class untuk Tabel Inventaris
type Inventaris struct {
	Id_inventaris   string         `json:"id_inventaris"`
	Nama_inventaris string         `json:"nama_inventaris"`
	Stok_inventaris int            `json:"stok_inventaris"`
	Foto_inventaris sql.NullString `json:"foto_inventaris"`
}

// Select Data Inventaris
func DataInventaris() (Response, error) {
	var obj Inventaris
	var arrobj []Inventaris
	var res Response

	con, err := db.DbConnection()

	sqlStatement := "SELECT * FROM inventaris"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()
	//fmt.Println(con)

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(
			&obj.Id_inventaris,
			&obj.Nama_inventaris,
			&obj.Stok_inventaris,
			&obj.Foto_inventaris,
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
