package pgsql

import (
	"github.com/cquery/importer/lib"
)

type Creator func(dataSourceName string) (lib.Updater, error)

var Updaters = map[string]Creator{}

func Add(name string, creator Creator) {
	Updaters[name] = creator
}
