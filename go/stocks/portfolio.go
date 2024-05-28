package stocks

type Portfolio struct {
	money []Money
}

func (p *Portfolio) Add(money Money) {
	p.money = append(p.money, money)
}

func (p *Portfolio) Get(c Currency) Money {
	sum := float32(0)
	for _, m := range p.money {
		if m.currency == c {
			sum += m.amount
		}
	}
	return Money{sum, c}
}
