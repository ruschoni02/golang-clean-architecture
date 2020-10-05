package find

import (
	transactionFind "github.com/golang-clean-architecture/core/modules/transaction/find"
)

type TransactionAdapter struct {
}

func NewTransactionAdapter() TransactionAdapter {
	return TransactionAdapter{}
}

func (adapter TransactionAdapter) Find(
	id string,
) (transaction transactionFind.Transaction, err error) {
	transaction.ID = id
	transaction.Payee = transactionFind.Payee{
		ID:   1,
		Type: transactionFind.SELLER,
	}
	transaction.Payer = transactionFind.Payer{
		ID:   2,
		Type: transactionFind.CONSUMER,
	}
	transaction.Amount = 100
	return transaction, err
}
