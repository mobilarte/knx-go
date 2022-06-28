// Copyright 2017 Ole Krüger.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import (
	"encoding/binary"
	"errors"
	"math"
)

var (
	// ErrInvalidLength is returned when the application data has unexpected length.
	ErrInvalidLength = errors.New("application data has invalid length")
	// ErrInvalidData is returned when a bit or content denotes invalid data per KNX specification.
	ErrInvalidData = errors.New("data is noted as invalid")
	// ErrOutOfRange is returned when the unpacked value is out of range.
	ErrOutOfRange = errors.New("payload is out of range")
	// ErrBadReservedBits is returned when reserved bits are populated. E.g. if bit number 5 of a r4B4 field is populated.
	ErrBadReservedBits = errors.New("reserved bits in the input data have been populated")
)

func packB1(b bool) []byte {
	if b {
		return []byte{1}
	}

	return []byte{0}
}

func unpackB1(data []byte, b *bool) error {
	if len(data) != 1 {
		return ErrInvalidLength
	}

	*b = data[0]&1 == 1

	return nil
}

func packB4(bs [4]bool) byte {
	var b byte = 0
	if bs[0] {
		b |= 1 << 0
	}
	if bs[1] {
		b |= 1 << 1
	}
	if bs[2] {
		b |= 1 << 2
	}
	if bs[3] {
		b |= 1 << 3
	}

	return byte(b)
}

func unpackB4(data byte, b0 *bool, b1 *bool, b2 *bool, b3 *bool) error {

	if uint8(data) > 15 {
		return ErrBadReservedBits
	}

	*b0 = ((data >> 0) & 1) != 0
	*b1 = ((data >> 1) & 1) != 0
	*b2 = ((data >> 2) & 1) != 0
	*b3 = ((data >> 3) & 1) != 0

	return nil
}

func packU8(i uint8) []byte {
	return []byte{0, i}
}

func unpackU8(data []byte, i *uint8) error {
	if len(data) != 2 {
		return ErrInvalidLength
	}

	*i = uint8(data[1])

	return nil
}

func packV8(i int8) []byte {
	return []byte{0, byte(i)}
}

func unpackV8(data []byte, i *int8) error {
	if len(data) != 2 {
		return ErrInvalidLength
	}

	*i = int8(data[1])

	return nil
}

func packU16(i uint16) []byte {
	buffer := make([]byte, 3)

	binary.BigEndian.PutUint16(buffer[1:], i)

	return buffer
}

func unpackU16(data []byte, i *uint16) error {
	if len(data) != 3 {
		return ErrInvalidLength
	}

	*i = binary.BigEndian.Uint16(data[1:])

	return nil
}

func packV16(i int16) []byte {
	buffer := make([]byte, 3)

	buffer[0] = 0
	buffer[1] = byte((i >> 8) & 0xff)
	buffer[2] = byte(i & 0xff)

	return buffer
}

func unpackV16(data []byte, i *int16) error {
	if len(data) != 3 {
		return ErrInvalidLength
	}

	*i = int16(data[1])<<8 | int16(data[2])

	return nil
}

func packU32(i uint32) []byte {
	buffer := make([]byte, 5)

	binary.BigEndian.PutUint32(buffer[1:], i)

	return buffer
}

func unpackU32(data []byte, i *uint32) error {
	if len(data) != 5 {
		return ErrInvalidLength
	}

	*i = binary.BigEndian.Uint32(data[1:])

	return nil
}

func packV32(i int32) []byte {
	buffer := make([]byte, 5)

	buffer[0] = 0
	buffer[1] = byte((i >> 24) & 0xff)
	buffer[2] = byte((i >> 16) & 0xff)
	buffer[3] = byte((i >> 8) & 0xff)
	buffer[4] = byte(i & 0xff)

	return buffer
}

func unpackV32(data []byte, i *int32) error {
	if len(data) != 5 {
		return ErrInvalidLength
	}

	*i = int32(data[1])<<24 | int32(data[2])<<16 | int32(data[3])<<8 | int32(data[4])

	return nil
}

func packF16(f float64) []byte {
	buffer := make([]byte, 3)

	if f > 670433.28 {
		f = 670433.28
	} else if f < -671088.64 {
		f = -671088.64
	}

	signedMantissa := f * 100.0
	exp := 0

	for signedMantissa > 2048 || signedMantissa < -2048 {
		signedMantissa /= 2
		exp++
	}

	buffer[1] |= uint8(exp&0xF) << 3

	if signedMantissa < 0 {
		signedMantissa += 2048
		buffer[1] |= 0x1 << 7
	}

	mantissa := uint(signedMantissa)

	buffer[1] |= uint8(mantissa>>8) & 0x7
	buffer[2] |= uint8(mantissa)

	return buffer
}

func unpackF16(data []byte, f *float64) error {
	if len(data) != 3 {
		return ErrInvalidLength
	}

	// This value denotes invalid data, but only applies to DPT 9.00x
	if data[1]&0x7F == 0x7F && data[2]&0xFF == 0xFF {
		return ErrInvalidData
	}

	m := int(data[1]&0x7)<<8 | int(data[2])
	e := (data[1] >> 3) & 0x0F
	if data[1]&0x80 == 0x80 {
		m -= 2048
	}
	value := float64(m<<e) / 100
	*f = value

	return nil
}

func packF32(f float32) []byte {
	buffer := make([]byte, 5)

	binary.BigEndian.PutUint32(buffer[1:], math.Float32bits(f))

	return buffer
}

func unpackF32(data []byte, f *float32) error {
	if len(data) != 5 {
		return ErrInvalidLength
	}

	*f = math.Float32frombits(binary.BigEndian.Uint32(data[1:]))

	return nil
}
