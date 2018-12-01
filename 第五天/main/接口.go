package main

import (
	"fmt"
)

/*
在go语言中，接口是一个自定义类型，接口类型具体描述了一系列方法的集合
接口类型是一种抽象的类型，他不会暴露出他所代表的对象的内部值的结构和这个对象支持的基础操作的集合，他们只会展示出他们自己的
方法，因此接口类型不能实例化，go通过接口实现了鸭子类型，并不关心对象是什么类型，只关心行为，不关心由哪个类型实现，只关心实现之后的行为
接口默认以er结尾或者I开头，接口只有方法声明，没有数据字段，接口可以匿名嵌入其他接口，或嵌入到结构中
接口可以实现多态
*/

//定义接口类型
type Humaner interface {
	//方法，只有声明，没有实现，由别的类型（自定义类型）实现
	sayhi()
}

type Person interface {
	Humaner
	singing()
}

type Student02 struct {
	name string
	id   int
}

type Teacher struct {
}

func (s Teacher) sayhi() {
	fmt.Println("hello,this is a teacher")
}

//student02这个类型实现了方法
func (s Student02) sayhi() {
	fmt.Println("hello,this is a student")
}

/*
实现多态的例子
定义一个普通函数，函数的参数为接口类型
只有一个函数，可以有不同的表现，多态，调用同一接口
*/
func Whosayhi(i Humaner) {
	i.sayhi()
}

//Whosayhi(s)

/*
接口的继承
如果一个接口1 作为接口2的嵌入字段，那么接口2就隐式的包含了接口1里面的方法
*/

/*
接口的转换
超集可以转换为子集，反过来不可以,超集的变量可以赋给子集
*/

/*
空接口 不包含任何方法，正因为如此，所有类型都实现了空接口，因此空接口可以存储任意类型的数值，
当函数可以接收任意的对象实例时，我们会将其声明为空接口，标准库中很多这种函数,空接口经常作为匿名函数，进行传参
*/
func MyPrintln(args ...interface{}) {
	fmt.Println(args)
}

/*
类型查询/类型断言
我们知道interface变量可以接收任意类型，通过comma-OK断言和switch测试，两种方式来反向知道这个变量实际保存了哪种类型
*/

func main() {
	//定义一个接口类型的变量
	var a Humaner
	var b Humaner

	//只要实现了此接口方法的类型，那么这个类型的变量（接收者类型）就可以给a赋值
	var s Student02
	var t Teacher
	a = s
	b = t
	a.sayhi() //调用同一个接口，拥有不同表现
	b.sayhi()

	//定义语句要在函数体内，不要在函数体外
	//创建一个切片
	//x:=make([]int,3,5 )
	//fmt.Println(len(x),cap(x))
	fmt.Println("========================================")
	//创建一个切片
	x := make([]Humaner, 2) //这种写法应该检查边界，确保slice中每个元素都被赋值，否则将报错，若改成3，则会报错
	x[0] = s
	x[1] = t

	for _, i := range x {
		i.sayhi()
	} //第一个值返回的是下标。第二个返回的是下标所对应的值
	fmt.Println("========================================")
	fmt.Println("空接口")
	var v1 interface{} = 1
	fmt.Println(v1)
	MyPrintln(v1)

	fmt.Println("========================================")
	//类型查询
	i := make([]interface{}, 2)
	i[0] = 1
	i[1] = "hello go "

	for index, data := range i {

		//第一个返回的是值，第二个返回的是判断结果的真假
		if value, ok := data.(int); ok == true { //if 语句也要求同一行
			fmt.Println("是int")
		}
	}

}
