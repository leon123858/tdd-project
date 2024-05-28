package main

import (
	"github.com/stretchr/testify/assert"
	s "money/stocks"
	"reflect"
	"testing"
)

func TestMoney_times(t *testing.T) {
	type fields struct {
		amount   float32
		currency s.Currency
	}
	type args struct {
		multiplier float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   s.Money
	}{
		{
			name: "Test usd: 5*2",
			fields: fields{
				amount:   5,
				currency: s.USD,
			},
			args: args{
				multiplier: 2,
			},
			want: s.NewMoney(10, s.USD),
		},
		{
			name: "Test eur: 10*2",
			fields: fields{
				amount:   10,
				currency: s.EUR,
			},
			args: args{
				multiplier: 2,
			},
			want: s.NewMoney(20, s.EUR),
		},
		{
			name: "Test twd: 15.5*3",
			fields: fields{
				amount:   15.5,
				currency: s.TWD,
			},
			args: args{
				multiplier: 3,
			},
			want: s.NewMoney(46.5, s.TWD),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := s.NewMoney(tt.fields.amount, tt.fields.currency)
			d := &m
			if got := d.Times(tt.args.multiplier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("times() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoney_Divide(t *testing.T) {
	original := s.NewMoney(10, s.USD)
	result := original.Divide(2)
	assert.Equal(t, s.NewMoney(5, s.USD), result)
	original = s.NewMoney(10, s.EUR)
	result = original.Divide(4)
	assert.Equal(t, s.NewMoney(2.5, s.EUR), result)
	original = s.NewMoney(10, s.TWD)
	result = original.Divide(-5)
	assert.Equal(t, s.NewMoney(-2, s.TWD), result)
}

func TestMoney_Add(t *testing.T) {
	portfolio := s.Portfolio{}
	portfolio.Add(s.NewMoney(10, s.USD))
	portfolio.Add(s.NewMoney(30, s.USD))
	bank := s.NewBank()
	dollars, _ := portfolio.Get(bank, s.USD)
	assert.Equal(t, s.NewMoney(40, s.USD), dollars)
}

func TestMoney_AddMixedCurrencies(t *testing.T) {
	portfolio := s.Portfolio{}
	portfolio.Add(s.NewMoney(10, s.USD))
	portfolio.Add(s.NewMoney(20, s.TWD))
	portfolio.Add(s.NewMoney(30, s.USD))
	portfolio.Add(s.NewMoney(40, s.EUR))
	bank := s.NewBank()
	bank.AddExchangeRate(s.USD, s.TWD, 30.5)
	bank.AddExchangeRate(s.EUR, s.TWD, 40.5)
	bank.AddExchangeRate(s.USD, s.EUR, 0.8)
	bank.AddExchangeRate(s.TWD, s.EUR, 0.025)
	bank.AddExchangeRate(s.TWD, s.USD, 0.0328)
	bank.AddExchangeRate(s.EUR, s.USD, 1.25)
	dollars, _ := portfolio.Get(bank, s.TWD)
	assert.Equal(t, s.NewMoney(2860, s.TWD), dollars)
	dollars, _ = portfolio.Get(bank, s.EUR)
	assert.Equal(t, s.NewMoney(72.5, s.EUR), dollars)
	dollars, _ = portfolio.Get(bank, s.USD)
	assert.Equal(t, s.NewMoney(90.656, s.USD), dollars)
}

func TestMoney_ConvertNotExist(t *testing.T) {
	portfolio := s.Portfolio{}
	portfolio.Add(s.NewMoney(10, s.USD))
	portfolio.Add(s.NewMoney(20, s.TWD))
	portfolio.Add(s.NewMoney(30, s.USD))
	portfolio.Add(s.NewMoney(40, s.EUR))
	bank := s.NewBank()
	_, err := portfolio.Get(bank, s.JPY)
	assert.NotNil(t, err)
	assert.Equal(t, "no exchange rate found: USD to JPY", err.Error())

	portfolio = s.Portfolio{}
	portfolio.Add(s.NewMoney(50, s.JPY))
	_, e := portfolio.Get(bank, s.USD)
	assert.NotNil(t, e)
	assert.Equal(t, "no exchange rate found: JPY to USD", e.Error())
}

func TestMoney_Convert(t *testing.T) {
	bank := s.NewBank()
	bank.AddExchangeRate(s.USD, s.EUR, 0.8)
	m := s.NewMoney(10, s.USD)
	converted, e := bank.Convert(m, s.EUR)
	assert.Nil(t, e)
	assert.Equal(t, s.NewMoney(8, s.EUR), converted)
}

func TestMoney_ConvertSameCurrency(t *testing.T) {
	bank := s.NewBank()
	m := s.NewMoney(10, s.USD)
	converted, e := bank.Convert(m, s.USD)
	assert.Nil(t, e)
	assert.Equal(t, s.NewMoney(10, s.USD), converted)
}

func TestMoney_ConvertRateEdit(t *testing.T) {
	bank := s.NewBank()
	m := s.NewMoney(10, s.USD)
	bank.AddExchangeRate(s.USD, s.EUR, 0.8)
	converted, e := bank.Convert(m, s.EUR)
	assert.Nil(t, e)
	assert.Equal(t, s.NewMoney(8, s.EUR), converted)
	bank.AddExchangeRate(s.USD, s.EUR, 0.9)
	converted, e = bank.Convert(m, s.EUR)
	assert.Nil(t, e)
	assert.Equal(t, s.NewMoney(9, s.EUR), converted)
}
