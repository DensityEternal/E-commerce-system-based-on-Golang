package main

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var products = []Product{
	{ID: 1, Name: "Laptop", Price: 999.99},
	{ID: 2, Name: "Phone", Price: 499.9},
}
var productMap = make(map[int]Product)

func getProducts(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}
func getProductID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		log.Printf("Invalid Product ID: %s, error: %v", idStr, err)
		return
	}
	if product, ok := productMap[id]; ok {
		c.JSON(http.StatusOK, product)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
	}

}
func addProduct(c *gin.Context) {
	log.Println("addProduct endpoint triggered")
	var newProduct Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Printf("ERROR:%v", err)
		return

	}
	if newProduct.Name == "" || newProduct.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product data"})
		log.Printf("Invalid product data: %+v", newProduct)
		return
	}

	newProduct.ID = len(products) + 1
	products = append(products, newProduct)
	c.JSON(http.StatusCreated, newProduct)
}
func keepUniqueProduct() {
	for _, p := range products {
		productMap[p.ID] = p
	}
}
func main() {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
	keepUniqueProduct()
	r.GET("/productMap", getProducts)
	r.GET("/productMap/:id", getProductID)
	r.POST("/productMap", addProduct)
	err := r.Run(":8090")
	if err != nil {
		log.Fatalf("ERR:%v", err)
	}
}
