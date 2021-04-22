package value

import (
	"errors"
	"time"
)

func FromBool(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}
func FromByte(b *byte) byte {
	if b == nil {
		return byte(0)
	}
	return *b
}
func FromComplex128(c *complex128) complex128 {
	if c == nil {
		return complex128(0)
	}
	return *c
}
func FromComplex64(c *complex64) complex64 {
	if c == nil {
		return complex64(0)
	}
	return *c
}
func FromDuration(d *time.Duration) time.Duration {
	if d == nil {
		return time.Duration(0)
	}
	return *d
}
func FromError(e *error) error {
	if e == nil {
		return errors.New("")
	}
	return *e
}
func FromFloat32(f *float32) float32 {
	if f == nil {
		return float32(0)
	}
	return *f
}
func FromFloat64(f *float64) float64 {
	if f == nil {
		return float64(0)
	}
	return *f
}
func FromInt(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}
func FromInt16(i *int16) int16 {
	if i == nil {
		return 0
	}
	return *i
}
func FromInt32(i *int32) int32 {
	if i == nil {
		return 0
	}
	return *i
}
func FromInt64(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}
func FromInt8(i *int8) int8 {
	if i == nil {
		return 0
	}
	return *i
}
func FromRune(r *rune) rune {
	if r == nil {
		return rune(0)
	}
	return *r
}
func FromString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
func FromTime(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	}
	return *t
}
func FromUint(u *uint) uint {
	if u == nil {
		return 0
	}
	return *u
}
func FromUint16(u *uint16) uint16 {
	if u == nil {
		return 0
	}
	return *u
}
func FromUint32(u *uint32) uint32 {
	if u == nil {
		return 0
	}
	return *u
}
func FromUint64(u *uint64) uint64 {
	if u == nil {
		return 0
	}
	return *u
}
func FromUint8(u *uint8) uint8 {
	if u == nil {
		return 0
	}
	return *u
}
func FromUintptr(u *uintptr) uintptr {
	if u == nil {
		return 0
	}
	return *u
}
