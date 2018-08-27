package main

import (
	"github.com/pulumi/pulumi-statuscake"
	"github.com/pulumi/pulumi-statuscake/pkg/version"
	"github.com/pulumi/pulumi-terraform/pkg/tfgen"
)

func main() {
	tfgen.Main("statuscake", version.Version, statuscake.Provider())
}
