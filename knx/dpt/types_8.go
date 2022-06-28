// Copyright 2017 Ole Krüger.
// Copyright 2022 Martin Müller.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import (
	"fmt"
)

// DPT_8001 represents DPT 8.001 (G) / DPT_Value_2_Count.
type DPT_8001 int16

func (d DPT_8001) Pack() []byte {
	return packV16(int16(d))
}

func (d *DPT_8001) Unpack(data []byte) error {
	return unpackV16(data, (*int16)(d))
}

func (d DPT_8001) Unit() string {
	return "pulses"
}

func (d DPT_8001) String() string {
	return fmt.Sprintf("%d pulses", d)
}

// DPT_8012 represents DPT 8.012 (FB) / DPT_Length_m.
type DPT_8012 int16

func (d DPT_8012) Pack() []byte {
	return packV16(int16(d))
}

func (d *DPT_8012) Unpack(data []byte) error {
	return unpackV16(data, (*int16)(d))
}

func (d *DPT_8012) Unit() string {
	return "m"
}

func (d DPT_8012) String() string {
	return fmt.Sprintf("%d m", d)
}
