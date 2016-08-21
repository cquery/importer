package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"

	"bytes"
	"fmt"
	"strings"
)

const (
	SQL_UPSERT_1 = "UPSERT INTO ec2 "
	SQL_UPSERT_2 = "(instance_id,image_id,instance_type,state_name) "
	SQL_UPSERT_3 = "VALUES ( "
	SQL_UPSERT_0 = ")\n"
	SQL_V        = `'%s'`
)

type EC2Resources []*ec2.Reservation

//TODO(anarcher): SQL(sqlbuilder *sqlbuilder.Builder) error { ...
func (e EC2Resources) SQL() string {
	var buf bytes.Buffer
	for _, r := range e {
		for _, i := range r.Instances {
			//els := []string{*i.InstanceId, *i.ImageId, *i.VpcId}
			els := []string{fmt.Sprintf(SQL_V, *i.InstanceId),
				fmt.Sprintf(SQL_V, *i.ImageId),
				fmt.Sprintf(SQL_V, *i.InstanceType),
				fmt.Sprintf(SQL_V, *i.State.Name)}
			buf.WriteString(SQL_UPSERT_1)
			buf.WriteString(SQL_UPSERT_2)
			buf.WriteString(SQL_UPSERT_3)
			buf.WriteString(strings.Join(els, ", "))
			buf.WriteString(SQL_UPSERT_0)
		}
	}
	return buf.String()
}

type EC2APICaller struct {
	service *ec2.EC2
}

func NewEC2APICaller(awsConfig *aws.Config) *EC2APICaller {

	sess, _ := session.NewSession()
	svc := ec2.New(sess, awsConfig)

	e := &EC2APICaller{
		service: svc,
	}
	return e
}

func (e *EC2APICaller) Run() (EC2Resources, error) {
	params := &ec2.DescribeInstancesInput{}
	resp, err := e.service.DescribeInstances(params)
	if err != nil {
		return nil, err
	}

	res := EC2Resources(resp.Reservations)
	return res, err
}
