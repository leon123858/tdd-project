package stocks

type Currency string

const (
	USD Currency = "USD"
	EUR Currency = "EUR"
	TWD Currency = "TWD"
)

type Money struct {
	amount   float32
	currency Currency
}

func NewMoney(amount float32, currency Currency) Money {
	return Money{amount, currency}
}

func (d *Money) Times(multiplier float32) Money {
	return Money{d.amount * multiplier, d.currency}
}

func (d *Money) Divide(m float32) Money {
	return Money{d.amount / m, d.currency}
}
