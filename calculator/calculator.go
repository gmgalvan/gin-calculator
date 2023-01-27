package calculator

import (
	"github.com/gin-calculator/entities"
)

type Calculator struct{}

// NewService return a calculator strcut with methods
func NewService() *Calculator {
	return &Calculator{}
}

// Sum method sums quantities on operands struct and returns an int64 result
func (c *Calculator) Sum(op entities.Operands) int64 {
	var result int64
	for _, e := range op.Quantities {
		result += e
	}
	return result
}
