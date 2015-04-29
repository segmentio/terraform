package aws

import (
	// "github.com/awslabs/aws-sdk-go/aws"
	// "github.com/awslabs/aws-sdk-go/service/ecs"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAwsEcsService() *schema.Resource {
	return &schema.Resource{
		Create: nil,
		Read:   nil,
		Update: nil,
		Delete: nil,

		Schema: map[string]*schema.Schema{
			"service_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"cluster": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"task_definition": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"client_token": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"desired_count": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},

			"load_balancers": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"load_balancer_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"container_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"container_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
		},
	}
}
