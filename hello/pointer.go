package main

//Wallet struct
type Wallet struct {
	Amount float64
}

//Deposit money into Wallet
func (w *Wallet) Deposit(money float64) {
	w.Amount += money
}

//Balance returns the current amount in wallet
func (w *Wallet) Balance() float64 {
	return w.Amount
}
