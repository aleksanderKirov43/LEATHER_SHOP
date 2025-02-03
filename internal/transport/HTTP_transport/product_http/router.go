package product_http

import (
	"github.com/gin-gonic/gin"
	"leather-shop/internal/transport/HTTP_transport"
)

func NewRouterProduct(engine *gin.RouterGroup, controller HTTP_transport.Products) {

	productGroup := engine.Group("/products")

	productGroup.POST("/product", controller.CreateProduct)
	productGroup.GET("/:id", controller.GetProduct)
	productGroup.GET("", controller.GetProducts)
	productGroup.DELETE("/:id", controller.DeleteProduct)
	productGroup.PUT("/:id", controller.EditProduct)
}
