// Copyright 2017 Ole Krüger.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import (
	"fmt"
)

// DPT_9001 represents DPT 9.001 (G) / DPT_Value_Temp.
type DPT_9001 float64

func (d DPT_9001) Pack() []byte {
	if d <= -273 {
		return packF16(-273.12)
	} else if d > 670433.28 {
		return packF16(670433.28)
	} else {
		return packF16(float64(d))
	}
}

func (d *DPT_9001) Unpack(data []byte) error {
	var value float64
	if err := unpackF16(data, &value); err != nil {
		return err
	}

	if value < -273 {
		value = -273
	}
	// Check the value for valid range
	if value < -273 || value > 670433.28 {
		return ErrOutOfRange
	}

	*d = DPT_9001(value)
	return nil
}

func (d DPT_9001) Unit() string {
	return "°C"
}

func (d DPT_9001) String() string {
	return fmt.Sprintf("%.2f °C", float64(d))
}

// DPT_9002 represents DPT 9.002 (G) / DPT_Value_Tempd.
type DPT_9002 float64

func (d DPT_9002) Pack() []byte {
	if d < -671088.64 {
		return packF16(-671088.64)
	} else if d > 670433.28 {
		return packF16(670433.28)
	} else {
		return packF16(float64(d))
	}
}

func (d *DPT_9002) Unpack(data []byte) error {
	var value float64
	if err := unpackF16(data, &value); err != nil {
		return err
	}

	// Check the value for valid range
	if value < -671088.64 || value > 670433.28 {
		return ErrOutOfRange
	}

	*d = DPT_9002(value)
	return nil
}

func (d DPT_9002) Unit() string {
	return "K"
}

func (d DPT_9002) String() string {
	return fmt.Sprintf("%.2f K", float64(d))
}

// DPT_9003 represents DPT 9.003 (G) / DPT_Value_Tempa.
type DPT_9003 float64

func (d DPT_9003) Pack() []byte {
	if d < -671088.64 {
		return packF16(-671088.64)
	} else if d > 670433.28 {
		return packF16(670433.28)
	} else {
		return packF16(float64(d))
	}
}

func (d *DPT_9003) Unpack(data []byte) error {
	var value float64
	if err := unpackF16(data, &value); err != nil {
		return err
	}

	// Check the value for valid range
	if value < -671088.64 || value > 670433.28 {
		return ErrOutOfRange
	}

	*d = DPT_9003(value)
	return nil
}

func (d DPT_9003) Unit() string {
	return "K/h"
}

func (d DPT_9003) String() string {
	return fmt.Sprintf("%.2f K/h", float64(d))
}

// DPT_9004 represents DPT 9.004 (G) / DPT_Value_Lux.
type DPT_9004 float64

func (d DPT_9004) Pack() []byte {
	if d <= 0 {
		return packF16(0)
	} else if d > 670433.28 {
		return packF16(670433.28)
	} else {
		return packF16(float64(d))
	}
}

func (d *DPT_9004) Unpack(data []byte) error {
	var value float64
	if err := unpackF16(data, &value); err != nil {
		return err
	}

	// Check the value for valid range
	if value < 0 || value > 670433.28 {
		return ErrOutOfRange
	}

	*d = DPT_9004(value)
	return nil
}

func (d DPT_9004) Unit() string {
	return "Lux"
}

func (d DPT_9004) String() string {
	return fmt.Sprintf("%.2f Lux", float64(d))
}

// DPT_9005 represents DPT 9.005 (G) / DPT_Value_Wsp.
type DPT_9005 float64

func (d DPT_9005) Pack() []byte {
	if d <= 0 {
		return packF16(0)
	} else if d > 670433.28 {
		return packF16(670433.28)
	} else {
		return packF16(float64(d))
	}
}

func (d *DPT_9005) Unpack(data []byte) error {
	var value float64

	if err := unpackF16(data, &value); err != nil {
		return err
	}

	// Check the value for valid range
	if value < 0 || value > 670433.28 {
		return ErrOutOfRange
	}

	*d = DPT_9005(value)
	return nil
}

func (d DPT_9005) Unit() string {
	return "m/s"
}

func (d DPT_9005) String() string {
	return fmt.Sprintf("%.2f m/s", float64(d))
}

// DPT_9006 represents DPT 9.006 (G) / DPT_Value_Pres.
type DPT_9006 float64

func (d DPT_9006) Pack() []byte {
	if d <= 0 {
		return packF16(0)
	} else if d > 670433.28 {
		return packF16(670433.28)
	} else {
		return packF16(float64(d))
	}
}

