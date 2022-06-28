// Copyright 2017 Ole Krüger.
// Copyright 2022 Martin Müller.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import (
	"fmt"
)

// DPT_6001 represents DPT 6.001 (G) / DPT_Percent_V8.
type DPT_6001 int8

func (d DPT_6001) Pack() []byte {
	return packV8(int8(d))
}

func (d *DPT_6001) Unpack(data []byte) error {
	var value int8

	if err := unpackV8(data, &value); err != nil {
		return err
	}

	*d = DPT_6001(value)

	return nil
}

func (d DPT_6001) Unit() string {
	return "%"
}

func (d DPT_6001) String() string {
	return fmt.Sprintf("%d %%", int8(d))
}

// DPT_6010 represents DPT 6.010 (G) / DPT_Value_1_Count.
type DPT_6010 int8

func (d DPT_6010) Pack() []byte {
	return packV8(int8(d))
}

func (d *DPT_6010) Unpack(data []byte) error {
	var value int8

	if err := unpackV8(data, &value); err != nil {
		return err
	}

	*d = DPT_6010(value)

	return nil
}

func (d DPT_6010) Unit() string {
	return "counter pulses"
}

func (d DPT_6010) String() string {
	return fmt.Sprintf("%d counter pulses", int8(d))
}
