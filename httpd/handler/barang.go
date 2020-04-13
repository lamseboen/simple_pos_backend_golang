package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Barang struct {
	Barang_kode      string `json:barang_kode`
	Barang_nama      string `json:barang_nama`
	Barang_kategori  int    `json:barang_kategori`
	Barang_hargabeli int    `json:barang_hargabeli`
	Barang_hargajual int    `json:barang_hargajual`
	Barang_stok      int    `json:barang_stok`
	Barang_satuan    int    `json:barang_satuan`
}

func GetBarang(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var barang []Barang
		rows, err := db.Query(`SELECT barang_kode, barang_nama, barang_kategori, barang_hargabeli, barang_hargajual, barang_stok, barang_satuan FROM master.barang`)
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		for rows.Next() {
			brg := Barang{}
			err = rows.Scan(
				&brg.Barang_kode,
				&brg.Barang_nama,
				&brg.Barang_kategori,
				&brg.Barang_hargabeli,
				&brg.Barang_hargajual,
				&brg.Barang_stok,
				&brg.Barang_satuan,
			)
			if err != nil {
				panic(err)
			}
			barang = append(barang, brg)

			c.JSON(http.StatusOK, barang)
		}
	}
}

func GetBarangDetail(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var barang []Barang

		id := c.Param("id")
		fmt.Println(id)
		rows, err := db.Query(`SELECT barang_kode, barang_nama, barang_kategori, barang_hargabeli, barang_hargajual, barang_stok, barang_satuan FROM master.barang where barang_kode=$1`, id)
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		for rows.Next() {
			brg := Barang{}
			err = rows.Scan(
				&brg.Barang_kode,
				&brg.Barang_nama,
				&brg.Barang_kategori,
				&brg.Barang_hargabeli,
				&brg.Barang_hargajual,
				&brg.Barang_stok,
				&brg.Barang_satuan,
			)
			if err != nil {
				panic(err)
			}
			barang = append(barang, brg)

			c.JSON(http.StatusOK, barang)
		}
	}
}
