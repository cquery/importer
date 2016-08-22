package pgsql

import (
	"testing"

	"database/sql"
	"database/sql/driver"
	testdb "github.com/erikstmartin/go-testdb"
)

type testResult struct {
	lastId       int64
	affectedRows int64
}

func (r testResult) LastInsertId() (int64, error) {
	return r.lastId, nil
}

func (r testResult) RowsAffected() (int64, error) {
	return r.affectedRows, nil
}

func Test_UpdaterOneSet(t *testing.T) {

	testdb.SetExecWithArgsFunc(func(query string, args []driver.Value) (result driver.Result, err error) {
		if args[0] == "i-12345" {
			return testResult{1, 1}, nil
		}
		return testResult{1, 0}, nil
	})

	db, _ := sql.Open("testdb", "")

	updater := &Updater{
		db: db,
	}

	set := updater.Set("ec2")
	set.AddString("instance_id", "i-12345")
	set.AddString("image_id", "ami-12345")

	if err := updater.Update(set); err != nil {
		t.Error(err)
		return
	}

}
