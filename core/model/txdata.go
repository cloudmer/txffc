package model

type Txdata struct {
	Id     			int
	Alias  		 	string
	DataTxt			string
	Start			int
	End				int
	RegretNumber 	int
	Forever 		int
	State			int
	Time			int
}

func (Txdata) Query() []*Txdata  {
	sql_str := `SELECT * FROM txdata WHERE state=1;`
	rows, err := DB.Query(sql_str)
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	data := make([]*Txdata, 0)
	for rows.Next() {
		rows.Columns()
		model := new(Txdata)
		var err error
		err = rows.Scan(&model.Id, &model.Alias, &model.DataTxt, &model.Start, &model.End, &model.RegretNumber, &model.Forever, &model.Start, &model.Time)
		if err != nil {
			panic(err.Error())
		}
		data = append(data, model)
	}
	return data
}