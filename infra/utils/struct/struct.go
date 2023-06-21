package _struct

import (
	"encoding/json"
	"errors"
	"reflect"
)

// CopyStruct 将src的值赋值到dst
func CopyStruct(dst interface{}, src interface{}) (err error) {

	if dst == nil || src == nil {
		err = errors.New("CopyStruct: dst and src must not be nil")
		return
	}

	dstVal := reflect.ValueOf(dst)

	if dstVal.Kind() != reflect.Ptr {
		err = errors.New("CopyStruct: dst reflect kind need to be a pointer")
		return
	}

	if dstVal.IsNil() {
		err = errors.New("CopyStruct: dstVal.IsNil()")
		return
	}

	bys, err := json.Marshal(src)
	if err != nil {
		return
	}

	err = json.Unmarshal(bys, dst)

	return
}
