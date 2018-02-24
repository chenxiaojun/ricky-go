package main

import (
	"fmt"
	"reflect"
)

func main() {
	u := User{"张三", 20}
	// reflect.TypeOf可以获取任意对象的具体类型
	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)
	fmt.Println(t)
	fmt.Println(v)
	fmt.Println(reflect.TypeOf("Ricky"))
	fmt.Printf("%T\n", u)
	fmt.Printf("%v\n", u)

	u1 := v.Interface().(User)
	fmt.Println(u1)
	fmt.Println(v.Type())
	fmt.Println(v.Kind())

	fmt.Println("")

	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Name)
	}

	for i := 0; i < t.NumMethod(); i++ {
		fmt.Println(t.Method(i).Name)
	}

	animal := Animal{"Kitty"}

	an := reflect.ValueOf(animal)
	fmt.Println(an)

	mPrint := an.MethodByName("Print")
	args := []reflect.Value{reflect.ValueOf("哈哈")}
	fmt.Println(mPrint.Call(args))
	fmt.Println(reflect.ValueOf("哈哈"))
	fmt.Println([]reflect.Value{reflect.ValueOf("哈哈")})
}

type User struct {
	Name string
	Age  int
}

type Animal struct {
	Name string
}

func (a Animal) Print(prfix string) {
	fmt.Printf("%s:Name is %s", prfix, a.Name)
}
