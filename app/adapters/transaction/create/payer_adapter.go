package create

import (
	transactionCreate "github.com/golang-clean-architecture/core/modules/transaction/create"
)

type PayerAdapter struct {
}

func NewPayerAdapter() PayerAdapter {
	return PayerAdapter{}
}

func (adapter PayerAdapter) Find(id uint) (payer transactionCreate.Payer, err error) {
	payer.ID = id
	payer.Type = transactionCreate.CONSUMER
	return payer, err
}