func (d *DPT_9006) Unpack(data []byte) error {
	var value float64

	if err := unpackF16(data, &value); err != nil {
		return err
	}

	// Check the value for valid range
	if value < 0 || value > 670433.28 {
		return ErrOutOfRange
	}

	*d = DPT_9006(value)
	return nil
}

func (d DPT_9006) Unit() string {
	return "Pa"
}

func (d DPT_9006) String() string {
	return fmt.Sprintf("%.2f Pa", float64(d))
}

// DPT_9007 represents DPT 9.007 (G) / DPT_Value_Humidity.
type DPT_9007 float64

func (d DPT_9007) Pack() []byte {
	if d <= 0 {
		return packF16(0)
	} else if d > 670433.28 {
		return packF16(670433.28)
	} else {
		return packF16(float64(d))
	}
}

func (d *DPT_9007) Unpack(data []byte) error {
	var value float64

	if err := unpackF16(data, &value); err != nil {
		return err
	}

	// Check the value for valid range
	if value < 0 || value > 670433.28 {
		return ErrOutOfRange
	}

	*d = DPT_9007(value)
	return nil
}

func (d DPT_9007) Unit() string {
	return "%"
}

func (d DPT_9007) String() string {
	return fmt.Sprintf("%.2f %%", float64(d))
}

// DPT_9008 represents DPT 9.008 (G) / DPT_Value_AirQuality.
type DPT_9008 float64

func (d DPT_9008) Pack() []byte {
	if d <= 0 {
		return packF16(0)
	} else if d > 670433.28 {
		return packF16(670433.28)
	} else {
		return packF16(float64(d))
	}
}

func (d *DPT_9008) Unpack(data []byte) error {
	var value float64

	if err := unpackF16(data, &value); err != nil {
		return err
	}

	// Check the value for valid range
	if value < 0 || value > 670433.28 {
		return ErrOutOfRange
	}

	*d = DPT_9008(value)

	return nil
}

func (d DPT_9008) Unit() string {
	return "ppm"
}

func (d DPT_9008) String() string {
	return fmt.Sprintf("%.2f ppm", float64(d))
}

// DPT_9009 represents DPT 9.009 (G) / DPT_Value_AirFlow.
type DPT_9009 float64

func (d DPT_9009) Pack() []byte {
	if d <= 0 {
		return packF16(0)
	} else if d > 670433.28 {
		return packF16(670433.28)
	} else {
		return packF16(float64(d))
	}
}

func (d *DPT_9009) Unpack(data []byte) error {
	var value float64

	if err := unpackF16(data, &value); err != nil {
		return err
	}

	// Check the value for valid range
	if value < 0 || value > 670433.28 {
		return ErrOutOfRange
	}

	*d = DPT_9009(value)

	return nil
}

func (d DPT_9009) Unit() string {
	return "m³/h"
}

func (d DPT_9009) String() string {
	return fmt.Sprintf("%.2f m³/h", float64(d))
}

// DPT_9010 represents DPT 9.010 (FB) / DPT_Value_Time1.
type DPT_9010 float64

func (d DPT_9010) Pack() []byte {
	if d < -671088.64 {
		return packF16(-671088.64)
	} else if d > 670433.28 {
		return packF16(670433.28)
	} else {
		return packF16(float64(d))
	}
}

func (d *DPT_9010) Unpack(data []byte) error {
	var value float64
	if err := unpackF16(data, &value); err != nil {
		return err
	}

	// Check the value for valid range
	if value < -671088.64 || value > 670433.28 {
		return ErrOutOfRange
	}

	*d = DPT_9010(value)
	return nil
}

func (d DPT_9010) Unit() string {
	return "s"
}

func (d DPT_9010) String() string {
	return fmt.Sprintf("%.2f s", float64(d))
}

// DPT_9011 represents DPT 9.011 (G) / DPT_Value_Time2.
type DPT_9011 float64

func (d DPT_9011) Pack() []byte {
	if d < -671088.64 {
		return packF16(-671088.64)
	} else if d > 670433.28 {
		return packF16(670433.28)
	} else {
		return packF16(float64(d))
	}
}

func (d *DPT_9011) Unpack(data []byte) error {
	var value float64
	if err := unpackF16(data, &value); err != nil {
		return err
	}

	// Check the value for valid range
	if value < -671088.64 || value > 670433.28 {
		return ErrOutOfRange
	}

	*d = DPT_9011(value)
	return nil
}

func (d DPT_9011) Unit() string {
	return "ms"
}

func (d DPT_9011) String() string {
	return fmt.Sprintf("%.2f ms", float64(d))
}

