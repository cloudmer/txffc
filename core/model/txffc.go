package model

import (
	"strconv"
)

type Txffc struct {
	Id    int
	Qishu string
	One   string
	Two   string
	Three string
	Four  string
	Five  string
	Time  int
}

func (Txffc) Query(limit string) []*Txffc {
	sql_str := `SELECT * FROM (
					SELECT id,qishu,one,two,three,four,five,time FROM txffc  ORDER BY time DESC LIMIT ` + limit + `
				) AS ssc ORDER BY time ASC`

	rows, err := DB.Query(sql_str)
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	data := make([]*Txffc, 0)

	for rows.Next() {
		rows.Columns()
		cq := new(Txffc)
		var err error
		err = rows.Scan(&cq.Id, &cq.Qishu, &cq.One, &cq.Two, &cq.Three, &cq.Four, &cq.Five, &cq.Time)
		if err != nil {
			panic(err.Error())
		}
		data = append(data, cq)
	}
	return  data
}

// 获取最新开奖号
func (Txffc) GetNewsCode() string {
	sql_str := `SELECT one,two,three,four,five FROM txffc  ORDER BY time DESC LIMIT 1`
	rows, err := DB.Query(sql_str)
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}


	type dataStruct struct {
		One string
		Two string
		Three string
		Four string
		Five string
	}

	data := new(dataStruct)
	for rows.Next() {
		rows.Columns()
		var err error
		err = rows.Scan(&data.One, &data.Two, &data.Three, &data.Four, &data.Five)
		if err != nil {
			panic(err.Error())
		}
	}
	return data.One+data.Two+data.Three+data.Four+data.Five
}

// 查询 当前数据包 前三中奖最新一次的出现位置
func (Txffc) GetCodes(packet_id string) []*Txffc {
	sql_q3 := `SELECT txffc.id FROM txffc LEFT JOIN analysisTxffc ON(txffc.id=analysisTxffc.txffc_id) WHERE analysisTxffc.front_three_lucky_txt != '' AND analysisTxffc.type = ` + packet_id + ` ORDER BY txffc.time DESC LIMIT 1;`
	sql_z3 := `SELECT txffc.id FROM txffc LEFT JOIN analysisTxffc ON(txffc.id=analysisTxffc.txffc_id) WHERE analysisTxffc.center_three_lucky_txt != '' AND analysisTxffc.type = ` + packet_id + ` ORDER BY txffc.time DESC LIMIT 1;`
	sql_h3 := `SELECT txffc.id FROM txffc LEFT JOIN analysisTxffc ON(txffc.id=analysisTxffc.txffc_id) WHERE analysisTxffc.after_three_lucky_txt != '' AND analysisTxffc.type = ` + packet_id + ` ORDER BY txffc.time DESC LIMIT 1;`


	var q3_lucky_index_id int
	var z3_lucky_index_id int
	var h3_lucky_index_id int
	DB.QueryRow(sql_q3).Scan(&q3_lucky_index_id)
	DB.QueryRow(sql_z3).Scan(&z3_lucky_index_id)
	DB.QueryRow(sql_h3).Scan(&h3_lucky_index_id)

	var least_id int
	if q3_lucky_index_id < z3_lucky_index_id {
		least_id = q3_lucky_index_id
	} else {
		least_id = z3_lucky_index_id
	}

	if least_id < h3_lucky_index_id {
		if q3_lucky_index_id < z3_lucky_index_id {
			least_id = q3_lucky_index_id
		} else {
			least_id = z3_lucky_index_id
		}
	} else {
		least_id = h3_lucky_index_id
	}

	str_sql := `SELECT id,qishu,one,two,three,four,five,time FROM txffc WHERE id >= `+ strconv.Itoa(least_id) +` ORDER BY time ASC`
	rows, err := DB.Query(str_sql)
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	data := make([]*Txffc, 0)

	for rows.Next() {
		rows.Columns()
		cq := new(Txffc)
		var err error
		err = rows.Scan(&cq.Id, &cq.Qishu, &cq.One, &cq.Two, &cq.Three, &cq.Four, &cq.Five, &cq.Time)
		if err != nil {
			panic(err.Error())
		}
		data = append(data, cq)
	}
	return  data
}
