package transport

import (
	docs "github.com/etherwatt/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRoutes returns a gin engine router
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/etherwatt/v1"

	router.GET("/etherwatt/v1/etherwatt/wallet/:address", h.EtherWattWallet)
	router.GET("/etherwatt/v1/ethereum/wallet/:address", h.EtherWallet)

	router.POST("/etherwatt/v1/conection/:address", h.Conection)
	router.POST("/etherwatt/v1/transfer/etherwatt", h.Transfer)

	// use ginSwagger middleware to serve the API docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
