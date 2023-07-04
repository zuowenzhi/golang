package main

import (
	"fmt"
	//为了使用utils.go文件的变量或者函数，我们需要先引入该model包
	//"chapter03/demo04/model"
)

//变量使用的注意事项
func main() {

	//该区域的数据值可以在同一类型范围内不断变化
	//局部变量只能在当前函数里边使用,不能在其他函数中使用
	//在一个函数,代码块中,函数名不能重复
	//
	var i int = 10
	i = 30
	i = 50
	fmt.Println("i=", i)
	//i = 1.2 //int ,原因是不能改变数据类型

	//变量在同一个作用域(在一个函数或者在代码块)内不能重名
	//var i int = 59
	//i := 99
	//我们使用utils.go 的heroName 包名.标志符
	//fmt.Println(model.HeroName)

}
