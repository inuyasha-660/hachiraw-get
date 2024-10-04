package main

import (
	"INIGet/utils"
	"fmt"
)

func init() {
	fmt.Print(utils.ASCII)
	fmt.Println(utils.INFO)
	fmt.Println(utils.LINE)
}

func main() {
	var url_input string
	var paint_input_start int
	var paint_input_end int

	fmt.Print("输入目标链接:")
	fmt.Scanln(&url_input)

	fmt.Print("==> 连通性测试... ")
	get_test_string, get_test_bool := utils.Get_test(url_input)
	if get_test_bool {
		fmt.Println("成功")
	} else {
		fmt.Println("失败")
		panic(get_test_string)
	}

	fmt.Print("==> 发送请求[Get]... ")
	utils.Get(url_input)

	fmt.Print("输入目标画开始:")
	fmt.Scanln(&paint_input_start)
	fmt.Print("输入目标画结束:")
	fmt.Scanln(&paint_input_end)
	utils.Get_paint(paint_input_start, paint_input_end, url_input)
}
