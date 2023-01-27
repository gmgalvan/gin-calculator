package calculator

import (
	"testing"

	"github.com/gin-calculator/entities"
)

func TestCalculator_Sum(t *testing.T) {
	type args struct {
		op entities.Operands
	}
	tests := []struct {
		name string
		c    *Calculator
		args args
		want int64
	}{
		{
			name: "Correct sum total",
			c:    &Calculator{},
			args: args{
				op: entities.Operands{
					Quantities: []int64{1, 2, 45, 6},
				},
			},
			want: 54,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calculator{}
			if got := c.Sum(tt.args.op); got != tt.want {
				t.Errorf("Calculator.Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
