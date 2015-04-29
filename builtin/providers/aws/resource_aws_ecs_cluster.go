package aws

import (
	"fmt"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/ecs"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAwsEcsCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsEcsClusterCreate,
		Read:   resourceAwsEcsClusterRead,
		Update: nil,
		Delete: resourceAwsEcsClusterDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAwsEcsClusterCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).ecsconn
	req := &ecs.CreateClusterInput{
		ClusterName: aws.String(d.Get("name").(string)),
	}

	res, err := conn.CreateCluster(req)
	if err != nil {
		return fmt.Errorf("Error creating ecs cluster: %s", err)
	}

	d.SetId(*res.Cluster.ClusterName)

	return nil
}

func resourceAwsEcsClusterRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).ecsconn
	req := &ecs.DescribeClustersInput{
		Clusters: []*string{
			aws.String(d.Get("name").(string)),
		},
	}

	_, err := conn.DescribeClusters(req)
	if err != nil {
		return fmt.Errorf("Error reading ecs cluster: %s", err)
	}

	return nil
}

func resourceAwsEcsClusterDelete(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).ecsconn
	req := &ecs.DeleteClusterInput{
		Cluster: aws.String(d.Id()),
	}

	_, err := conn.DeleteCluster(req)
	if err != nil {
		return fmt.Errorf("Error deleting ecs cluster: %s", err)
	}

	return nil
}
