package main

import (
	"runtime"
	"fmt"
)

func say(s string)  {
	for i := 0;i<5;i++{
		runtime.Gosched()//表示让cpu把时间片让给别人，下次某个时候继续恢复执行该goroutine
		fmt.Println(s, i)
	}
}

func main()  {
	go say("world")
	say("hello")
}