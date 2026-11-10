package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// 导入 net/http 会触发 cgo（因为 net 包依赖 cgo 解析器）
	// 这里仅演示导入，不实际发起网络请求
	fmt.Println("Hello, Go 1.22")
	fmt.Printf("HTTP Package imported: %v\n", http.MethodGet != "")

	// 退出码 0
	os.Exit(0)
}
