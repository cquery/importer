package main

import (
	"os"
	"time"

	"github.com/urfave/cli"
	"golang.org/x/net/context"

	"github.com/cquery/importer/lib"
	"github.com/cquery/importer/lib/worker"
)

func main() {
	app := cli.NewApp()
	app.Name = "cquery/importer"

	app.Commands = []cli.Command{
		{
			Name:    "daemon",
			Aliases: []string{"d"},
			Action:  DaemonAction,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "awsRegion",
					Value:  "",
					EnvVar: "AWS_REGION",
				},
				cli.StringFlag{
					Name:   "awsAccessKey",
					Value:  "",
					EnvVar: "AWS_ACCESS_KEY",
				},
				cli.StringFlag{
					Name:   "awsSecertKey",
					Value:  "",
					EnvVar: "AWS_SECERT_KEY",
				},
				cli.StringFlag{
					Name:   "pgsqlUser",
					Value:  "maxroach",
					EnvVar: "PGSQL_USER",
				},
				cli.StringFlag{
					Name:   "pgsqlAddr",
					Value:  "localhost:26257",
					EnvVar: "PGSQL_ADDR",
				},
				cli.StringFlag{
					Name:   "interval,i",
					Value:  "1m",
					Usage:  "The duration of fetching",
					EnvVar: "INTERVAL",
				},
			},
		},
	}

	app.Run(os.Args)
}

func DaemonAction(c *cli.Context) error {
	lib.Logger.Log("cquery/importer", "start", "version", Version, "gitcommit", GitCommit)

	if c.String("awsAccessKey") != "" {
		os.Setenv("AWS_ACCESS_KEY", c.String("awsAccessKey"))
	}
	if c.String("awsSecertKey") != "" {
		os.Setenv("AWS_SECERT_KEY", c.String("awsSecertKey"))
	}
	/*
		fmt.Println(os.Getenv("AWS_ACCESS_KEY"))
		fmt.Println(os.Getenv("AWS_SECERT_KEY"))
	*/

	ctx, cancelFunc := context.WithCancel(context.Background())

	interval, err := time.ParseDuration(c.String("interval"))
	if err != nil {
		lib.Logger.Log("err", err)
		cancelFunc()
		return err
	}

	_, err = worker.NewAWSPGSQLWorker(
		ctx,
		c.String("awsRegion"),
		c.String("pgsqlUser"),
		c.String("pgsqlAddr"),
		interval,
	)
	if err != nil {
		lib.Logger.Log("err", err, "worker", "aws-pgsql")
		return err
	}

	lib.WaitGroup.Wait()

	lib.Logger.Log("cquery/importer", "end", "version", Version, "gitcommit", GitCommit)

	return nil
}
