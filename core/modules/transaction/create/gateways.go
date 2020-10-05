package transaction_create

type TransactionGateway interface {
	Create(payer Payer, payee Payee, amount uint) (Transaction, error)
}

type PayerGateway interface {
	Find(id uint) (Payer, error)
}

type PayeeGateway interface {
	Find(id uint) (Payee, error)
}
