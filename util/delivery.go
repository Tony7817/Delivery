package util

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

func GetPrice(weight float64) (float64, error) {
	if weight < 1 {
		return 18.00, nil
	}

	if weight > 100 {
		return 0, errors.New("maximum weight 100KG has reached")
	}

	weight = math.Floor(weight + 1)
	var price = 18.00

	for i := 1; i < int(weight); i++ {
		price += 5 + (price * 0.01)
	}

	price, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", price), 64)
	return price, nil
}
