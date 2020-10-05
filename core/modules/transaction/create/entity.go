package transaction_create

const (
	CONSUMER = "CONSUMER"
	SELLER   = "SELLER"
)

type Transaction struct {
	ID     string
	Payee  Payee
	Payer  Payer
	Amount uint
}

type Payer struct {
	ID   uint
	Type string
}

type Payee struct {
	ID   uint
	Type string
}
