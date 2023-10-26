package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/jmoiron/sqlx"
)

func API(db *sqlx.DB, engine *gin.Engine) {
	productHandler := Product{
		DB: db,
	}
	engine.POST("/v1/product", productHandler.CreateNewProduct)
	engine.GET("/v1/product/:id", productHandler.GetProductByID)
}
