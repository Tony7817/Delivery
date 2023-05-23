package test

import (
	"testing"
	"toaster/util"
)

var tests = []struct {
	weight float64
	price  float64
}{
	{0.4, 18},
	{2.2, 28.41},
}

func TestDelivery(t *testing.T) {
	for _, test := range tests {
		price, err := util.GetPrice(test.weight)
		if err != nil {
			t.Errorf("wrong weight %2.f", test.weight)
		}

		if price != test.price {
			t.Errorf("Expected %.2f but got %.2f", test.price, price)
		}
	}
}
