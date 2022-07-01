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
	ErrInvalidData = errors.New("application data is noted as invalid")
	// ErrOutOfRange is returned when the unpacked value is out of range.
	ErrOutOfRange = errors.New("application data is out of range")
	// ErrBadReservedBits is returned when reserved bits are populated. E.g. if bit number 5 of a r4B4 field is populated.
	ErrBadReservedBits = errors.New("reserved bits in the application data have been populated")
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

	*b = data[0]&0x1 == 0x1

	return nil
}

func packB2(b1 bool, b0 bool) []byte {
	var b byte

	if b1 {
		b |= 0x2
	}
	if b0 {
		b |= 0x1
	}

	return []byte{b}
}

func unpackB2(data []byte, b1 *bool, b0 *bool) error {

	if data[0] > 0x3 {
		return ErrBadReservedBits
	}

	*b1 = data[0]&0x2 == 0x2
	*b0 = data[0]&0x1 == 0x1

	return nil
}

func packB4(b3 bool, b2 bool, b1 bool, b0 bool) byte {
	var b byte

	if b3 {
		b |= 0x8
	}
	if b2 {
		b |= 0x4
	}
	if b1 {
		b |= 0x2
	}
	if b0 {
		b |= 0x1
	}

	return byte(b)
}

func unpackB4(data byte, b3 *bool, b2 *bool, b1 *bool, b0 *bool) error {

	if uint8(data) > 15 {
		return ErrBadReservedBits
	}

	*b3 = data&0x8 == 0x8
	*b2 = data&0x4 == 0x4
	*b1 = data&0x2 == 0x2
	*b0 = data&0x1 == 0x1

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

	binary.BigEndian.PutUint16(buffer[1:], uint16(i))

	return buffer
}

func unpackV16(data []byte, i *int16) error {
	if len(data) != 3 {
		return ErrInvalidLength
	}

	*i = int16(binary.BigEndian.Uint16(data[1:]))

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

	binary.BigEndian.PutUint32(buffer[1:], uint32(i))

	return buffer
}

func unpackV32(data []byte, i *int32) error {
	if len(data) != 5 {
		return ErrInvalidLength
	}

	*i = int32(binary.BigEndian.Uint32(data[1:]))

	return nil
}

func packV64(i int64) []byte {
	buffer := make([]byte, 9)

	binary.BigEndian.PutUint64(buffer[1:], uint64(i))

	return buffer
}

func unpackV64(data []byte, i *int64) error {
	if len(data) != 9 {
		return ErrInvalidLength
	}

	*i = int64(binary.BigEndian.Uint64(data[1:]))

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

	for signedMantissa > 2047 || signedMantissa < -2048 {
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

	// This value denotes invalid data. It only applies to DPT 9.00x!
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
