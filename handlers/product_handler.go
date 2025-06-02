package handlers

import (
	"E_commerce_System/config"
	"E_commerce_System/models"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetProducts(c *gin.Context) {
	// Define the slice to hold products
	var products []models.Product
	result := config.DB.Find(&products)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": result.Error.Error()})
		return
	}
	// If the table is empty, return a proper message
	if len(products) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "No products found",
			"data":    []models.Product{},
		})
		return
	}

	c.JSON(http.StatusOK, products)
}
func AddProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if product.Name == "" || product.Price < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Invalid product data"})
		return
	}
	result := config.DB.Create(&product)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, product)
}
func DelProduct(c *gin.Context) {
	var req models.DELProduct

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if req.Name == "" || req.Price < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Invalid product data"})
	}
	//find the product by Name+Price
	var product models.Product
	result := config.DB.Where("name = ? and price = ?", req.Name, req.Price).First(&product)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"err": "Product not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"err": result.Error.Error()})
		}
		return

	}
	//delete the product by ID
	if err := config.DB.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
		"deleted": product,
	})

}
func UpdateProduct(c *gin.Context) {
	var req models.UpdateProduct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Invalid product data"})
		return
	}
	if req.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Product ID is required"})
		return
	}
	//find the product by ID
	var product models.Product
	result := config.DB.Where("id = ?", req.ID).First(&product)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"err": "Product not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"err": result.Error.Error()})
		}
	}
	if req.Name != "" {
		product.Name = req.Name
	}
	if req.Price != 0 {
		product.Price = req.Price
	}
	//update the product
	if err := config.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Product updated successfully",
		"updated": product,
	})

}
