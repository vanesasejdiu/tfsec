package eks

import (
	"github.com/aquasecurity/defsec/rules/aws/eks"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
	"github.com/aquasecurity/tfsec/pkg/rule"
)

func init() {
	scanner.RegisterCheckRule(rule.Rule{
		LegacyID: "AWS068",
		BadExample: []string{`
 resource "aws_eks_cluster" "bad_example" {
     // other config 
 
     name = "bad_example_cluster"
     role_arn = var.cluster_arn
     vpc_config {
         endpoint_public_access = true
     }
 }
 `},
		GoodExample: []string{`
 resource "aws_eks_cluster" "good_example" {
     // other config 
 
     name = "good_example_cluster"
     role_arn = var.cluster_arn
     vpc_config {
         endpoint_public_access = true
         public_access_cidrs = ["10.2.0.0/8"]
     }
 }
 `},
		Links: []string{
			"https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/eks_cluster#vpc_config",
		},
		RequiredTypes:  []string{"resource"},
		RequiredLabels: []string{"aws_eks_cluster"},
		Base:           eks.CheckNoPublicClusterAccessToCidr,
	})
}
