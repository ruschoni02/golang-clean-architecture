package transaction_controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-clean-architecture/app/adapters"
	createAdapters "github.com/golang-clean-architecture/app/adapters/transaction/create"
	findAdapters "github.com/golang-clean-architecture/app/adapters/transaction/find"
	transactionCreate "github.com/golang-clean-architecture/core/modules/transaction/create"
	transactionFind "github.com/golang-clean-architecture/core/modules/transaction/find"
	"github.com/golang-clean-architecture/pkg/http"
)

type controller struct {
}

func NewController() *controller {
	return &controller{}
}

type HttpRequest struct {
	Amount uint `form:"amount" json:"amount" binding:"required"`
	Payer  uint `form:"payer" json:"payer" binding:"required"`
	Payee  uint `form:"payee" json:"payee" binding:"required"`
}

type responseData map[string]interface{}

// POST /create
func (controller *controller) Create(context *gin.Context) {
	var request HttpRequest
	if err := context.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		http.BadRequest(context, err)
		return
	}

	useCase := transactionCreate.NewUseCase(
		createAdapters.NewTransactionAdapter(),
		createAdapters.NewPayeeAdapter(),
		createAdapters.NewPayerAdapter(),
		adapters.NewLoggerAdapter(),
	)

	response, err := useCase.Execute(transactionCreate.Request{
		Amount: request.Amount,
		Payer:  request.Payer,
		Payee:  request.Payee,
	})

	if err != nil {
		http.InternalServerError(context, err)
		return
	}

	context.JSON(200, responseData{
		"transaction": responseData{
			"id": response.Transaction.ID,
			"payee": responseData{
				"id":   response.Transaction.Payee.ID,
				"type": response.Transaction.Payee.Type,
			},
			"payer": responseData{
				"id":   response.Transaction.Payer.ID,
				"type": response.Transaction.Payer.Type,
			},
			"amount": response.Transaction.Amount,
		},
	})
}

// POST /create
func (controller *controller) Find(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		http.BadRequest(context, errors.New("transaction id is required"))
		return
	}

	useCase := transactionFind.NewUseCase(
		findAdapters.NewTransactionAdapter(),
		adapters.NewLoggerAdapter(),
	)

	response, err := useCase.Execute(transactionFind.Request{
		Id: id,
	})

	if err != nil {
		http.InternalServerError(context, err)
		return
	}

	context.JSON(200, responseData{
		"transaction": responseData{
			"id": response.Transaction.ID,
			"payee": responseData{
				"id":   response.Transaction.Payee.ID,
				"type": response.Transaction.Payee.Type,
			},
			"payer": responseData{
				"id":   response.Transaction.Payer.ID,
				"type": response.Transaction.Payer.Type,
			},
			"amount": response.Transaction.Amount,
		},
	})
}
