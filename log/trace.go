package log

import (
	"fmt"
	"time"
)

//用途：用于在函数进入时和退出时打印日志，其中包含耗时
//用法：
func testFunc() {
	defer TraceLog("testFunc")()
}

func TraceLog(funcName string) func() {
	start := time.Now()
	fmt.Printf("enter func: %s\n", funcName)
	return func() {
		fmt.Printf("exit func: %s, (%s)\n", funcName, time.Since(start))
	}
}
