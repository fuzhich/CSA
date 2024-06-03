//package main
//
//import (
//	"fmt"
//	"reflect"
//)
//
//type 猿神 int   //int特征的新类型
//type 原神 = int //别名
//
//func main() {
//	var a 猿神
//	var b 原神
//
//	fmt.Print("type of a:%T\n", a)
//	fmt.Print("type of b:%T\n", b)
//
//	rfTypeOf(a)
//	rfTypeOf(b)
//
//	TypeOf(a)
//	TypeOf(b)
//}
//func TypeOf(data interface{}) {
//	switch data.(type) {
//	case 猿神:
//
//		fmt.Println("Type is int")
//	case 原神:
//		fmt.Println("Type is int")
//	default:
//		fmt.Println("Type Not Found")
//	}
//}
//func rfTypeOf(data interface{}) {
//	of := reflect.TypeOf(data)
//	fmt.Println(of)
//}

//结构体的定义
//type 类型名 struct{
//	字段名 类型名
//}

//package main
//
//import "fmt"
//
//type yuanshen struct {
//	Name string
//	Age  int
//}
//
//func main() {
//
//	a := yuanshen{
//		Name: "yuan1",
//		Age:  18,
//	}
//
//	//我觉得最明了的方法
//	var b yuanshen
//	b.Name = "yuan2"
//	b.Age = 18
//
//	c := struct {
//		Name string
//		Age  int
//	}{
//		"yuan3",
//		18,
//	}
//
//	d := NewYuanShen("猿神", 18)
//
//	fmt.Printf("%#v\n", a)
//	fmt.Printf("%#v\n", b)
//	fmt.Printf("%#v\n", c)
//	fmt.Printf("%#v\n", d)
//
//	//在Go中通过字母大小写来控制可见性的，如果定义的常量、变量、类型、接口、结构、函数等的名称是
//	//大写字母开头表示能被其它包访问或调用（相当于public），非大写开头就只能在包内使用（相当于p
//	//rivate，变量或常量也可以下划线开头）
//
//	//结构体的访问
//}
//
//// 构造函数
//func NewYuanShen(name string, age int) *yuanshen {
//	return &yuanshen{
//		Name: name,
//		Age:  age,
//	}
//}

