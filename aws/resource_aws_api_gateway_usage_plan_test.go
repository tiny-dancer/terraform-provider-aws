package aws

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccAWSAPIGatewayUsagePlan_basic(t *testing.T) {
	var conf apigateway.UsagePlan
	name := acctest.RandString(10)
	updatedName := acctest.RandString(10)
	resourceName := "aws_api_gateway_usage_plan.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSAPIGatewayUsagePlanDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSApiGatewayUsagePlanBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSAPIGatewayUsagePlanExists(resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", ""),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccAWSApiGatewayUsagePlanBasicUpdatedConfig(updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", ""),
				),
			},
			{
				Config: testAccAWSApiGatewayUsagePlanBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSAPIGatewayUsagePlanExists(resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", ""),
				),
			},
		},
	})
}

func TestAccAWSAPIGatewayUsagePlan_description(t *testing.T) {
	var conf apigateway.UsagePlan
	name := acctest.RandString(10)
	resourceName := "aws_api_gateway_usage_plan.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSAPIGatewayUsagePlanDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSApiGatewayUsagePlanBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSAPIGatewayUsagePlanExists(resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "description", ""),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccAWSApiGatewayUsagePlanDescriptionConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSAPIGatewayUsagePlanExists(resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "description", "This is a description"),
				),
			},
			{
				Config: testAccAWSApiGatewayUsagePlanDescriptionUpdatedConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", "This is a new description"),
				),
			},
			{
				Config: testAccAWSApiGatewayUsagePlanDescriptionConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSAPIGatewayUsagePlanExists(resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "description", "This is a description"),
				),
			},
			{
				Config: testAccAWSApiGatewayUsagePlanBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSAPIGatewayUsagePlanExists(resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "description", ""),
				),
			},
		},
	})
}

func TestAccAWSAPIGatewayUsagePlan_productCode(t *testing.T) {
	var conf apigateway.UsagePlan
	name := acctest.RandString(10)
	resourceName := "aws_api_gateway_usage_plan.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSAPIGatewayUsagePlanDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSApiGatewayUsagePlanBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSAPIGatewayUsagePlanExists(resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "product_code", ""),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccAWSApiGatewayUsagePlanProductCodeConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSAPIGatewayUsagePlanExists(resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "product_code", "MYCODE"),
				),
			},
			{
				Config: testAccAWSApiGatewayUsagePlanProductCodeUpdatedConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "product_code", "MYCODE2"),
				),
			},
			{
				Config: testAccAWSApiGatewayUsagePlanProductCodeConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSAPIGatewayUsagePlanExists(resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "product_code", "MYCODE"),
				),
			},
			{
				Config: testAccAWSApiGatewayUsagePlanBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSAPIGatewayUsagePlanExists(resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "product_code", ""),
				),
			},
		},
	})
}

func TestAccAWSAPIGatewayUsagePlan_throttling(t *testing.T) {
	var conf apigateway.UsagePlan
	name := acctest.RandString(10)
	resourceName := "aws_api_gateway_usage_plan.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSAPIGatewayUsagePlanDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSApiGatewayUsagePlanBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSAPIGatewayUsagePlanExists(resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckNoResourceAttr(resourceName, "throttle_settings"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccAWSApiGatewayUsagePlanThrottlingConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSAPIGatewayUsagePlanExists(resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "throttle_settings.4173790118.burst_limit", "2"),
					resource.TestCheckResourceAttr(resourceName, "throttle_settings.4173790118.rate_limit", "5"),
				),
			},
			{
				Config: testAccAWSApiGatewayUsagePlanThrottlingModifiedConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSAPIGatewayUsagePlanExists(resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "throttle_settings.1779463053.burst_limit", "3"),
					resource.TestCheckResourceAttr(resourceName, "throttle_settings.1779463053.rate_limit", "6"),
				),
			},
			{
				Config: testAccAWSApiGatewayUsagePlanBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSAPIGatewayUsagePlanExists(resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckNoResourceAttr(resourceName, "throttle_settings"),
				),
			},
		},
	})
}

// https://github.com/terraform-providers/terraform-provider-aws/issues/2057
func TestAccAWSAPIGatewayUsagePlan_throttlingInitialRateLimit(t *testing.T) {
	var conf apigateway.UsagePlan
	name := acctest.RandString(10)
	resourceName := "aws_api_gateway_usage_plan.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSAPIGatewayUsagePlanDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSApiGatewayUsagePlanThrottlingConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSAPIGatewayUsagePlanExists(resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "throttle_settings.4173790118.rate_limit", "5"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAWSAPIGatewayUsagePlan_quota(t *testing.T) {
	var conf apigateway.UsagePlan
	name := acctest.RandString(10)
	resourceName := "aws_api_gateway_usage_plan.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSAPIGatewayUsagePlanDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSApiGatewayUsagePlanBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSAPIGatewayUsagePlanExists(resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckNoResourceAttr(resourceName, "quota_settings"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccAWSApiGatewayUsagePlanQuotaConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSAPIGatewayUsagePlanExists(resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "quota_settings.1956747625.limit", "100"),
					resource.TestCheckResourceAttr(resourceName, "quota_settings.1956747625.offset", "6"),
					resource.TestCheckResourceAttr(resourceName, "quota_settings.1956747625.period", "WEEK"),
				),
			},
			{
				Config: testAccAWSApiGatewayUsagePlanQuotaModifiedConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSAPIGatewayUsagePlanExists(resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "quota_settings.3909168194.limit", "200"),
					resource.TestCheckResourceAttr(resourceName, "quota_settings.3909168194.offset", "20"),
					resource.TestCheckResourceAttr(resourceName, "quota_settings.3909168194.period", "MONTH"),
				),
			},
			{
				Config: testAccAWSApiGatewayUsagePlanBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSAPIGatewayUsagePlanExists(resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckNoResourceAttr(resourceName, "quota_settings"),
				),
			},
		},
	})
}

