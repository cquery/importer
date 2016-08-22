package aws

import (
	"github.com/cquery/importer/lib"

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

	updater := &MockUpdater{
		updateFunc: func(sets ...*lib.UpdateSet) error {
			if len(sets) <= 0 {
				t.Error("updateset is zero")
			}
			return nil
		},
	}

	res.Update(updater)

}

type MockUpdater struct {
	updateFunc func(sets ...*lib.UpdateSet) error
}

func (u *MockUpdater) Set(name string) *lib.UpdateSet {
	set := &lib.UpdateSet{Name: name}
	return set
}

func (u *MockUpdater) Update(sets ...*lib.UpdateSet) error {
	return u.updateFunc(sets...)
}
