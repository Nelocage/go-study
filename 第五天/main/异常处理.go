package main

import (
	"errors"
	"fmt"
)

/*
普通情况返回error，致命错误返回panic，比如数组访问越界，空指针引用等，一般情况下不应通过调用panic来报告普通的错误
而应该只把他作为报告致命错误的一种方式，
当panic异常发生时，程序会中断运行，并立即执行在该goroutine中被延迟的函数（defer机制）。随后程序崩溃并输出日志信息
日志信息包括panic value和函数调用的堆栈跟踪信息


不是所有的panic异常都来自运行时，直接调用内置的panic函数也会引发panic异常，panic函数接收任何值作为参数
*/
func MyDiv(a, b int) (result int, err error) {

	defer func() {
		if err := recover(); err != nil { //如果没发生错误就不打印 ，go语言中的if语句是可以初始化的
			fmt.Println(err) //可以打印panic信息
		}

	}() //匿名函数需要加括号来调用

	if b == 0 {
		err = errors.New("除数不应为0 ")
		return

	} else {
		result = a / b
		return
	}

}

/*
recover函数
运行时panic异常一旦被引发就会导致程序崩溃，go语言提供了专门用于“拦截”运行时panic的内建函数，他可以使当前程序从运行时
panic的状态中恢复并重新获得流程控制权，recover只有在defer调用的函数中有效
*/

func main() {
	err1 := errors.New("this is a normal err1")
	fmt.Println(err1)

	b, error := MyDiv(1, 0)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(b)
	}
	//显式调用panic
	//panic("panic")

}
