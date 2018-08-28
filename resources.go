package statuscake

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pulumi/pulumi-terraform/pkg/tfbridge"
	"github.com/pulumi/pulumi/pkg/tokens"
	"github.com/terraform-providers/terraform-provider-statuscake/statuscake"
	"unicode"
)

const (
	statuscakePkg = "statuscake"
	statuscakeMod = "index"
)

// statusCakeMember manufactures a type token for the Digital Ocean package and the given module and type.
func statusCakeMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(statuscakePkg + ":" + mod + ":" + mem)
}

// statusCakeType manufactures a type token for the Digital Ocean package and the given module and type.
func statusCakeType(mod string, typ string) tokens.Type {
	return tokens.Type(statusCakeMember(mod, typ))
}

func statusCakeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return statusCakeType(mod+"/"+fn, res)
}

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
		Resources: map[string]*tfbridge.ResourceInfo{
			"statuscake_test": {Tok: statusCakeResource(statuscakeMod, "Test")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi":    "^0.15.0",
				"builtin-modules":   "3.0.0",
				"read-package-tree": "^5.2.1",
				"resolve":           "^1.8.1",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.9.2",
			},
		},
	}

	return prov
}
