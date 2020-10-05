package transaction_create

import "errors"

var PayerNotFoundError = errors.New("payer not found")
var PayeeNotFoundError = errors.New("payee not found")
var AmountNotValidError = errors.New("amount not valid")
var SellerCantBePayerError = errors.New("seller cant be payer")
var CreateTransactionError = errors.New("error to create transaction")
