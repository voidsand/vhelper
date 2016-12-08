package vhelper

import (
	"errors"
	"reflect"
)

func StructMethodCall(o ...interface{}) (interface{}, error) {
	var nom string
	var ok bool
	args := make([]reflect.Value, 0)
	rtns := make([]interface{}, 0)
	v := reflect.ValueOf(o[0])
	if v.Kind() != reflect.Ptr {
		error := errors.New("第一个参数必须为Struct指针")
		return nil, error
	}
	if nom, ok = o[1].(string); !ok {
		error := errors.New("第二个参数必须为方法名字符串")
		return nil, error
	}
	m := v.MethodByName(nom)
	if !m.IsValid() {
		error := errors.New("调用的方法不存在")
		return nil, error
	}
	if m.Type().NumIn() != len(o)-2 {
		error := errors.New("方法参数个数不正确")
		return nil, error
	}
	for i := 2; i < len(o); i++ {
		if reflect.TypeOf(o[i]) != m.Type().In(i-2) {
			error := errors.New("方法参数类型不正确")
			return nil, error
		}
		args = append(args, reflect.ValueOf(o[i]))
	}
	r := m.Call(args)
	for i := range r {
		rtns = append(rtns, r[i].Interface())
	}
	return rtns, nil
}
