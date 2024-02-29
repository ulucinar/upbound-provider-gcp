package vpcaccess

import (
	"crypto/rand"

	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	rand.Int(nil, nil)
	p.AddResourceConfigurator("google_vpc_access_connector", func(r *config.Resource) {
		r.References["network"] = config.Reference{
			Type: "github.com/upbound/provider-gcp/apis/compute/v1beta1.Network",
		}
		config.MarkAsRequired(r.TerraformResource, "region")
	})
}
