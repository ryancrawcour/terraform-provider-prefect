package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/tf6server"
	"github.com/prefecthq/terraform-provider-prefect/internal/provider"
)

const providerAddress = "registry.terraform.io/prefecthq/prefect"

func main() {
	provider := providerserver.NewProtocol6(&provider.Provider{})

	err := tf6server.Serve(providerAddress, provider)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to start starting plugin server: %s", err)
		os.Exit(1)
	}
}
