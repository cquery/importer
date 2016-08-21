package aws

import (
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
)

func Test_EC2APICaller(t *testing.T) {
	if os.Getenv("AWS_ACCESS_KEY") == "" {
		t.Log("AWS_ACCESS_KEY not found. this test is skipped")
		return
	}

	cfg := &aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	}
	caller := NewEC2APICaller(cfg)
	res, err := caller.Run()
	if err != nil {
		t.Error(err)
	}

	t.Log(res.SQL())

}
