//本项目实现了一个简单的命令行交换的婚姻问候

package main

import (
	"flag"
	"fmt"
	"os"
)

// 自定义错误提示
func Usage() {
	fmt.Fprintf(os.Stderr, "用法:[options] [name A] [name B]\n")
	flag.PrintDefaults()
}

func main() {

	//定义命令行参数
	//这些变量都是 指针 类型!
	greeting := flag.String("g", " New couples , have a nice day!", "打印问候")
	married := flag.Bool("m", true, "是否已婚")

	flag.Usage = Usage //自定义错误
	flag.Parse()       // 必须调用

	arg := flag.Args() //要写在 parse 语句后面 否则无法正常解析
	name := [2]string{"", ""}

	//未婚:
	if !(*married) {
		fmt.Println("您还未婚!无法输出!")
		Usage()
		return
	}
	//名字个数不对
	if len(arg) != 2 {
		fmt.Println("请输入您和您的配偶名字!")
		Usage()
		return
	}

	//已婚:
	name = [2]string(arg)
	fmt.Println(name[0], name[1], *greeting)

}
