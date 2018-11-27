package main//每个go文件必须具有package关键字，且在第一行，表示go文件属于哪个包

import "fmt"

//关键字，标示符，注释，基础结构
//go 的注释与C，C++一样
/*
多行注释
 */

 //类型都是在名字之后定义，在这块即使不定义也可以，go会自动推断
const NAME int = 56  //常量的名字最好是大写
var name  string= "imooc"   //在main函数之外定义的变量就是全局变量，可以在包里面被调用

//一般类型声明,相当于类型别名
type ChenInt int

//结构的声明
type learn struct {

}

//接口的声明
type Ilearn interface {  //接口的名字一般以I开头

}
//函数定义
func machine_learning()  {

}

 //记不住包名没有问题，只要记住函数名即可，goland会自动补全，引入多个包时，goland会
 //自动添加括号
//main函数
func main()  {
	fmt.Println(NAME)
	fmt.Println(name)
}





//package，import别名，路径，".",   "_"的使用说明

/*
1.package是最基本的分发单位和工程管理中依赖关系的体现
2.每个go语言源代码文件开头都拥有一个package声明，表示源码文件所属代码包
3.要生成go语言可执行程序，必须要有main的package包，且必须在该包下有main（）函数
4.同一个路径下只能存在一个package，一个package可以拆成多个源文件组成；所以在同一个
	工程中一般都会出现很多文件夹
 */
 //package 包名 尽量与所属文件夹名称一致

/*
go基础语法----import
1，import语句可以导入源代码文件所依赖的package包
2，不得导入源代码文件中没有用到的package，否则go语言编译器会报编译错误
3，import语法格式主要有以下两种:
	(1)第一种	import "package1"
			import "package2"
			import "package3"
	（2）第二种  import（    注意是小括号，而不是大括号
						"package1"
						"package2"
						"package3")
import 原理：
	1.如果一个main导入其他包，包将被顺序导入
	2.如果导入的包中依赖其他包（包B），会首先带入B包，然后初始化B包中的常量和变量，
		最后如果B包中有init，会自动执行init（）
	3.所有包导入完成后才会对main中常量和变量进行初始化，然后执行main中的init函数（如果存在）
		，最后执行main函数
	4.如果一个包被导入多次，则该包只会被导入一次

综上：所以代码执行的顺序与导入包的顺序有直接关系，不同的导入顺序可能会有完全不同的输出
 */








//go变量，函数可见规则