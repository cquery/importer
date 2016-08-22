package worker

import (
	"github.com/cquery/importer/lib"

	"golang.org/x/net/context"

	"time"
)

type Worker struct {
	ctx      context.Context
	callers  map[string]lib.APICaller
	updaters map[string]lib.Updater
	interval time.Duration
}
