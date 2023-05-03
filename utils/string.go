package utils

import "reflect"

// IsAnyParameterEmpty
// @Description 判断是否存在参数为该类型零值
// @Author John 2023-05-03 21:12:23
// @Param params
// @Return bool
func IsAnyParameterEmpty(params ...interface{}) bool {
	for _, param := range params {
		v := reflect.ValueOf(param)
		switch v.Kind() {
		case reflect.Invalid:
			return true
		case reflect.Slice, reflect.Map, reflect.Chan:
			if v.IsNil() {
				return true
			}
		case reflect.Bool:
			if !v.Bool() {
				return true
			}
		case reflect.String:
			if v.String() == "" {
				return true
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if v.Int() == 0 {
				return true
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			if v.Uint() == 0 {
				return true
			}
		case reflect.Float32, reflect.Float64:
			if v.Float() == 0 {
				return true
			}
		case reflect.Interface, reflect.Ptr:
			if v.IsNil() || IsZeroValue(v.Elem().Interface()) {
				return true
			}
		default:
			if reflect.DeepEqual(reflect.Zero(v.Type()).Interface(), param) {
				return true
			}
		}
	}
	return false
}

func IsZeroValue(i interface{}) bool {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Invalid:
		return true
	case reflect.Slice, reflect.Map, reflect.Chan:
		return v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.String:
		return v.String() == ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil() || IsZeroValue(v.Elem().Interface())
	default:
		return reflect.DeepEqual(reflect.Zero(v.Type()).Interface(), i)
	}
}
