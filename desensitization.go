package desensitization

import (
	"reflect"
)

func Desensitization(obj interface{}) error {
	rt := reflect.TypeOf(obj)
	rv := reflect.ValueOf(obj)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
		rt = rt.Elem()
	}
	switch rv.Kind() {
	case reflect.Struct:
		for idx := 0; idx < rv.NumField(); idx++ {
			fieldValue := rv.Field(idx)
			fieldType := rt.Field(idx)
			switch fieldType.Type.Kind() {
			case reflect.Slice, reflect.Array:
				// 判断字段类型是否为结构体数组
				if fieldType.Type.Elem().Kind() == reflect.Struct {
					_ = Desensitization(fieldValue.Addr().Interface())
				} else {
					desensitizationTag := fieldType.Tag.Get(Type)
					if desensitizationTag != "" {
						for i := 0; i < fieldValue.Len(); i++ {
							elemValue := fieldValue.Index(i)
							newValue, err := OperateByRule(desensitizationTag, elemValue.Interface())
							if err == nil {
								elemValue.Set(reflect.ValueOf(newValue))
							}
						}
					}
				}
			case reflect.Struct:
				_ = Desensitization(fieldValue.Addr().Interface())
			default:
				desensitizationTag := fieldType.Tag.Get(Type)
				if desensitizationTag != "" {
					newValue, err := OperateByRule(desensitizationTag, fieldValue.Interface())
					if err == nil {
						fieldValue.Set(reflect.ValueOf(newValue))
					}
				}
			}
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < rv.Len(); i++ {
			elemValue := rv.Index(i)
			if reflect.TypeOf(elemValue).Kind() == reflect.Ptr {
				_ = Desensitization(elemValue.Interface())
			} else {
				_ = Desensitization(elemValue.Addr().Interface())
			}
		}
	default:
	}
	return nil
}
