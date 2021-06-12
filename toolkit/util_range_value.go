package toolkit

import "math/rand"

type RangeValue struct {
	From float64
	To   float64
}

func (rv RangeValue) CreateRand() float64 {
	value := rand.Float64()
	return value*(rv.To-rv.From) + rv.From
}
