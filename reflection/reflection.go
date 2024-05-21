package reflection

import "reflect"

// call fn on each attribute of type string in x
func walk(x interface{}, fn func(string)) {
	val := reflect.ValueOf(x)

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
