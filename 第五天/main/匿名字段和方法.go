package main

import "fmt"

//匿名字段相当于正常面向对象编程中的继承
//匿名字段的初始化
//人
type Person struct {
	name string
	sex  string
	age  int
}

//学生
type Student struct {
	Person //匿名字段，那么默认student就包含了person的所有字段,只有类型，没有名字
	id     int
	addr   string
}

//结构体指针类型
type Student1 struct {
	*Person
}

//方法：给某个类型绑定一个函数,相当于正常面向对象中的封装函数，go语言中的方法不支持重载，也就是说不能定义名字相同但是
//不同参数的方法,但是接收者类型不同可以。调用者（也叫作接收者）的类型不能是接口或是指针
//只要接收者类型不一样，即使同名，也是不同的方法，不会出现重复
func (student Student) link() (age_id int) { //值语义
	age_id = student.age + student.id
	return
} //如果要是想修改则应该传递一个指针，默认为值引用

//通过一个函数，给成员赋值，引用语义
func (p *Student) SetInfo(name string, sex string, age int) {
	p.age = age
	p.sex = sex
	p.name = name
}

/*
方法集：
类型的方法集是指可以被该类型的值调用的所有方法的集合，无论是值语义还是引用语义，编译器总能查找全部方法，转化工作内部自动实现
*/

/*
方法的继承：
如果匿名字段实现了一个方法，那么包含这个匿名字段的struct也能调用该方法,子类也可以调用父类的方法，会覆盖父类方法
若是想调用父类方法，可以进行显式调用
s.person.balabla  显式调用
方法值相当于函数指针，函数类型
*/

//给person类型实现一个方法
func (p Person) MyPrint() {
	fmt.Println("person的方法")
}

func main() {
	//顺序初始化，必须所有参数都要赋值
	s1 := Student{Person{"mike", "18", 18}, 01, "beijing"} //结构体变量初始化，或者写成var s1 student=student{}
	fmt.Println(s1)
	fmt.Printf("s1=%+v\n", s1) //对结构体变量显示更详细的参数   %+v

	//指定成员初始化，没有初始化的则使用默认值
	s2 := Student{id: 1, Person: Person{name: "mike"}}
	fmt.Println(s2)

	//成员的操作
	var s3 Student //变量声明
	//给成员赋值
	s3.name = "mike" //等价于s3.Person.name="mike"
	s3.Person = Person{age: 10}

	//指针变量两种赋值
	//s4:=Stuent1{&Person{"mike","m",18}}

	var s5 Student1
	s5.Person = new(Person)
	s5.name = "yoyo"
	fmt.Println("========================================")
	fmt.Println("方法的应用例子：")
	s6 := Student{id: 01, Person: Person{age: 18}}
	fmt.Println(s6.link())
	fmt.Println("========================================")
	var s7 Student
	(&s7).SetInfo("lily", "w", 19) //不要忘记取地址符
	fmt.Printf("s7=%+v\n", s7)
	fmt.Println("========================================")
	s5.MyPrint()           //student 继承了person，也就继承了person的方法
	a := s5.MyPrint        //方法值，调用函数时，无需再传递接收者，隐藏了接收者
	a()                    //等价于s5.MyPrint()
	b := (Student).MyPrint //方法表达式，直接对类型进行操作，但是需要显式的传入接收者，但是不会完成自动转换的工作，b(s5)会报错

}
