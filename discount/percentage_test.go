package discount

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPercentageDiscount(t *testing.T) {
	var pd *PercentageDiscount
	var v float64

	v = 5
	pd = NewPercentageDiscount(v)

	assert.Equal(t, v, pd.Value())
}

func TestPercentageDiscount_Calculate(t *testing.T) {
	var pd *PercentageDiscount
	var v float64

	v = 5
	pd = NewPercentageDiscount(v)

	discountValue, _ := pd.Calculate(10000)

	assert.Equal(t, 9500.0, discountValue)
}
