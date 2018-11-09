package tfsql

type tfsqlResult struct {
	affectedRows int64
	insertId     int64
}

func (res *tfsqlResult) LastInsertId() (int64, error) {
	return res.insertId, nil
}

func (res *tfsqlResult) RowsAffected() (int64, error) {
	return res.affectedRows, nil
}
