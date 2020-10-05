package transaction_find

import (
	"github.com/golang-clean-architecture/core/depedencies"
)

type Request struct {
	Id string
}

type Response struct {
	Transaction Transaction
}

type UseCase struct {
	transactionGateway TransactionGateway
	loggerGateway      depedencies.LoggerGateway
}

func NewUseCase(
	transactionGateway TransactionGateway,
	loggerGateway depedencies.LoggerGateway,
) UseCase {
	return UseCase{
		transactionGateway: transactionGateway,
		loggerGateway:      loggerGateway,
	}
}

func (useCase UseCase) Execute(request Request) (*Response, error) {
	useCase.loggerGateway.Info("Init transaction find use case", depedencies.Event{
		"request": request,
	})

	transaction, err := useCase.transactionGateway.Find(request.Id)
	if err != nil {
		return nil, TransactionFoundError
	}
	response := Response{
		Transaction: transaction,
	}

	useCase.loggerGateway.Info("Finish transaction find use case", depedencies.Event{
		"response": response,
	})

	return &response, nil
}
