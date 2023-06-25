package util

import (
	"fmt"
	"reflect"
	"unsafe"
)

func GetPtrUnExportFiled(s interface{}, field string) reflect.Value {
	v := reflect.ValueOf(s).Elem().FieldByName(field)
	// 必须要调用 Elem()
	return reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr())).Elem()
}

func SetPtrUnExportFiled(s interface{}, field string, val interface{}) error {
	v := GetPtrUnExportFiled(s, field)
	rv := reflect.ValueOf(val)
	if v.Kind() != v.Kind() {
		return fmt.Errorf("invalid kind, expected kind: %v, got kind:%v", v.Kind(), rv.Kind())
	}

	v.Set(rv)
	return nil
}