// DPT_9020 represents DPT 9.020 (G) / DPT_Value_Volt.
type DPT_9020 float64

func (d DPT_9020) Pack() []byte {
	if d < -671088.64 {
		return packF16(-671088.64)
	} else if d > 670433.28 {
		return packF16(670433.28)
	} else {
		return packF16(float64(d))
	}
}

func (d *DPT_9020) Unpack(data []byte) error {
	var value float64
	if err := unpackF16(data, &value); err != nil {
		return err
	}

	// Check the value for valid range
	if value < -671088.64 || value > 670433.28 {
		return ErrOutOfRange
	}

	*d = DPT_9020(value)
	return nil
}

func (d DPT_9020) Unit() string {
	return "mV"
}

func (d DPT_9020) String() string {
	return fmt.Sprintf("%.2f mV", float64(d))
}

// DPT_9021 represents DPT 9.021 (G) / DPT_Value_Curr.
type DPT_9021 float64

func (d DPT_9021) Pack() []byte {
	if d < -671088.64 {
		return packF16(-671088.64)
	} else if d > 670433.28 {
		return packF16(670433.28)
	} else {
		return packF16(float64(d))
	}
}

func (d *DPT_9021) Unpack(data []byte) error {
	var value float64
	if err := unpackF16(data, &value); err != nil {
		return err
	}

	// Check the value for valid range
	if value < -671088.64 || value > 670433.28 {
		return ErrOutOfRange
	}

	*d = DPT_9021(value)
	return nil
}

func (d DPT_9021) Unit() string {
	return "mA"
}

func (d DPT_9021) String() string {
	return fmt.Sprintf("%.2f mA", float64(d))
}

// DPT_9022 represents DPT 9.022 (FB) / DPT_PowerDensity.
type DPT_9022 float64

func (d DPT_9022) Pack() []byte {
	if d < -671088.64 {
		return packF16(-671088.64)
	} else if d > 670433.28 {
		return packF16(670433.28)
	} else {
		return packF16(float64(d))
	}
}

func (d *DPT_9022) Unpack(data []byte) error {
	var value float64
	if err := unpackF16(data, &value); err != nil {
		return err
	}

	// Check the value for valid range
	if value < -671088.64 || value > 670433.28 {
		return ErrOutOfRange
	}

	*d = DPT_9022(value)
	return nil
}

func (d DPT_9022) Unit() string {
	return "W/m²"
}

func (d DPT_9022) String() string {
	return fmt.Sprintf("%.2f W/m²", float64(d))
}

// DPT_9023 represents DPT 9.023 (FB) / DPT_KelvinPerPercent.
type DPT_9023 float64

func (d DPT_9023) Pack() []byte {
	if d < -671088.64 {
		return packF16(-671088.64)
	} else if d > 670433.28 {
		return packF16(670433.28)
	} else {
		return packF16(float64(d))
	}
}

func (d *DPT_9023) Unpack(data []byte) error {
	var value float64
	if err := unpackF16(data, &value); err != nil {
		return err
	}

	// Check the value for valid range
	if value < -671088.64 || value > 670433.28 {
		return ErrOutOfRange
	}

	*d = DPT_9023(value)
	return nil
}

func (d DPT_9023) Unit() string {
	return "K/%"
}

func (d DPT_9023) String() string {
	return fmt.Sprintf("%.2f K/%%", float64(d))
}

// DPT_9024 represents DPT 9.024 (FB) / DPT_Power.
type DPT_9024 float64

func (d DPT_9024) Pack() []byte {
	if d < -671088.64 {
		return packF16(-671088.64)
	} else if d > 670433.28 {
		return packF16(670433.28)
	} else {
		return packF16(float64(d))
	}
}

func (d *DPT_9024) Unpack(data []byte) error {
	var value float64
	if err := unpackF16(data, &value); err != nil {
		return err
	}

	// Check the value for valid range
	if value < -671088.64 || value > 670433.28 {
		return ErrOutOfRange
	}

	*d = DPT_9024(value)
	return nil
}

func (d DPT_9024) Unit() string {
	return "kW"
}

func (d DPT_9024) String() string {
	return fmt.Sprintf("%.2f kW", float64(d))
}

// DPT_9025 represents DPT 9.025 (FB) / DPT_Value_Volume_Flow.
type DPT_9025 float64

func (d DPT_9025) Pack() []byte {
	if d < -671088.64 {
		return packF16(-671088.64)
	} else if d > 670433.28 {
		return packF16(670433.28)
	} else {
		return packF16(float64(d))
	}
}

