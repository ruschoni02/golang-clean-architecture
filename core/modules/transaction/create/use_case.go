package transaction_create

import (
	"github.com/golang-clean-architecture/core/depedencies"
)

type Request struct {
	Amount uint
	Payer  uint
	Payee  uint
}

type Response struct {
	Transaction Transaction
}

type UseCase struct {
	transactionGateway TransactionGateway
	payeeGateway       PayeeGateway
	payerGateway       PayerGateway
	loggerGateway      depedencies.LoggerGateway
}

func NewUseCase(
	transactionGateway TransactionGateway,
	payeeGateway PayeeGateway,
	payerGateway PayerGateway,
	loggerGateway depedencies.LoggerGateway,
) UseCase {
	return UseCase{
		transactionGateway: transactionGateway,
		payeeGateway:       payeeGateway,
		payerGateway:       payerGateway,
		loggerGateway:      loggerGateway,
	}
}

func (useCase UseCase) Execute(request Request) (*Response, error) {
	useCase.loggerGateway.Info("Init transaction create use case", depedencies.Event{
		"request": request,
	})

	if request.Amount < 1 {
		return nil, AmountNotValidError
	}

	payer, err := useCase.payerGateway.Find(request.Payer)
	if err != nil {
		useCase.loggerGateway.Error("Error to find payer", err)
		return nil, PayerNotFoundError
	}

	payee, err := useCase.payeeGateway.Find(request.Payee)
	if err != nil {
		useCase.loggerGateway.Error("Error to find payee", err)
		return nil, PayeeNotFoundError
	}

	if payer.Type == SELLER {
		return nil, SellerCantBePayerError
	}

	transaction, err := useCase.transactionGateway.Create(payer, payee, request.Amount)
	if err != nil {
		return nil, CreateTransactionError
	}
	response := Response{
		Transaction: transaction,
	}

	useCase.loggerGateway.Info("Finish transaction create use case", depedencies.Event{
		"response": response,
	})

	return &response, nil
}
