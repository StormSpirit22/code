package main

import (
	"fmt"
	"reflect"
)

type X int
type Y int

func main() {
	testElem2()
}

func typeOf() {
	var a, b X = 100, 200
	var c Y = 300
	ta, tb := reflect.TypeOf(a), reflect.TypeOf(b)
	tc := reflect.TypeOf(c)

	fmt.Println(ta, tb, tc)
	fmt.Println(ta == tb, ta == tc, ta.Kind() == tc.Kind())
	fmt.Println(reflect.ValueOf(a))
}

func testElem() {
	fmt.Println(reflect.TypeOf(map[string]float64{}).Elem())
	fmt.Println(reflect.TypeOf([]int32{}).Elem())
}

type user struct {
	name string
	age int
}

type manager struct {
	user
	title string
}


func testElem2() {
	var m manager
	t := reflect.TypeOf(&m)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Type, f.Offset)

		if f.Anonymous {													// 输出匿名字段结构
			for x := 0; x < f.Type.NumField(); x++ {
				af := f.Type.Field(x)
				fmt.Println(" ", af.Name, af.Type)
			}
		}
	}
}