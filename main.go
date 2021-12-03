package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	// "github.com/hashicorp/terraform-plugin-sdk/terraform"
	// "github.com/richmondang/custom-tf/api/client"

	"github.com/richmondang/custom-tf/provider"
)

// func main() {
// 	plugin.Serve(&plugin.ServeOpts{
// 		ProviderFunc: func() terraform.ResourceProvider {
// 			return Provider()
// 		},
// 	})
// }

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	})
}
