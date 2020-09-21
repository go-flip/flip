// +build !go1.8

package flip

import "reflect"

var reflectValueOf = reflect.ValueOf

func reflectSwapper(x interface{}) func(int, int) {
	v := reflectValueOf(x)
	tmp := reflect.New(v.Type().Elem()).Elem()
	return func(i, j int) {
		a, b := v.Index(i), v.Index(j)
		tmp.Set(a)
		a.Set(b)
		b.Set(tmp)
	}
}
