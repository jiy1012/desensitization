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
		return processStruct(rv)
	case reflect.Slice, reflect.Array:
		return processSlice(rv, "")
	case reflect.Interface:
		if !rv.IsNil() {
			return Desensitization(rv.Elem().Interface())
		}
	default:
		return nil
	}
	return nil
}

func processStruct(rv reflect.Value) error {
	rt := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		fieldType := rt.Field(i)

		if !field.CanSet() {
			continue
		}

		tag := fieldType.Tag.Get(TagKey)
		switch field.Kind() {
		case reflect.Struct:
			if err := processStruct(field); err != nil {
				return err
			}
		case reflect.Slice, reflect.Array:
			if err := processSlice(field, tag); err != nil {
				return err
			}
		case reflect.Interface:
			if !field.IsNil() {
				if err := Desensitization(field.Interface()); err != nil {
					return err
				}
			}
		default:
			if tag != "" {
				if err := processField(field, tag); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func processSlice(field reflect.Value, tag string) error {
	for i := 0; i < field.Len(); i++ {
		elem := field.Index(i)
		switch elem.Kind() {
		case reflect.Ptr:
			if !elem.IsNil() {
				if err := Desensitization(elem.Interface()); err != nil {
					return err
				}
			}
		case reflect.Struct:
			if err := processStruct(elem); err != nil {
				return err
			}
		case reflect.Interface:
			if !elem.IsNil() {
				if err := Desensitization(elem.Interface()); err != nil {
					return err
				}
			}
		default:
			if tag != "" {
				if err := processField(elem, tag); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func processField(field reflect.Value, tag string) error {
	newVal, err := OperateByRule(tag, field.Interface())
	if err != nil {
		return err
	}
	nv := reflect.ValueOf(newVal)
	if nv.Type().ConvertibleTo(field.Type()) {
		field.Set(nv.Convert(field.Type()))
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
