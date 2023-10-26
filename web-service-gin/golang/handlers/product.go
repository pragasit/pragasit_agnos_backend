package handlers

import (
	"log"
	"net/http"
	"time"

	"./product"

	"github.com/gin-gonic/gin"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	DB        *sqlx.DB
	ProductDB product.ProductDB
}

func (prod Product) CreateNewProduct(context *gin.Context) {
	var newProduct product.NewProduct
	err := context.ShouldBindJSON(&newProduct)
	if err != nil {
		log.Println("CreateNewProduct ShouldBindJSON error: ", err)
		context.Status(http.StatusBadRequest)
		return
	}
	newItemp, err := prod.ProductDB.CreateNewProduct(newProduct, time.Now())
	if err != nil {
		log.Println("Handlers CreateNewProduct error: ", err)
		context.Status(http.StatusInternalServerError)
		return
	}
	log.Print("new", newItemp)
	context.JSON(http.StatusOK, newItemp)
}

func (prod Product) GetProductByID(context *gin.Context) {
	productID := context.Param("id")

	product, err := prod.ProductDB.GetProductByID(productID)
	if err != nil {
		log.Println("Handlers GetProductByID error: ", err)
		context.Status(http.StatusInternalServerError)
		return
	}
	context.JSON(http.StatusOK, product)
}
