package model

type AnalysisTxffc struct {
	Id  int
	FrontThreeLuckyTxt  interface{}
	CenterThreeLuckyTxt interface{}
	AfterThreeLuckyTxt  interface{}
}

func (AnalysisTxffc) GetAnalysis (txffc_id string) *AnalysisTxffc {
	str_sql := `SELECT id, front_three_lucky_txt, center_three_lucky_txt, after_three_lucky_txt FROM analysisTxffc WHERE txffc_id=` + txffc_id
	rows, err := DB.Query(str_sql)
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	data := new(AnalysisTxffc)
	for rows.Next() {
		rows.Columns()
		var err error
		err = rows.Scan(&data.Id, &data.FrontThreeLuckyTxt, &data.CenterThreeLuckyTxt, &data.AfterThreeLuckyTxt)
		if err != nil {
			panic(err.Error())
		}
	}
	return data
}