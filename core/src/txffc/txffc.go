package txffc

import (
	"txffc/core/model"
	"fmt"
	"strings"
	"log"
	"time"
	"txffc/core/logger"
	"strconv"
	"txffc/core/mail"
)

// 数据包
var contain_datapackage []*model.Txdata

// 最新开奖号
var new_code string = ""

// 计算
func Calculation()  {
	// 查询
	txffc := new(model.Txffc)

	// 最新开奖号
	newcode := txffc.GetNewsCode()

	// 是否重复计算
	if new_code != "" && new_code == newcode {
		fmt.Println("最新一期 腾讯分分彩 已经计算过了 等待新号码出现")
		return
	}

	// 刷新 最新开奖号
	new_code = newcode

	// 查询数据包
	txdata := new(model.Txdata)
	contain_datapackage = txdata.Query()

	for i := range contain_datapackage {
		go containAnalysisCodes(contain_datapackage[i])
	}
}

// 开始任务
func containAnalysisCodes(packet *model.Txdata)  {
	slice_dataTxt := strings.Split(packet.DataTxt, "\r\n")
	//slice data txt to slice data txt map
	dataTxtMap := make(map[string]string)
	for i := range slice_dataTxt {
		dataTxtMap[slice_dataTxt[i]] = slice_dataTxt[i]
	}

	//检查是否在报警时间段以内
	if (packet.Start >0 && packet.End >0) && (time.Now().Hour() < packet.Start || time.Now().Hour() > packet.End)  {
		log.Println("彩票类型: 腾讯分分彩", "数据包别名:", packet.Alias, "报警通知非接受时间段内")
		logger.Log("彩票类型: 腾讯分分彩 数据包别名: " + packet.Alias + " 报警通知非接受时间段内 ")
		return
	}


	// 查询 开奖号
	txffc := new(model.Txffc)
	analysisTxffc := new(model.AnalysisTxffc)
	codes := txffc.GetCodes(strconv.Itoa(packet.Id))

	var q3_lucky_number int
	var z3_lucky_number int
	var h3_lucky_number int

	for i := range codes {
		//获取当前彩种 分析数据中的 数据包 N 的分析数据
		anaData := analysisTxffc.GetAnalysis(strconv.Itoa(codes[i].Id))

		//当前 N 期内 前3号码 中过奖
		if anaData.FrontThreeLuckyTxt != nil {
			q3_lucky_number = 0
		} else  {
			q3_lucky_number += 1
		}

		//当前 N 期内 中3号码 中过奖
		if anaData.CenterThreeLuckyTxt != nil {
			z3_lucky_number = 0
		} else {
			z3_lucky_number += 1
		}

		//当前 N 期内 后3号码 中过奖
		if anaData.AfterThreeLuckyTxt != nil {
			h3_lucky_number =0
		} else {
			h3_lucky_number += 1
		}
	}

	//初始化邮件内容
	mail_contents := ""
	//前三 中奖次数是否达到 报警状态 大于报警期数 不报警 等到周期走完
	if q3_lucky_number == packet.RegretNumber {
		mail_contents += "前:" + strconv.Itoa(packet.RegretNumber) + " N<br/>"
	}

	//中三 中奖次数是否达到 报警状态 大于报警期数 不报警 等到周期走完
	if z3_lucky_number == packet.RegretNumber {
		mail_contents += "中:" + strconv.Itoa(packet.RegretNumber) + " N<br/>"
	}

	//后三 中奖次数是否达到 报警状态 大于报警期数 不报警 等到周期走完
	if h3_lucky_number == packet.RegretNumber {
		mail_contents += "后:" + strconv.Itoa(packet.RegretNumber) + " N<br/>"
	}

    title := "通知类型: 腾讯分分彩 - 当前" + strconv.Itoa(packet.RegretNumber) + "期 警报<br/>"

	//如果 达到报警条件 则报警
	if mail_contents != "" {
		mail_contents = title + mail_contents

		// 发送邮件
		go mail.SendMail(title, mail_contents)
	}
}