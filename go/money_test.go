package main

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMoney_times(t *testing.T) {
	type fields struct {
		amount   float32
		currency currency
	}
	type args struct {
		multiplier float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Money
	}{
		{
			name: "Test usd: 5*2",
			fields: fields{
				amount:   5,
				currency: USD,
			},
			args: args{
				multiplier: 2,
			},
			want: Money{
				amount:   10,
				currency: USD,
			},
		},
		{
			name: "Test eur: 10*2",
			fields: fields{
				amount:   10,
				currency: EUR,
			},
			args: args{
				multiplier: 2,
			},
			want: Money{
				amount:   20,
				currency: EUR,
			},
		},
		{
			name: "Test twd: 15.5*3",
			fields: fields{
				amount:   15.5,
				currency: TWD,
			},
			args: args{
				multiplier: 3,
			},
			want: Money{
				amount:   46.5,
				currency: TWD,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Money{
				amount:   tt.fields.amount,
				currency: tt.fields.currency,
			}
			if got := d.times(tt.args.multiplier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("times() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoney_Divide(t *testing.T) {
	original := Money{amount: 10, currency: USD}
	result := original.Divide(2)
	assert.Equal(t, Money{amount: 5, currency: USD}, result)
	original = Money{amount: 10, currency: EUR}
	result = original.Divide(4)
	assert.Equal(t, Money{amount: 2.5, currency: EUR}, result)
	original = Money{amount: 10, currency: TWD}
	result = original.Divide(-5)
	assert.Equal(t, Money{amount: -2, currency: TWD}, result)
}

func TestMoney_Add(t *testing.T) {
	portfolio := Portfolio{}
	portfolio.Add(Money{amount: 20, currency: USD})
	portfolio.Add(Money{amount: 20, currency: EUR})
	portfolio.Add(Money{amount: 30, currency: TWD})
	dollars := portfolio.Get(USD)
	assert.Equal(t, Money{amount: 20, currency: USD}, dollars)
}
