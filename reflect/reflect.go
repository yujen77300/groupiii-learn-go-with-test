package reflect

import (
	"reflect"
)

func walk(x interface{}, fn func(input string)) {

	val := getValue(x)

	if val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
		return
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		// field.Kind() 會回傳一個 reflect.Kind 類型的值，這個值代表了該欄位的型別

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Slice:
			for i := 0; i < val.Len(); i++ {
				walk(val.Index(i).Interface(), fn)
			}
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
