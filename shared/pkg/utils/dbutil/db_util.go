package dbutil

import (
	"errors"
	"reflect"
	"sync"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	fieldCache sync.Map // map[string][]fieldSetter
)

type fieldSetter func(src, dst reflect.Value)

func MapStruct(src, dst interface{}) error {
	if src == nil || dst == nil {
		return errors.New("src or dst is nil")
	}

	sv := reflect.ValueOf(src)
	dv := reflect.ValueOf(dst)

	if dv.Kind() != reflect.Ptr || dv.IsNil() {
		return errors.New("dst must be a non-nil pointer")
	}

	sv = indirect(sv)
	dv = indirect(dv)

	if sv.Kind() != reflect.Struct || dv.Kind() != reflect.Struct {
		return errors.New("src and dst must be struct or *struct")
	}

	setters := getSetters(sv.Type(), dv.Type())
	for _, set := range setters {
		set(sv, dv)
	}

	return nil
}

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

		srcField, ok := srcType.FieldByName(dstField.Name)
		if !ok || !srcField.IsExported() {
			continue
		}

		dstIndex := dstField.Index
		srcIndex := srcField.Index

		srcFieldType := srcField.Type
		dstFieldType := dstField.Type

		// ----------------------------------------------------------
		// Special case: time.Time <-> timestamppb.Timestamp
		// ----------------------------------------------------------
		if (dstField.Name == "CreatedAt" || dstField.Name == "UpdatedAt") &&
			srcFieldType == reflect.TypeOf(time.Time{}) &&
			dstFieldType == reflect.TypeOf(&timestamppb.Timestamp{}) {

			setters = append(setters, func(sv, dv reflect.Value) {
				svField := fieldByIndex(sv, srcIndex)
				if !svField.IsValid() {
					return
				}
				ts := timestamppb.New(svField.Interface().(time.Time))
				fieldByIndex(dv, dstIndex).Set(reflect.ValueOf(ts))
			})
			continue
		}

		if (dstField.Name == "CreatedAt" || dstField.Name == "UpdatedAt") &&
			srcFieldType == reflect.TypeOf(&timestamppb.Timestamp{}) &&
			dstFieldType == reflect.TypeOf(time.Time{}) {

			setters = append(setters, func(sv, dv reflect.Value) {
				svField := fieldByIndex(sv, srcIndex)
				if !svField.IsValid() {
					return
				}
				t := svField.Interface().(*timestamppb.Timestamp).AsTime()
				fieldByIndex(dv, dstIndex).Set(reflect.ValueOf(t))
			})
			continue
		}

		// ----------------------------------------------------------
		// Nested struct or pointer-to-struct → recursively map
		// ----------------------------------------------------------
		if isStructOrPtrStruct(srcFieldType) &&
			isStructOrPtrStruct(dstFieldType) {

			nestedSrc := indirectType(srcFieldType)
			nestedDst := indirectType(dstFieldType)

			nestedSet := getSetters(nestedSrc, nestedDst)

			setters = append(setters, func(sv, dv reflect.Value) {
				svField := fieldByIndex(sv, srcIndex)
				if !svField.IsValid() {
					return
				}

				svField = indirect(svField)
				if svField.Kind() != reflect.Struct {
					return
				}

				dvField := fieldByIndex(dv, dstIndex)

				// allocate pointer target if needed
				if dvField.Kind() == reflect.Ptr {
					if dvField.IsNil() {
						dvField.Set(reflect.New(dvField.Type().Elem()))
					}
					dvField = dvField.Elem()
				}

				if dvField.Kind() != reflect.Struct {
					return
				}

				for _, n := range nestedSet {
					n(svField, dvField)
				}
			})

			continue
		}

		// ----------------------------------------------------------
		// Normal assign / convert
		// ----------------------------------------------------------
		setters = append(setters, func(sv, dv reflect.Value) {
			svField := fieldByIndex(sv, srcIndex)
			if !svField.IsValid() {
				return
			}

			svField = indirect(svField)

			dvField := fieldByIndex(dv, dstIndex)
			if !dvField.CanSet() {
				return
			}

			if dvField.Kind() == reflect.Ptr {
				if dvField.IsNil() {
					dvField.Set(reflect.New(dvField.Type().Elem()))
				}
				dvField = dvField.Elem()
			}

			if svField.Type().AssignableTo(dvField.Type()) {
				dvField.Set(svField)
			} else if svField.Type().ConvertibleTo(dvField.Type()) {
				dvField.Set(svField.Convert(dvField.Type()))
			}
		})
	}

	fieldCache.Store(key, setters)
	return setters
}

// ----------------------------------------------------------
// Helpers
// ----------------------------------------------------------

func indirect(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return v
		}
		v = v.Elem()
	}
	return v
}

func indirectType(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

func fieldByIndex(v reflect.Value, index []int) reflect.Value {
	for _, i := range index {
		if v.Kind() == reflect.Ptr {
			if v.IsNil() {
				return reflect.Value{}
			}
			v = v.Elem()
		}
		v = v.Field(i)
	}
	return v
}

func isStructOrPtrStruct(t reflect.Type) bool {
	if t.Kind() == reflect.Struct {
		return true
	}
	return t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
}
