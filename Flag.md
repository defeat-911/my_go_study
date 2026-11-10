#  Flag

## `os.Args`

以下示例程序用于获取命令行参数(只是简单获取命令行参数)

```go
package main

import (
    "fmt"
    "os"
)

//os.Args demo
func main() {
    //os.Args是一个[]string
    if len(os.Args) > 0 {
        //这里的判断始终为真 
        //Go 的运行时在启动程序时会把命令行填进 os.Args
        //其中第 0 个元素固定是程序自身即程序的名字或路径。
        for index, arg := range os.Args {
            fmt.Printf("args[%d]=%v\n", index, arg)
        }
    }
```

需要注意的是 `os.Arg` 是一个存储字符串的切片,第一个元素为文件的名称,所以在遍历`os.Arg` 的时候要从 1 开始

---

## Flag

### flag的定义

**`flag.Type()`**

> flag.Type(flag名, 默认值, 帮助信息)*Type

````go
name := flag.String("name", "张三", "姓名")
age := flag.Int("age", 18, "年龄")
married := flag.Bool("married", false, "婚否")
delay := flag.Duration("d", 0, "时间间隔")
````

*定义的变量为对应类型的指针,使用这些变量时要解引用*



**`flag.TypeVar`**

```go
var name string
var age int
var married bool
var delay time.Duration
flag.StringVar(&name, "name", "张三", "姓名")
flag.IntVar(&age, "age", 18, "年龄")
flag.BoolVar(&married, "married", false, "婚否")
flag.DurationVar(&delay, "d", 0, "时间间隔")
```

通过`TypeVar()`变量写更加方便(变量引用)



### flag.Parse

*parse : (中文) 解析*

通过以上两种方法定义好命令行flag参数后，需要通过调用**flag.Parse()**来对命令行参数进行解析

命令行参数格式:

1. `--flag xxx` (使用空格赋值)
2. `--flag=xxx`(用=赋值)

*--flag 两个-是我们熟知的GNU风格*



### flag其他函数

- `flag.Args()` //返回命令行未被解析的参数，以**`[]string`**类型
- `flag.NArg()` //返回命令行参数后的其他参数个数
- `flag.NFlag()` //返回使用的命令行参数个数

### flag.usage

`flag.usage` 是默认情况下的报错:

```bash
$go run .\01flag.go --unknown                           
flag provided but not defined: -unknown                                             
Usage of C:\Users\Administrator\AppData\Local\go-build\23\23adc171e0fa527a4012fe59c4f3345c39ae2876ae6cd6d33f6a82abd6cbf7c2-d\01flag.exe:                                               
  -age int                                                                          
        结婚人的年龄                                                                
  -married                                                                          
        是否已婚(默认未婚)                                                          
  -name string                                                                      
        结婚人的名字  
```

但是我们可以通过定义一个函数来达到我们想要的报错内容:

```go
// 自定义错误提示
func Usage() {
	fmt.Fprintf(os.Stderr, "用法:\n")
	flag.PrintDefaults()
}
//然后在main里调用
flag.Usage = Usage //自定义错误
```

这样子新的报错:

```bash
$go run .\01flag.go --unknown

flag provided but not defined: -unknown
用法:
  -age int
        结婚人的年龄
  -married
        是否已婚(默认未婚)
  -name string
        结婚人的名字
exit status 2
```



## 完整代码

```go
package main

import (
	"flag"
	"fmt"
	"os"
)

// 自定义错误提示
func Usage() {
	fmt.Fprintf(os.Stderr, "用法:\n")
	flag.PrintDefaults()
}

func main() {

	//定义三个命令行参数 name age married
	//这些变量都是 指针 类型!
	name := flag.String("name", "", "结婚人的名字")
	age := flag.Int("age", 0, "结婚人的年龄")
	married := flag.Bool("married", false, "是否已婚(默认未婚)")

	flag.Usage = Usage //自定义错误
	flag.Parse()       // 必须调用

	//返回的是内存地址,因为变量类型是指针
	fmt.Println(name, age, married)
	fmt.Println(*name, *age, *married)

	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())

	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())

	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())

}

```

```bash
PS E:\collage\_project\01Go> go run .\01flag.go --name 张三 --age 22 --married=True
0x2ed49421c030 0x2ed49420e0a8 0x2ed49420e0d0
张三 22 true
[]
0
3
```

---

# Flag实战

本项目实现了一个简单的命令行交换的婚姻问候,主要的功能如下:

1. 输入两个名字,在终端输出`nameA nameB New couples , have a nice day!`的字样
2. 错误处理: 针对输入名字不为2,可以调用提示提醒用户
3. 自定义问候:通过`-g`参数,自定义问候内容

```go
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
	}
	//名字个数不对
	if len(arg) != 2 {
		fmt.Println("请输入您和您的配偶名字!")
		Usage()
	}

	//已婚:
	name = [2]string(arg)
	fmt.Println(name[0], name[1], *greeting)

}

```

## 项目改进 

### 逻辑判断

在第一次开发中,采用嵌套的`if`语句,显得十分臃肿,不方便识别:

```go
//已婚:
  if *married {
      if len(arg) >= 2 {
          name = [2]string(arg)
          fmt.Println(name[0], name[1], *greeting)
      }

      if len(arg) < 2 {
          fmt.Println("请输入您和您的配偶名字!")
          Usage()
      }
  }

  //未婚:
  if !(*married) {
      fmt.Println("您还未婚!无法输出!")
      Usage()
  }
```

改进后,利用卫语句,使得代码清晰易读:

> *卫语句的核心就是:先把异常情况处理掉并退出,让主流程留
>   在最外层。*

```go
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
```

