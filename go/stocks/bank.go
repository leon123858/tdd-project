package stocks

import (
	"errors"
	"fmt"
)

type Bank struct {
	exchangeRates map[Currency]map[Currency]float32
}

func NewBank() *Bank {
	return &Bank{exchangeRates: map[Currency]map[Currency]float32{}}
}

func (b *Bank) AddExchangeRate(from, to Currency, rate float32) {
	if _, ok := b.exchangeRates[from]; !ok {
		b.exchangeRates[from] = map[Currency]float32{}
	}
	b.exchangeRates[from][to] = rate
}

func (b *Bank) Convert(from Money, to Currency) (Money, error) {
	if from.currency == to {
		return from, nil
	}
	rate, ok := b.exchangeRates[from.currency][to]
	if !ok {
		s := fmt.Sprintf("no exchange rate found: %s to %s", from.currency, to)
		return Money{}, errors.New(s)
	}
	newAmount := rate * from.amount
	return Money{newAmount, to}, nil
}
