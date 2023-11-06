package util

import (
	"encoding/json"
	"reflect"
	"runtime"
)

func WhoAmI(skip int) string {
	pc, _, _, ok := runtime.Caller(skip + 1)
	if !ok {
		return "unknown"
	}

	me := runtime.FuncForPC(pc)
	if me == nil {
		return "unnamed"
	}

	return me.Name()
}

func IsJSONBody(x any) bool {
	v := reflect.ValueOf(x)
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Pointer:
		return v.IsNil()
	}
	return false
}

func getBodyString(body any) string {
	var responseBytes []byte

	value := reflect.ValueOf(body)
	switch value.Kind() {
	case reflect.String:
		responseBytes = []byte(body.(string))
	case reflect.Array, reflect.Slice:
		if value.Type().Elem().Kind() == reflect.Uint8 {
			responseBytes = value.Bytes()
		} else {
			responseBytes, _ = json.Marshal(body)
		}
	case reflect.Pointer:
		if value.Type().Elem().Kind() == reflect.String {
			strPtr := value.Interface().(*string)
			if strPtr != nil {
				responseBytes = []byte(*strPtr)
				break
			}
		}
		fallthrough
	default:
		if body != nil {
			responseBytes, _ = json.Marshal(body)
		}
	}

	return string(responseBytes)
}
