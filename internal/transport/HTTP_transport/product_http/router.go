package product_http

import (
	"github.com/gin-gonic/gin"
	"leather-shop/internal/transport/HTTP_transport"
)

func NewRouterProduct(engine *gin.RouterGroup, controller HTTP_transport.Products) {

	engine.POST("/product", controller.CreateProduct)
	engine.GET("/:id", controller.GetProduct)
	engine.GET("", controller.GetProducts)
	engine.DELETE("/:id", controller.DeleteProduct)
	engine.PUT("/:id", controller.EditProduct)
}
