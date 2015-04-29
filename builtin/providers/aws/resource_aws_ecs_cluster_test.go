package aws

import (
	"fmt"
	"testing"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/ecs"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccAwsEcsCluster(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckEcsClusterDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAwsEcsClusterConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEcsClusterExists("aws_ecs_cluster.foo"),
				),
			},
		},
	})
}

var testAccAwsEcsClusterConfig = fmt.Sprintf(`
resource "aws_ecs_cluster" "foo" {
    name = "api-%03d"
}
`, genRandInt())

func testAccCheckEcsClusterExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		conn := testAccProvider.Meta().(*AWSClient).ecsconn
		res, err := conn.DescribeClusters(&ecs.DescribeClustersInput{
			Clusters: []*string{aws.String(rs.Primary.ID)},
		})

		if err != nil {
			return err
		}

		if len(res.Clusters) == 0 {
			return fmt.Errorf("Not found: %s", n)
		}

		return nil
	}
}

func testAccCheckEcsClusterDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*AWSClient).ecsconn

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aws_ecs_cluster" {
			continue
		}
		fmt.Printf("XXXXX: %s", []*string{aws.String(rs.Primary.ID)})
		res, err := conn.DescribeClusters(&ecs.DescribeClustersInput{
			Clusters: []*string{aws.String(rs.Primary.ID)},
		})
		if err != nil {
			return nil
		}
		if len(res.Clusters) > 0 {
			return fmt.Errorf("Still exists")
		}
	}

	return nil
}
