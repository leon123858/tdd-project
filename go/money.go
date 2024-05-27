package main

type currency string

const (
	USD currency = "USD"
	EUR currency = "EUR"
	TWD currency = "TWD"
)

type Money struct {
	amount   float32
	currency currency
}

type Portfolio struct {
	money []Money
}

func (p *Portfolio) Add(money Money) {
	p.money = append(p.money, money)
}

func (p *Portfolio) Get(c currency) Money {
	sum := float32(0)
	for _, m := range p.money {
		if m.currency == c {
			sum += m.amount
		}
	}
	return Money{sum, c}
}

func (d *Money) times(multiplier float32) Money {
	return Money{d.amount * multiplier, d.currency}
}

func (d *Money) Divide(m float32) Money {
	return Money{d.amount / m, d.currency}
}

func main() {
	println("Hello, Money!")
}
