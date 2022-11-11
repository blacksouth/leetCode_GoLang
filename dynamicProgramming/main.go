package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now() // 获取当前时间
	fmt.Println(Fib(45))
	elapsed := time.Since(start)
	fmt.Println("该函数执行完成耗时：", elapsed)

}
