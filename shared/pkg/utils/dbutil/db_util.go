package dbutil

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"time"
)

// fieldCache stores mapping from dst type to slice of src/dst field indices
var (
	fieldCache sync.Map // map[reflect.Type][]fieldPair
)

type fieldPair struct {
	srcIndex int
	dstIndex int
}

// MapStruct copies exported fields from src to dst efficiently.
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

	dstType := dv.Type()
	pairs := getFieldPairs(sv.Type(), dstType)

	for _, pair := range pairs {
		sField := sv.Field(pair.srcIndex)
		dField := dv.Field(pair.dstIndex)

		if dField.CanSet() {
			fieldName := dstType.Field(pair.dstIndex).Name
			if err := setValue(fieldName, sField, dField); err != nil {
				return fmt.Errorf("cannot copy field %s: %w", fieldName, err)
			}
		}
	}

	return nil
}

// getFieldPairs returns cached src->dst field index mapping
func getFieldPairs(srcType, dstType reflect.Type) []fieldPair {
	if cached, ok := fieldCache.Load(dstType); ok {
		return cached.([]fieldPair)
	}

	var pairs []fieldPair
	for i := 0; i < dstType.NumField(); i++ {
		dstField := dstType.Field(i)
		if !dstField.IsExported() {
			continue
		}
		if srcField, ok := srcType.FieldByName(dstField.Name); ok && srcField.IsExported() {
			pairs = append(pairs, fieldPair{srcIndex: srcField.Index[0], dstIndex: i})
		}
	}

	fieldCache.Store(dstType, pairs)
	return pairs
}

// setValue handles pointers, structs, and basic types recursively
func setValue(fieldName string, sv, dv reflect.Value) error {
	if !sv.IsValid() || !dv.CanSet() {
		return nil
	}

	// Special cases
	switch fieldName {
	case "CreatedAt", "UpdatedAt":
		if sv.Type() == reflect.TypeOf(time.Time{}) {
			ts := sv.Interface().(time.Time).Unix()
			dv.Set(reflect.ValueOf(ts))
			return nil
		}
	}

	switch sv.Kind() {
	case reflect.Ptr:
		if sv.IsNil() {
			return nil
		}
		if dv.Kind() == reflect.Ptr {
			if dv.IsNil() {
				dv.Set(reflect.New(dv.Type().Elem()))
			}
			return setValue(fieldName, sv.Elem(), dv.Elem())
		}
		return setValue(fieldName, sv.Elem(), dv)

	case reflect.Struct:
		if dv.Kind() != reflect.Struct {
			return fmt.Errorf("cannot assign struct to %s", dv.Type())
		}
		for i := 0; i < dv.NumField(); i++ {
			df := dv.Field(i)
			if df.CanSet() {
				sf := sv.FieldByName(dv.Type().Field(i).Name)
				setValue(fieldName, sf, df)
			}
		}
		return nil
	}

	// Assignable or convertible types
	if sv.Type().AssignableTo(dv.Type()) {
		dv.Set(sv)
		return nil
	}
	if sv.Type().ConvertibleTo(dv.Type()) {
		dv.Set(sv.Convert(dv.Type()))
		return nil
	}

	return fmt.Errorf("type mismatch: src=%s dst=%s", sv.Type(), dv.Type())
}
