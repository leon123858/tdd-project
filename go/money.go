package main

type currency string

const (
	USD currency = "USD"
	EUR currency = "EUR"
	TWD currency = "TWD"
)

type Money struct {
	amount   int
	currency currency
}

func (d *Money) times(multiplier int) Money {
	return Money{d.amount * multiplier, d.currency}
}

func main() {
	println("Hello, Money!")
}
