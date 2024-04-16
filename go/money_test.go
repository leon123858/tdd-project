package main

import (
	"reflect"
	"testing"
)

func TestMoney_times(t *testing.T) {
	type fields struct {
		amount   int
		currency currency
	}
	type args struct {
		multiplier int
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
