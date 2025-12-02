package dbutil

import (
	"errors"
	"reflect"
	"sync"
	"time"
)

var (
	fieldCache sync.Map // map[srcType][dstType][]fieldSetter
)

type fieldSetter func(src, dst reflect.Value)

func MapStruct(src, dst interface{}) error {
	if src == nil || dst == nil {
		return errors.New("src or dst is nil")
	}

	dv := reflect.ValueOf(dst)
	if dv.Kind() != reflect.Ptr || dv.IsNil() {
		return errors.New("dst must be a non-nil pointer")
	}
	dv = dv.Elem()

	sv := reflect.ValueOf(src)
	if sv.Kind() == reflect.Ptr && !sv.IsNil() {
		sv = sv.Elem()
	}

	if sv.Kind() != reflect.Struct || dv.Kind() != reflect.Struct {
		return errors.New("src and dst must be structs or pointers to structs")
	}

	setters := getSetters(sv.Type(), dv.Type())
	for _, set := range setters {
		set(sv, dv)
	}

	return nil
}

// getSetters returns or builds cached setter functions
func getSetters(srcType, dstType reflect.Type) []fieldSetter {
	key := srcType.String() + "->" + dstType.String()
	if cached, ok := fieldCache.Load(key); ok {
		return cached.([]fieldSetter)
	}

	var setters []fieldSetter
	for i := 0; i < dstType.NumField(); i++ {
		dstField := dstType.Field(i)
		if !dstField.IsExported() {
			continue
		}

		if srcField, ok := srcType.FieldByName(dstField.Name); ok && srcField.IsExported() {
			dstIndex := i
			srcIndex := srcField.Index[0]

			// special handling for time.Time -> int64
			if dstField.Name == "CreatedAt" || dstField.Name == "UpdatedAt" {
				if srcField.Type == reflect.TypeOf(time.Time{}) && dstField.Type.Kind() == reflect.Int64 {
					setters = append(setters, func(sv, dv reflect.Value) {
						t := sv.Field(srcIndex).Interface().(time.Time).Unix()
						dv.Field(dstIndex).SetInt(t)
					})
					continue
				}
			}

			// assignable/convertible types
			setters = append(setters, func(sv, dv reflect.Value) {
				svField := sv.Field(srcIndex)
				dvField := dv.Field(dstIndex)
				if !dvField.CanSet() {
					return
				}
				if svField.Type().AssignableTo(dvField.Type()) {
					dvField.Set(svField)
				} else if svField.Type().ConvertibleTo(dvField.Type()) {
					dvField.Set(svField.Convert(dvField.Type()))
				}
			})
		}
	}

	fieldCache.Store(key, setters)
	return setters
}