func (d *DPT_9025) Unpack(data []byte) error {
	var value float64
	if err := unpackF16(data, &value); err != nil {
		return err
	}

	// Check the value for valid range
	if value < -671088.64 || value > 670433.28 {
		return ErrOutOfRange
	}

	*d = DPT_9025(value)
	return nil
}

func (d DPT_9025) Unit() string {
	return "l/h"
}

func (d DPT_9025) String() string {
	return fmt.Sprintf("%.2f l/h", float64(d))
}

// DPT_9026 represents DPT 9.026 (G)/ DPT_Rain_Amount.
type DPT_9026 float64

func (d DPT_9026) Pack() []byte {
	if d < -671088.64 {
		return packF16(-671088.64)
	} else if d > 670433.28 {
		return packF16(670433.28)
	} else {
		return packF16(float64(d))
	}
}

func (d *DPT_9026) Unpack(data []byte) error {
	var value float64
	if err := unpackF16(data, &value); err != nil {
		return err
	}

	// Check the value for valid range
	if value < -671088.64 || value > 670433.28 {
		return ErrOutOfRange
	}

	*d = DPT_9026(value)
	return nil
}

func (d DPT_9026) Unit() string {
	return "l/m²"
}

func (d DPT_9026) String() string {
	return fmt.Sprintf("%.2f l/m²", float64(d))
}

// DPT_9027 represents DPT 9.027 (G) / DPT_Value_Temp_F.
type DPT_9027 float64

func (d DPT_9027) Pack() []byte {
	if d <= -459.6 {
		return packF16(-459.84)
	} else if d > 670433.28 {
		return packF16(670433.28)
	} else {
		return packF16(float64(d))
	}
}

func (d *DPT_9027) Unpack(data []byte) error {
	var value float64
	if err := unpackF16(data, &value); err != nil {
		return err
	}

	if value < -459.6 {
		value = -459.6
	}
	// Check the value for valid range
	if value < -459.6 || value > 670433.28 {
		return ErrOutOfRange
	}

	*d = DPT_9027(value)
	return nil
}

func (d DPT_9027) Unit() string {
	return "°F"
}

func (d DPT_9027) String() string {
	return fmt.Sprintf("%.2f °F", float64(d))
}

// DPT_9028 represents DPT 9.028 (G) / DPT_Value_Wsp_kmh.
type DPT_9028 float64

func (d DPT_9028) Pack() []byte {
	if d <= 0 {
		return packF16(0)
	} else if d > 670433.28 {
		return packF16(670433.28)
	} else {
		return packF16(float64(d))
	}
}

func (d *DPT_9028) Unpack(data []byte) error {
	var value float64
	if err := unpackF16(data, &value); err != nil {
		return err
	}

	// Check the value for valid range
	if value < 0 || value > 670433.28 {
		return ErrOutOfRange
	}

	*d = DPT_9028(value)
	return nil
}

func (d DPT_9028) Unit() string {
	return "km/h"
}

func (d DPT_9028) String() string {
	return fmt.Sprintf("%.2f km/h", float64(d))
}

// DPT_9029 represents DPT 9.029 (G) / DPT_Value_Absolute_Humidity.
type DPT_9029 float64

func (d DPT_9029) Pack() []byte {
	if d <= 0 {
		return packF16(0)
	} else if d > 670433.28 {
		return packF16(670433.28)
	} else {
		return packF16(float64(d))
	}
}

func (d *DPT_9029) Unpack(data []byte) error {
	var value float64
	if err := unpackF16(data, &value); err != nil {
		return err
	}

	// Check the value for valid range
	if value < 0 || value > 670433.28 {
		return ErrOutOfRange
	}

	*d = DPT_9029(value)
	return nil
}

func (d DPT_9029) Unit() string {
	return "g/m³"
}

func (d DPT_9029) String() string {
	return fmt.Sprintf("%.2f g/m³", float64(d))
}

// DPT_9030 represents DPT 9.030 (G) / DPT_Concentration_μgm3.
type DPT_9030 float64

func (d DPT_9030) Pack() []byte {
	if d <= 0 {
		return packF16(0)
	} else if d > 670433.28 {
		return packF16(670433.28)
	} else {
		return packF16(float64(d))
	}
}

func (d *DPT_9030) Unpack(data []byte) error {
	var value float64
	if err := unpackF16(data, &value); err != nil {
		return err
	}

	// Check the value for valid range
	if value < 0 || value > 670433.28 {
		return ErrOutOfRange
	}

	*d = DPT_9030(value)
	return nil
}

func (d DPT_9030) Unit() string {
	return "μg/m³"
}

func (d DPT_9030) String() string {
	return fmt.Sprintf("%.2f μg/m³", float64(d))
}
