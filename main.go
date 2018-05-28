package main

import (
	"log"
	"os"
	"runtime"
	"time"
	"txffc/core/src/txffc"
	"txffc/core/src/ssccycle"
)

func main()  {
	log.Println("服务启动中．．．　进程ID:", os.Getpid())
	runtime.GOMAXPROCS(runtime.NumCPU())

	for  {
		select {
		// 腾讯分分彩 1分钟一开奖 10秒一计算
		case <-time.After(10 * time.Second):
			// 腾讯分分彩
			txffc.Calculation()
			// 腾讯分分彩 a包连续 周期 计算
			ssccycle.Calculation()
		}
	}
}
