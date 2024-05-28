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
	portfolio.Add(s.NewMoney(20, s.TWD))
	portfolio.Add(s.NewMoney(30, s.USD))
	dollars := portfolio.Get(s.USD)
	assert.Equal(t, s.NewMoney(40, s.USD), dollars)
}
