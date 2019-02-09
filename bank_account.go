package account

type BankAccount interface {
	Balance() float64
}

type bankAccount struct {
	balance float64
}

func Open(amt float64) *bankAccount {
	ba := &bankAccount{balance: amt}

	return ba
}

func (ba *bankAccount) Balance() (float64, bool) {
	return ba.balance, true
}

func (ba *bankAccount) Close() (float64, bool) {
	return ba.balance, true
}

func (ba *bankAccount) Deposit(depositAmount float64) (float64, bool) {
	ba.balance += depositAmount
	return ba.balance, true
}
