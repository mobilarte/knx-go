// Copyright 2017 Ole Krüger.
// Copyright 2022 Martin Müller.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import (
	"fmt"
	"math"
	"testing"
)

// Test DPT 5.xxx (U8)
func TestDPT_5(t *testing.T) {
	type DPT5 struct {
		Dpv    DatapointValue
		Min    int
		MinStr string
		Max    int
		MaxStr string
	}

	var types_5 = []DPT5{
		{new(DPT_5001), 0, "0.00 %", 100, "100.00 %"},
		{new(DPT_5003), 0, "0°", 360, "360°"},
		{new(DPT_5004), 0, "0 %", 255, "255 %"},
		{new(DPT_5005), 0, "0", 255, "255"},
		{new(DPT_5006), 0, "0", 255, "reserved"},
		{new(DPT_5010), 0, "0 counter pulses", 255, "255 counter pulses"},
	}

	for _, e := range types_5 {
		src := e.Dpv
		if fmt.Sprintf("%s", src) != e.MinStr {
			t.Errorf("%#v has wrong default value [%v]. Should be [%s].", e.Dpv, e.Dpv, e.MinStr)
		}

		err := e.Dpv.Unpack(packU8(255))
		if err != nil {
			if fmt.Sprintf("%v", err) != "payload is not valid" {
				t.Errorf("%v", err)
			}
		}
		if fmt.Sprintf("%s", e.Dpv) != e.MaxStr {
			t.Errorf("%#v has wrong true value [%v]. Should be [%s].", e.Dpv, e.Dpv, e.MaxStr)
		}
	}

	// Compute the quantization error we expect. We should get less than that
	const Q = float32(360) / 255.0

	for i := 0; i <= 36000; i++ {
		value := float32(i / 100.0)
		src := DPT_5003(value)
		buf := src.Pack()

		var dst DPT_5003

		dst.Unpack(buf)
		if math.IsNaN(float64(dst)) {
			t.Errorf("Value [%s] is not a valid number. Original value was [%v].", dst, i)
		}
		if abs(float32(dst)-value) >= Q {
			t.Errorf("Value [%s] after pack/unpack above quantization noise! Original value was [%v], noise is [%f].", dst, value, Q)
		}
	}

	src := DPT_5003(320)
	var dst DPT_5003
	dst.Unpack(src.Pack())
	if dst.String() != "319°" {
		t.Errorf("Unpacked value is [%v]. Shoul be [319°].", dst)
	}
}
