package model

type SscCycle struct {
	Id 			int
	Alias 		string
	DataTxt 	string
	Start 		int
	End 		int
	Continuity 	int
	Bnumber		int
	Status 		int
	Cycle 		int
	CreatedAt	string
}

func (SscCycle) Query() []*SscCycle  {
	sql_str := `SELECT * FROM ssc_cycle WHERE status=1;`
	rows, err := DB.Query(sql_str)
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	data := make([]*SscCycle, 0)
	for rows.Next() {
		rows.Columns()
		model := new(SscCycle)
		var err error
		err = rows.Scan(&model.Id, &model.Alias, &model.DataTxt, &model.Start, &model.End, &model.Continuity, &model.Bnumber, &model.Status, &model.Cycle, &model.CreatedAt)
		if err != nil {
			panic(err.Error())
		}
		data = append(data, model)
	}
	return data
}