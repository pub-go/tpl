package exp

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"

	"code.gopub.tech/errors"
)

var ErrNoSuchValue = fmt.Errorf("no such value")

// NameOfFunction 获取函数的全限定名
func NameOfFunction(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

// getValue 从 from 中获取 name 对应的值
func getValue(name string, from any) (value any, err error) {
	defer func() {
		if x := recover(); x != nil {
			err = errors.Errorf("cannot get `%v` from `%T`: %v", name, from, x)
		}
	}()

	if from == nil {
		return nil, errors.Wrapf(ErrNoSuchValue, "can not get `%s` from nil", name)
	}

	// type definition 用“类型定义”定义的新类型上 可以定义方法
	// type time.Duration int64
	// func (d Duration) Nanoseconds() int64 { return int64(d) }
	rv := reflect.ValueOf(from)
	method := rv.MethodByName(name)
	if !isZero(method) {
		return method.Interface(), nil
	}

	if rv.Kind() == reflect.Pointer {
		rv = rv.Elem()
	}
	switch rv.Kind() {
	case reflect.Struct:
		value := rv.FieldByName(name)
		if isZero(value) { // 无字段
			return nil, errors.Errorf("field not found `%v` in %v: %w", name, rv.Type(), ErrNoSuchValue)
		}
		return value.Interface(), nil
	case reflect.Map:
		value := rv.MapIndex(reflect.ValueOf(name))
		if isZero(value) {
			return nil, errors.Errorf("key not found `%v`: %w", name, ErrNoSuchValue)
		}
		return value.Interface(), nil
	case reflect.Array, reflect.Slice:
		idx, err := strconv.ParseInt(name, 10, 64)
		if err != nil {
			return nil, errors.Errorf("expected array index, but got `%s`: %w", name, err)
		}
		if idx < 0 {
			idx = int64(rv.Len()) + idx
		}
		return rv.Index(int(idx)).Interface(), nil
	default:
	}
	return nil, errors.Errorf("`%s` not found in `%T`: %w", name, from, ErrNoSuchValue)
}

// callFunc 反射调用函数
func callFunc(rv reflect.Value, in []reflect.Value) (out []reflect.Value, err error) {
	defer func() {
		if x := recover(); x != nil {
			// reflect: Call with too many input arguments
			err = errors.Errorf("call function failed: %v", x)
		}
	}()
	return rv.Call(in), nil
}

// isZero 检查值是否为零值
func isZero(value reflect.Value) bool {
	return value == reflect.Value{}
}
