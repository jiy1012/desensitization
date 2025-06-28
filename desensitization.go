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

	rt := rv.Type()
	switch rv.Kind() {
	case reflect.Struct:
		for i := 0; i < rv.NumField(); i++ {
			field := rv.Field(i)
			fieldType := rt.Field(i)
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
				elemType := fieldType.Type.Elem()
				isPtr := elemType.Kind() == reflect.Ptr
				if isPtr {
					elemType = elemType.Elem()
				}

				if elemType.Kind() == reflect.Struct {
					for j := 0; j < field.Len(); j++ {
						elem := field.Index(j)
						if isPtr {
							if !elem.IsNil() {
								_ = Desensitization(elem.Interface())
							}
						} else {
							if elem.CanAddr() {
								_ = Desensitization(elem.Addr().Interface())
							} else {
								_ = Desensitization(elem.Interface())
							}
						}
					}
				} else if tag != "" {
					for j := 0; j < field.Len(); j++ {
						elem := field.Index(j)
						newVal, err := OperateByRule(tag, elem.Interface())
						if err == nil {
							nv := reflect.ValueOf(newVal)
							if nv.Type().ConvertibleTo(elem.Type()) {
								elem.Set(nv.Convert(elem.Type()))
							}
						}
					}
				}
			default:
				if tag != "" {
					newVal, err := OperateByRule(tag, field.Interface())
					if err == nil {
						nv := reflect.ValueOf(newVal)
						if nv.Type().ConvertibleTo(field.Type()) {
							field.Set(nv.Convert(field.Type()))
						}
					}
				}
			}
		}
	}
	return nil
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
