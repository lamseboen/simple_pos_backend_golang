package main

import (
	"database/sql"
	"simple_pos_lib/httpd/handler"
	"simple_pos_lib/models"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func main() {
	db = models.ConnectDB()

	router := gin.Default()

	defaultName := router.Group("/simple_pos/api")

	barang := defaultName.Group("/barang")
	{
		barang.GET("/", handler.GetBarang(db))
		barang.GET("/:id", handler.GetBarangDetail(db))

	}

	router.Run(":8111")
}
