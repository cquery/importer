package aws

import (
	"github.com/cquery/importer/lib"

	"github.com/aws/aws-sdk-go/aws"
)

type Creator func(cfg *aws.Config) lib.APICaller

var APICallers = map[string]Creator{}

func Add(name string, creator Creator) {
	APICallers[name] = creator
}
