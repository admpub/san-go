package san

import (
	"fmt"
	"reflect"
)

// format error util
func e(format string, args ...interface{}) error {
	return fmt.Errorf("san: "+format, args...)
}

// Unmarshal decodes the contents of `p` in SAN format into a pointer `v`.
func Unmarshal(p []byte, v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		return e("Decode of non-pointer %s", reflect.TypeOf(v))
	}
	if rv.IsNil() {
		return e("Decode of nil %s", reflect.TypeOf(v))
	}
	return nil
}
