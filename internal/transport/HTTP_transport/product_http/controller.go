package product_http

import (
	"github.com/gin-gonic/gin"
	"leather-shop/internal/models"
	"leather-shop/internal/services"
	"leather-shop/internal/transport/HTTP_transport"
	"net/http"
	"strconv"
)

type productController struct {
	productService services.Products
}

func New(productService services.Products) HTTP_transport.Products {
	return &productController{
		productService: productService,
	}
}

func (pc *productController) GetProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	productId, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Ошибка": err.Error()})
		return
	}

	product, err := pc.productService.GetProduct(productId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Ошибка": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, product)
}
func (pc *productController) GetProducts(ctx *gin.Context) {
	products, err := pc.productService.GetProducts()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Ошибка": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}
func (pc *productController) CreateProduct(ctx *gin.Context) {
	var product models.Products
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Ошибка": err.Error()})
		return
	}

	if err := pc.productService.CreateProduct(&product); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Ошибка": err.Error()})
	}

	ctx.JSON(http.StatusOK, product)

}
func (pc *productController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	productId, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Ошибка": err.Error()})
		return
	}

	if err := pc.productService.DeleteProduct(productId); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Ошибка": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
func (pc *productController) EditProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	productId, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Ошибка": err.Error()})
		return
	}

	var product models.Products
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Ошибка": err.Error()})
		return
	}
	product.Id = productId

	if err := pc.productService.EditProduct(&product); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Ошибка": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, product)
}
