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
		return
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
	}
}
