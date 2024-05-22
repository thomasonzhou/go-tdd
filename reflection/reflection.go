package reflection

import "reflect"

// call fn on each attribute of type string in x
func walk(x interface{}, fn func(string)) {
	value := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch value.Kind() {
	case reflect.String:
		fn(x.(string))
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			walkValue(value.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			walkValue(value.Index(i))
		}
	case reflect.Map:
		for _, key := range value.MapKeys() {
			walkValue(value.MapIndex(key))
		}
	case reflect.Chan:
		for {
			if v, ok := value.Recv(); ok {
				walkValue(v)
			} else {
				break
			}
		}
	case reflect.Func:
		fnCallRes := value.Call(nil)
		for _, res := range fnCallRes {
			walkValue(res)
		}
	}

}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	return val
}
