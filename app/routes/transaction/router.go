package transaction_route

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/golang-clean-architecture/app/controllers/transaction"
	"github.com/golang-clean-architecture/pkg/config"
	"github.com/golang-clean-architecture/pkg/http"
)

type transactionHandler struct {
}

func (handler *transactionHandler) Handler(router gin.IRouter, config *config.Config, server *http.Server) error {
	controller := controllers.NewController()
	router.POST("/transactions", controller.Create)
	router.GET("/transactions/:id", controller.Find)

	return nil
}

func New() http.Endpoint {
	return &transactionHandler{}
}
