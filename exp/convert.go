package exp

import (
	"fmt"
	"reflect"
)

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
	return ReflectConvert[[]byte](a)
}

func ToRunes(a any) (s []rune) {
	switch i := a.(type) {
	case string:
		return []rune(i)
	case []rune:
		return []rune(i)
	}
	return ReflectConvert[[]rune](a)
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
	return ReflectConvert[T](a)
}

func ReflectConvert[T any](from any) (to T) {
	rt := reflect.TypeOf(to)
	rv := reflect.ValueOf(from)
	if rv.CanConvert(rt) {
		return rv.Convert(rt).Interface().(T)
	}
	panic(fmt.Sprintf("cannot convert %T(%v) to %T", from, from, to))
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
