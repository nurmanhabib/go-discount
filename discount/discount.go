package discount

type Discount interface {
	DisplayText() string
	Value() float64
	Calculate(float64) (float64, error)
}
