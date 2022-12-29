package pkg

import (
	"fmt"
	"runtime"
)

func FuncName() {
	pt, _, _, _ := runtime.Caller(1)
	funcName := fmt.Sprintf("func=%+v", runtime.FuncForPC(pt).Name())
	fmt.Println(">>>", funcName)
}
