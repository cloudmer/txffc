package main

import (
	"log"
	"os"
	"runtime"
	"time"
	"txffc/core/src/txffc"
)

func main()  {
	log.Println("服务启动中．．．　进程ID:", os.Getpid())
	runtime.GOMAXPROCS(runtime.NumCPU())

	for  {
		select {
		// 腾讯分分彩 1分钟一开奖 10秒一计算
		case <-time.After(10 * time.Second):
			txffc.Calculation()
		}
	}
}
