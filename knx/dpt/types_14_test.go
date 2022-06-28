package dpt

import (
	"fmt"
	"testing"

	"math"
	"math/rand"
)

func getRandValue() (float32, float32) {
	value := rand.Float32()

	// Scale the random number to the given range
	value *= 670760

	// Calculate the quantization error we expect
	Q := get_float_quantization_error(value, 0.01, 2047)
	return value, Q
}

// Test DPT 14.xxx (F32)
func TestDPT_14(t *testing.T) {
	type DPT14 struct {
		Dpv       DatapointValue
		Min       float32
		MinStr    string
		Middle    float32
		MiddleStr string
		Max       float32
		MaxStr    string
	}
	var types_14 = []DPT14{
		{new(DPT_14000), math.SmallestNonzeroFloat32, "0.00 m/s²", 0, "0.00 m/s²", math.MaxFloat32, "340282346638528859811704183484516925440.00 m/s²"},
	}
	for _, e := range types_14 {
		src := e.Dpv
		if fmt.Sprintf("%s", src) != e.MiddleStr {
			t.Errorf("%#v has wrong default (0) value [%v]. Should be [%s].", e.Dpv, e.Dpv, e.MiddleStr)
		}

		err := e.Dpv.Unpack(packF32(e.Min))
		if err != nil {
			t.Errorf("%v", err)
		}
		if fmt.Sprintf("%s", e.Dpv) != e.MinStr {
			t.Errorf("%#v has wrong smallest value [%v]. Should be [%s].", e.Dpv, e.Dpv, e.MinStr)
		}

		err = e.Dpv.Unpack(packF32(e.Max))
		if err != nil {
			t.Errorf("%v", err)
		}
		if fmt.Sprintf("%s", e.Dpv) != e.MaxStr {
			t.Errorf("%#v has wrong largest value [%v]. Should be [%s].", e.Dpv, e.Dpv, e.MaxStr)
		}
	}
}
