package aws

import (
	"github.com/cquery/importer/lib"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type EC2Resources []*ec2.Reservation

func (e EC2Resources) Update(updater lib.Updater) error {
	for _, r := range e {
		for _, i := range r.Instances {
			updateSet := updater.Set("ec2")
			updateSet.AddString("instance_id", *i.InstanceId)
			updateSet.AddString("image_id", *i.ImageId)
			updateSet.AddString("instance_type", *i.InstanceType)
			updateSet.AddString("state_name", *i.State.Name)
			if err := updater.Update(updateSet); err != nil {
				return err
			}
		}
	}
	return nil
}

type EC2APICaller struct {
	service *ec2.EC2
}

func NewEC2APICaller(awsConfig *aws.Config) lib.APICaller {

	sess, _ := session.NewSession()
	svc := ec2.New(sess, awsConfig)

	e := &EC2APICaller{
		service: svc,
	}
	return e
}

func (e *EC2APICaller) Call() (lib.Resources, error) {
	params := &ec2.DescribeInstancesInput{}
	resp, err := e.service.DescribeInstances(params)
	if err != nil {
		return nil, err
	}

	res := EC2Resources(resp.Reservations)
	return res, err
}

func init() {
	Add("ec2", NewEC2APICaller)
}
