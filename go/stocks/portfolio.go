package stocks

type Portfolio struct {
	money []Money
}

func (p *Portfolio) Add(money Money) {
	p.money = append(p.money, money)
}

func (p *Portfolio) Get(b *Bank, c Currency) (Money, error) {
	total := float32(0)
	for _, m := range p.money {
		converted, err := b.Convert(m, c)
		if err != nil {
			return Money{}, err
		}
		total += converted.amount
	}
	return Money{total, c}, nil
}
