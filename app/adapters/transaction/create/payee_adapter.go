package create

import (
	transactionCreate "github.com/golang-clean-architecture/core/modules/transaction/create"
)

type PayeeAdapter struct {
}

func NewPayeeAdapter() PayeeAdapter {
	return PayeeAdapter{}
}

func (adapter PayeeAdapter) Find(id uint) (payee transactionCreate.Payee, err error) {
	payee.ID = id
	payee.Type = transactionCreate.SELLER
	return payee, err
}
