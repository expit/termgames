package fastReader

import (
	"fmt"
	"testing"
)

func TestGenerator(t *testing.T) {

	rangeTests := []struct {
		min int
		max int
		length float64
	}{
		{min: 0, max:9,length:1},
		{min: 10,max:99,length:2},
		{100,999,3},
		{100,9999,3.1},
		{100,9999,3.9},
		{ 1000, 9999, 4 },
		{ 1000000000, 9999999999, 10 },
		{ 100000000000000, 999999999999999, 15 },
	}

	for _, tt := range rangeTests{
		t.Run(fmt.Sprintf("%v",tt.length),func(t *testing.T){
			got:= NumberGenerator(tt.length)

			if got < tt.min || got > tt.max {
				t.Errorf("got %d, wanted something %d < x< %cd", got, tt.min,tt.max)
			}
		})
	}
}
