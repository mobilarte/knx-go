// Copyright 2017 Ole Krüger.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import (
	"fmt"
)

// DPT_18001 represents DPT 18.001 (G) / DPT_SceneControl.
type DPT_18001 uint8

func (d DPT_18001) Pack() []byte {
	if d <= 63 || (d >= 128 && d <= 191) {
		return packU8(uint8(d))
	} else {
		return packU8(63)
	}
}

func (d *DPT_18001) Unpack(data []byte) error {
	var value uint8

	if err := unpackU8(data, &value); err != nil {
		return err
	}

	*d = DPT_18001(value)
	if !d.IsValid() {
		return ErrOutOfRange
	}
	return nil
}

func (d DPT_18001) Unit() string {
	return ""
}

func (d DPT_18001) IsValid() bool {
	return d <= 63 || (d >= 128 && d <= 191)
}

// KNX Association recommends to display the scene numbers [1..64].
// Displaying value like ETS.
// See note 6 of the KNX Specifications v2.1.
func (d DPT_18001) String() string {
	if d >= 128 && d <= 191 {
		return fmt.Sprintf("learn %d", uint8(d)-128+1)
	} else if d <= 63 {
		return fmt.Sprintf("activate %d", uint8(d)+1)
	}
	return "invalid payload"
}
