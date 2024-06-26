package reflect

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func DecodeKeyValue() {

	var (
		m0 = make(map[int]string)
		// m1 = make(map[int64]string)
		// m2 map[interface{}]string
		// m3 map[interface{}]interface{}
		// m4 map[bool]string
	)
	s := "1=foo,2=bar,3=baz"
	fmt.Println(decodeMap(s, &m0), m0) // pass
	// fmt.Println(decodeMap(s, &m1), m1) // pass
	// fmt.Println(decodeMap(s, &m2), m2) // pass
	// fmt.Println(decodeMap(s, &m3), m3) // pass
	// fmt.Println(decodeMap(s, &m4), m4) // fail

}

func decodeMap(s string, i interface{}) error {
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Pointer {
		return fmt.Errorf("non pointer %v", v.Type())
	}

	v = v.Elem()
	t := v.Type()

	if v.IsNil() {
		v.Set(reflect.MakeMap(t))
	}

	// fmt.Println(v.Addr().Elem(), v.Addr().Type())

	for _, kv := range strings.Split(s, ",") {
		s := strings.Split(kv, "=")
		n, err := strconv.Atoi(s[0])
		if err != nil {
			return fmt.Errorf("failed to parse number: %v", err)
		}
		k, e := reflect.ValueOf(n), reflect.ValueOf(s[1])
		// get the type of the key.
		kt := t.Key()
		if !k.Type().ConvertibleTo(kt) {
			return fmt.Errorf("can't convert key to type %v", kt.Kind())
		}
		k = k.Convert(kt)
		// get the element type.
		et := t.Elem()
		if et.Kind() != v.Kind() && !e.Type().ConvertibleTo(et) {
			return fmt.Errorf("can't assign value to type %v", kt.Kind())
		}
		v.SetMapIndex(k, e.Convert(et))

	}

	return nil
}