func TestAccAWSAPIGatewayUsagePlan_apiStages(t *testing.T) {
	var conf apigateway.UsagePlan
	name := acctest.RandString(10)
	resourceName := "aws_api_gateway_usage_plan.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSAPIGatewayUsagePlanDestroy,
		Steps: []resource.TestStep{
			// Create UsagePlan WITH Stages as the API calls are different
			// when creating or updating.
			{
				Config: testAccAWSApiGatewayUsagePlanApiStagesConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSAPIGatewayUsagePlanExists(resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "api_stages.0.stage", "test"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Handle api stages removal
			{
				Config: testAccAWSApiGatewayUsagePlanBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSAPIGatewayUsagePlanExists(resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckNoResourceAttr(resourceName, "api_stages"),
				),
			},
			// Handle api stages additions
			{
				Config: testAccAWSApiGatewayUsagePlanApiStagesConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSAPIGatewayUsagePlanExists(resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "api_stages.0.stage", "test"),
				),
			},
			// Handle api stages updates
			{
				Config: testAccAWSApiGatewayUsagePlanApiStagesModifiedConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSAPIGatewayUsagePlanExists(resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "api_stages.0.stage", "foo"),
				),
			},
			{
				Config: testAccAWSApiGatewayUsagePlanBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSAPIGatewayUsagePlanExists(resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckNoResourceAttr(resourceName, "api_stages"),
				),
			},
		},
	})
}

func testAccCheckAWSAPIGatewayUsagePlanExists(n string, res *apigateway.UsagePlan) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No API Gateway Usage Plan ID is set")
		}

		conn := testAccProvider.Meta().(*AWSClient).apigateway

		req := &apigateway.GetUsagePlanInput{
			UsagePlanId: aws.String(rs.Primary.ID),
		}
		up, err := conn.GetUsagePlan(req)
		if err != nil {
			return err
		}

		if *up.Id != rs.Primary.ID {
			return fmt.Errorf("APIGateway Usage Plan not found")
		}

		*res = *up

		return nil
	}
}

func testAccCheckAWSAPIGatewayUsagePlanDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*AWSClient).apigateway

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aws_api_gateway_usage_plan" {
			continue
		}

		req := &apigateway.GetUsagePlanInput{
			UsagePlanId: aws.String(s.RootModule().Resources["aws_api_gateway_rest_api.test"].Primary.ID),
		}
		describe, err := conn.GetUsagePlan(req)

		if err == nil {
			if describe.Id != nil && *describe.Id == rs.Primary.ID {
				return fmt.Errorf("API Gateway Usage Plan still exists")
			}
		}

		aws2err, ok := err.(awserr.Error)
		if !ok {
			return err
		}
		if aws2err.Code() != "NotFoundException" {
			return err
		}

		return nil
	}

	return nil
}

const testAccAWSAPIGatewayUsagePlanConfig = `
resource "aws_api_gateway_rest_api" "test" {
  name = "test"
}

resource "aws_api_gateway_resource" "test" {
  rest_api_id = "${aws_api_gateway_rest_api.test.id}"
  parent_id = "${aws_api_gateway_rest_api.test.root_resource_id}"
  path_part = "test"
}

resource "aws_api_gateway_method" "test" {
  rest_api_id = "${aws_api_gateway_rest_api.test.id}"
  resource_id = "${aws_api_gateway_resource.test.id}"
  http_method = "GET"
  authorization = "NONE"
}

resource "aws_api_gateway_method_response" "error" {
  rest_api_id = "${aws_api_gateway_rest_api.test.id}"
  resource_id = "${aws_api_gateway_resource.test.id}"
  http_method = "${aws_api_gateway_method.test.http_method}"
  status_code = "400"
}

resource "aws_api_gateway_integration" "test" {
  rest_api_id = "${aws_api_gateway_rest_api.test.id}"
  resource_id = "${aws_api_gateway_resource.test.id}"
  http_method = "${aws_api_gateway_method.test.http_method}"

  type = "HTTP"
  uri = "https://www.google.de"
  integration_http_method = "GET"
}

resource "aws_api_gateway_integration_response" "test" {
  rest_api_id = "${aws_api_gateway_rest_api.test.id}"
  resource_id = "${aws_api_gateway_resource.test.id}"
  http_method = "${aws_api_gateway_integration.test.http_method}"
  status_code = "${aws_api_gateway_method_response.error.status_code}"
}

resource "aws_api_gateway_deployment" "test" {
  depends_on = ["aws_api_gateway_integration.test"]

  rest_api_id = "${aws_api_gateway_rest_api.test.id}"
  stage_name = "test"
  description = "This is a test"

  variables = {
    "a" = "2"
  }
}

resource "aws_api_gateway_deployment" "foo" {
  depends_on = ["aws_api_gateway_deployment.test", "aws_api_gateway_integration.test"]

  rest_api_id = "${aws_api_gateway_rest_api.test.id}"
  stage_name = "foo"
  description = "This is a prod stage"
}
`

