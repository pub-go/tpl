package exp

import "fmt"

func ToString(a any) string {
	switch i := a.(type) {
	case string:
		return string(i)
	case []byte:
		return string(i)
	case []rune:
		return string(i)
	case fmt.Stringer:
		return i.String()
	default:
		return fmt.Sprintf("%v", i)
	}
}

func ToBytes(a any) (s []byte) {
	switch i := a.(type) {
	case string:
		return []byte(i)
	case []byte:
		return []byte(i)
	}
	panic(fmt.Sprintf("cannot convert %T(%v) to []byte", a, a))
}

func ToRunes(a any) (s []rune) {
	switch i := a.(type) {
	case string:
		return []rune(i)
	case []rune:
		return []rune(i)
	}
	panic(fmt.Sprintf("cannot convert %T(%v) to []rune", a, a))
}

func ToNumber[T Integer | Float](a any) (n T) {
	switch i := a.(type) {
	case uint:
		return T(i)
	case uint8:
		return T(i)
	case uint16:
		return T(i)
	case uint32:
		return T(i)
	case uint64:
		return T(i)
	case int:
		return T(i)
	case int8:
		return T(i)
	case int16:
		return T(i)
	case int32:
		return T(i)
	case int64:
		return T(i)
	case float32:
		return T(i)
	case float64:
		return T(i)
	}
	panic(fmt.Sprintf("cannot convert %T(%v) to %T", a, a, n))
}

type (
	Signed interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64
	}
	Unsigned interface {
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
	}
	Integer interface {
		Signed | Unsigned
	}
	Float interface {
		~float32 | ~float64
	}
	Complex interface {
		~complex64 | ~complex128
	}
	Ordered interface {
		Integer | Float | ~string
	}
)
