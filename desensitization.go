package desensitization

import (
	"reflect"
)

func Desensitization(obj interface{}) error {
	rv := reflect.ValueOf(obj)
	if !rv.IsValid() {
		return nil
	}

	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return nil
		}
		rv = rv.Elem()
	}

	if isBasicType(rv.Kind()) {
		return nil
	}

	switch rv.Kind() {
	case reflect.Struct:
		processStruct(rv)
	case reflect.Slice, reflect.Array:
		processSlice(rv, "")
	}
	return nil
}

func processStruct(rv reflect.Value) {
	rt := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		fieldType := rt.Field(i) // 正确获取StructField

		if !field.CanSet() {
			continue
		}

		tag := fieldType.Tag.Get(TagKey)
		switch field.Kind() {
		case reflect.Struct:
			if field.CanAddr() {
				_ = Desensitization(field.Addr().Interface())
			} else {
				_ = Desensitization(field.Interface())
			}
		case reflect.Slice, reflect.Array:
			processSlice(field, tag)
		default:
			if tag != "" {
				processField(field, tag)
			}
		}
	}
}

func processSlice(field reflect.Value, tag string) {
	for i := 0; i < field.Len(); i++ {
		elem := field.Index(i)
		if elem.Kind() == reflect.Ptr {
			if !elem.IsNil() {
				_ = Desensitization(elem.Interface())
			}
		} else if elem.Kind() == reflect.Struct {
			if elem.CanAddr() {
				_ = Desensitization(elem.Addr().Interface())
			} else {
				_ = Desensitization(elem.Interface())
			}
		} else if tag != "" {
			processField(elem, tag)
		}
	}
}

func processField(field reflect.Value, tag string) {
	newVal, err := OperateByRule(tag, field.Interface())
	if err == nil {
		nv := reflect.ValueOf(newVal)
		if nv.Type().ConvertibleTo(field.Type()) {
			field.Set(nv.Convert(field.Type()))
		}
	}
}

func isBasicType(kind reflect.Kind) bool {
	switch kind {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128, reflect.String:
		return true
	default:
		return false
	}
}