func testAccAWSApiGatewayUsagePlanBasicConfig(rName string) string {
	return fmt.Sprintf(testAccAWSAPIGatewayUsagePlanConfig+`
resource "aws_api_gateway_usage_plan" "test" {
  name = "%s"
}
`, rName)
}

func testAccAWSApiGatewayUsagePlanDescriptionConfig(rName string) string {
	return fmt.Sprintf(testAccAWSAPIGatewayUsagePlanConfig+`
resource "aws_api_gateway_usage_plan" "test" {
  name        = "%s"
  description = "This is a description"
}
`, rName)
}

func testAccAWSApiGatewayUsagePlanDescriptionUpdatedConfig(rName string) string {
	return fmt.Sprintf(testAccAWSAPIGatewayUsagePlanConfig+`
resource "aws_api_gateway_usage_plan" "test" {
  name        = "%s"
  description = "This is a new description"
}
`, rName)
}

func testAccAWSApiGatewayUsagePlanProductCodeConfig(rName string) string {
	return fmt.Sprintf(testAccAWSAPIGatewayUsagePlanConfig+`
resource "aws_api_gateway_usage_plan" "test" {
  name         = "%s"
  product_code = "MYCODE"
}
`, rName)
}

func testAccAWSApiGatewayUsagePlanProductCodeUpdatedConfig(rName string) string {
	return fmt.Sprintf(testAccAWSAPIGatewayUsagePlanConfig+`
resource "aws_api_gateway_usage_plan" "test" {
  name         = "%s"
  product_code = "MYCODE2"
}
`, rName)
}

func testAccAWSApiGatewayUsagePlanBasicUpdatedConfig(rName string) string {
	return fmt.Sprintf(testAccAWSAPIGatewayUsagePlanConfig+`
resource "aws_api_gateway_usage_plan" "test" {
  name = "%s"
}
`, rName)
}

func testAccAWSApiGatewayUsagePlanThrottlingConfig(rName string) string {
	return fmt.Sprintf(testAccAWSAPIGatewayUsagePlanConfig+`
resource "aws_api_gateway_usage_plan" "test" {
  name        = "%s"

  throttle_settings {
    burst_limit = 2
    rate_limit  = 5
  }
}
`, rName)
}

func testAccAWSApiGatewayUsagePlanThrottlingModifiedConfig(rName string) string {
	return fmt.Sprintf(testAccAWSAPIGatewayUsagePlanConfig+`
resource "aws_api_gateway_usage_plan" "test" {
  name        = "%s"

  throttle_settings {
    burst_limit = 3
    rate_limit  = 6
  }
}
`, rName)
}

func testAccAWSApiGatewayUsagePlanQuotaConfig(rName string) string {
	return fmt.Sprintf(testAccAWSAPIGatewayUsagePlanConfig+`
resource "aws_api_gateway_usage_plan" "test" {
  name        = "%s"

  quota_settings {
    limit  = 100
    offset = 6
    period = "WEEK"
  }
}
`, rName)
}

func testAccAWSApiGatewayUsagePlanQuotaModifiedConfig(rName string) string {
	return fmt.Sprintf(testAccAWSAPIGatewayUsagePlanConfig+`
resource "aws_api_gateway_usage_plan" "test" {
  name        = "%s"

  quota_settings {
    limit  = 200
    offset = 20
    period = "MONTH"
  }
}
`, rName)
}

func testAccAWSApiGatewayUsagePlanApiStagesConfig(rName string) string {
	return fmt.Sprintf(testAccAWSAPIGatewayUsagePlanConfig+`
resource "aws_api_gateway_usage_plan" "test" {
  name        = "%s"

  api_stages {
    api_id = "${aws_api_gateway_rest_api.test.id}"
    stage  = "${aws_api_gateway_deployment.test.stage_name}"
  }
}
`, rName)
}

func testAccAWSApiGatewayUsagePlanApiStagesModifiedConfig(rName string) string {
	return fmt.Sprintf(testAccAWSAPIGatewayUsagePlanConfig+`
resource "aws_api_gateway_usage_plan" "test" {
  name        = "%s"

  api_stages {
    api_id = "${aws_api_gateway_rest_api.test.id}"
    stage  = "${aws_api_gateway_deployment.foo.stage_name}"
  }
}
`, rName)
}
