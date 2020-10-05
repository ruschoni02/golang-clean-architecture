package transaction_find

type TransactionGateway interface {
	Find(id string) (Transaction, error)
}
