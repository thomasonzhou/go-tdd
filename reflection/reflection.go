package reflection

import "reflect"

// call fn on each attribute of type string in x
func walk(x interface{}, fn func(string)) {
	value := getValue(x)

	valueCount := 0
	var getField func(int) reflect.Value

	switch value.Kind() {
	case reflect.String:
		fn(x.(string))
	case reflect.Struct:
		valueCount = value.NumField()
		getField = value.Field
	case reflect.Slice, reflect.Array:
		valueCount = value.Len()
		getField = value.Index
	case reflect.Map:
		for _, key := range value.MapKeys() {
			walk(value.MapIndex(key).Interface(), fn)
		}
	}
	for i := 0; i < valueCount; i++ {
		walk(getField(i).Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	return val
}
