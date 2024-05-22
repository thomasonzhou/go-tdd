package reflection

import "reflect"

// call fn on each attribute of type string in x
func walk(x interface{}, fn func(string)) {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	if val.Kind() == reflect.String {
		fn(x.(string))
	} else if val.Kind() == reflect.Struct {
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)

			switch field.Kind() {
			case reflect.String:
				fn(field.String())
			case reflect.Struct:
				walk(field.Interface(), fn)
			}
		}
	} else if val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
	}
}
