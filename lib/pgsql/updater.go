package pgsql

import (
	"github.com/cquery/importer/lib"

	_ "github.com/lib/pq"

	"database/sql"
	"fmt"
	"strings"
)

const (
	sql_upsert = "UPSERT INTO %s (%s) VALUES (%s)"
)

type Updater struct {
	db *sql.DB
}

func NewUpdater(dataSourceName string) (lib.Updater, error) {

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	u := &Updater{
		db: db,
	}

	return u, nil
}

func (u *Updater) Set(name string) *lib.UpdateSet {
	set := &lib.UpdateSet{Name: name}
	return set
}

func (u *Updater) Update(sets ...*lib.UpdateSet) error {
	for _, set := range sets {

		//TODO(anarcher): Need refactoring
		var fs []string
		for i := 1; i <= len(set.Values); i++ {
			fs = append(fs, fmt.Sprintf("$%d", i))
		}
		q := fmt.Sprintf(sql_upsert, set.Name, strings.Join(set.Keys, ", "), strings.Join(fs, ","))

		values := make([]interface{}, len(set.Values))
		for i, v := range set.Values {
			values[i] = v
		}

		_, err := u.db.Exec(q, values...)
		if err != nil {
			return err
		}
	}

	return nil
}

func init() {
	Add("pgsql", NewUpdater)
}
