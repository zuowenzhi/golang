package main

import "fmt"

//定义全局变量
var n1 = 100
var n2 = 200
var name = "jack"

//上面的声明方式，也可以改成一次性声明
var (
	n3    = 300
	n4    = 900
	name2 = "mary"
)

func demo() {
	fmt.Println("自定义函数")
}

func main() {
	demo()

	//该案例演示golang如何一次性声明多个变量
	//var n1, n2, n3 int
	//fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)

	//一次性声明多个变量的方式2
	// 应该将=的左右两边分开来看,=左边有3个变量,=右边有3个值,
	// 按照顺序,n1,name,n3的值分别是100,888,tom
	//var n1, name, n3 = 100, 888, "tom"//定义的同时赋值
	//fmt.Println("n1=", n1, "name=", name, "n3=", n3)

	//一次性声明多个变量的方式3, 同样可以使用类型推导
	//n1, name , n3 := 100, "tom~", 888
	//fmt.Println("n1=",n1, "name=", name, "n3=", n3)

	//输出全局变量 在函数外边生成的变量,就是全局变量,
	fmt.Println("n1=", n1, "name=", name, "n2=", n2)
	fmt.Println("n3=", n3, "name2=", name2, "n4=", n4)

}
