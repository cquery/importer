package worker

import (
	"github.com/cquery/importer/lib"
	"github.com/cquery/importer/lib/aws"
	"github.com/cquery/importer/lib/pgsql"

	_aws "github.com/aws/aws-sdk-go/aws"
	"golang.org/x/net/context"

	"fmt"
	"time"
)

func NewAWSPGSQLWorker(ctx context.Context, awsRegion, pgsqlUser, pgsqlAddr string, interval time.Duration) (*Worker, error) {
	w := &Worker{
		ctx:      ctx,
		callers:  make(map[string]lib.APICaller),
		updaters: make(map[string]lib.Updater),
		interval: interval,
	}

	awsConfig := &_aws.Config{Region: _aws.String(awsRegion)}
	for n, c := range aws.APICallers {
		w.callers[n] = c(awsConfig)
	}

	dsString := fmt.Sprintf("postgresql://%s@%s/%s?sslmode=disable", pgsqlUser, pgsqlAddr, "aws")
	lib.Logger.Log("pgsql", dsString)

	for n, c := range pgsql.Updaters {
		u, err := c(dsString)
		if err != nil {
			return nil, err
		}
		w.updaters[n] = u
	}

	w.Run(time.Tick(w.interval))
	return w, nil
}

func (w *Worker) Run(ticker <-chan time.Time) {

	go func() {
		for {
			select {
			case <-w.ctx.Done():
				lib.WaitGroup.Done()
				return
			case <-ticker:
				for cn, c := range w.callers {
					r, err := c.Call()
					if err != nil {
						lib.Logger.Log("apicaller", cn, "err", err)
						continue
					}
					for un, u := range w.updaters {
						if err := r.Update(u); err != nil {
							lib.Logger.Log("updater", un, "apicaller", cn, "err", err)
						}
					}
				}
			}
		}
	}()

	lib.WaitGroup.Add(1)
}
