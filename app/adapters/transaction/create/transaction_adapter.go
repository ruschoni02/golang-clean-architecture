package create

import (
	transactionCreate "github.com/golang-clean-architecture/core/modules/transaction/create"
	"github.com/google/uuid"
)

type TransactionAdapter struct {
}

func NewTransactionAdapter() TransactionAdapter {
	return TransactionAdapter{}
}

func (adapter TransactionAdapter) Create(
	payer transactionCreate.Payer,
	payee transactionCreate.Payee,
	amount uint,
) (transaction transactionCreate.Transaction, err error) {
	transaction.ID = uuid.New().String()
	transaction.Payee = payee
	transaction.Payer = payer
	transaction.Amount = amount
	return transaction, err
}
