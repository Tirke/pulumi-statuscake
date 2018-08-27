package statuscake

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pulumi/pulumi-terraform/pkg/tfbridge"
	"github.com/terraform-providers/terraform-provider-statuscake/statuscake"
)

func Provider() tfbridge.ProviderInfo {
	p := statuscake.Provider().(*schema.Provider)
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "statuscake",
		Description: "A Pulumi package for creating StatusCake tests",
		Keywords:    []string{"pulumi", "statuscake"},
		License:     "MIT",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/Tirke/pulumi-statuscake",
	}

	return prov
}
