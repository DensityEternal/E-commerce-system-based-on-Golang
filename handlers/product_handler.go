package handlers

import (
	"E_commerce_System/config"
	"E_commerce_System/models"
	"E_commerce_System/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

//---------------------商品的增删改查---------------------//
//查询商品

func GetProducts(c *gin.Context) {
	// Define the slice to hold products
	var products []models.Product
	result := config.DB.Find(&products)
	if result.Error != nil {
		log.Printf("Database error: %s", result.Error.Error()) // 内部日志
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Internal Server Error"})
		return
	}
	// If the table is empty, return a proper message
	if len(products) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "No products found",
			"data":    products,
		})
		return
	}

	c.JSON(http.StatusOK, products)
}

//添加商品

func AddProduct(c *gin.Context) {
	var product models.Product
	// 解析 JSON 数据
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	// 检查商品是否有效
	if product.Name == "" || product.Price < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Invalid product data"})
		return
	}
	//检查商品是否已经存在过
	var existingProduct models.Product
	err := config.DB.Where("name = ? AND price = ?", product.Name, product.Price).First(&existingProduct).Error
	switch {
	case err == nil:
		c.JSON(http.StatusConflict, gin.H{"err": "Product already exists"})
		return
	case errors.Is(err, gorm.ErrRecordNotFound):
		// 商品不存在，允许创建
	case err != nil:
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	result := config.DB.Create(&product)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, product)
}

//删除商品

func DelProduct(c *gin.Context) {
	// 定义结构体用于接收数组形式的商品 ID
	var reqBody struct {
		IDs []uint `json:"id"` // 商品 ID 的数组
	}

	// 解析 JSON 数据
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Invalid request body"})
		return
	}

	// 检查是否提供了 ID
	if len(reqBody.IDs) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"err": "No product IDs provided"})
		return
	}

	// 删除商品
	result := config.DB.Delete(&models.Product{}, reqBody.IDs)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "Products deleted successfully",
		"deleted_count": result.RowsAffected,
	})
}

//更新商品

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
			return

		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"err": result.Error.Error()})
			return

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

//---------------------------注册与登录---------------------------//

//加密

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err

	}
	return string(bytes), err
}

//验证

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//---------------------------注册接口---------------------------//

func Register(c *gin.Context) {
	var req models.UserLogin
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Invalid request"})
		return
	}

	// 查找用户名是否已存在
	var existing models.User
	// 添加日志查看查询的用户名
	log.Printf("Checking username")

	err := config.DB.Where("user_name = ?", req.UserName).First(&existing).Error
	if err == nil {
		log.Printf("Username conflict: %s", req.UserName) // 添加日志
		c.JSON(http.StatusConflict, gin.H{"err": "Username already exists"})
		return
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// 数据库出错
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Database error"})
		return
	}

	// 加密密码
	hashPassword, hashErr := HashPassword(req.Password)
	if hashErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Error hashing password"})
		return
	}

	// 创建用户
	user := models.User{
		UserName: req.UserName,
		Password: hashPassword,
	}
	if result := config.DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// ---------------------------登录接口---------------------------//

func Login(c *gin.Context) {
	var req models.UserLogin
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Invalid request"})
		return
	}
	var user models.User
	LoginErr := config.DB.Where("user_name = ?", req.UserName).First(&user).Error
	if LoginErr != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Invalid username or password",
			"err": LoginErr.Error()})
		return
	}
	if !CheckPasswordHash(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid username or password,\nplease try again"})
		return
	}
	token, tokenErr := utils.GenerateJWT(user.ID)

	if tokenErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Failed to generate token"})
		return

	}
	c.JSON(http.StatusOK, gin.H{"message": "Login successfully",
		"token": token})

}
