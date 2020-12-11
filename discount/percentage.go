package discount

import "strconv"

type PercentageDiscount struct {
	value float64
}

var _ Discount = PercentageDiscount{}

func NewPercentageDiscount(percentage float64) *PercentageDiscount {
	return &PercentageDiscount{percentage}
}

func (p PercentageDiscount) DisplayText() string {
	return strconv.FormatFloat(p.Value(), 'f', 6, 64) + " %"
}

func (p PercentageDiscount) Value() float64 {
	return p.value
}

func (p PercentageDiscount) Calculate(f float64) (float64, error) {
	var result float64

	discount := f * (p.Value() / 100)
	result = f - discount

	return result, nil
}
