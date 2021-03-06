// Copyright 2019 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package azure

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-08-01/network"
	"github.com/Azure/go-autorest/autorest"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/hashicorp/go-azure-helpers/authentication"
)

type VirtualNetworkGenerator struct {
	AzureService
}

func (g VirtualNetworkGenerator) createResources(ctx context.Context, iterator network.VirtualNetworkListResultIterator) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	for iterator.NotDone() {
		virtualNetwork := iterator.Value()
		resources = append(resources, terraformutils.NewSimpleResource(
			*virtualNetwork.ID,
			*virtualNetwork.Name,
			"azurerm_virtual_network",
			"azurerm",
			[]string{}))
		if err := iterator.NextWithContext(ctx); err != nil {
			log.Println(err)
			return resources, err
		}
	}
	return resources, nil
}

func (g *VirtualNetworkGenerator) InitResources() error {
	ctx := context.Background()
	virtualNetworkClient := network.NewVirtualNetworksClient(g.Args["config"].(authentication.Config).SubscriptionID)

	virtualNetworkClient.Authorizer = g.Args["authorizer"].(autorest.Authorizer)

	var (
		output network.VirtualNetworkListResultIterator
		err    error
	)

	if rg := g.Args["resource_group"].(string); rg != "" {
		output, err = virtualNetworkClient.ListComplete(ctx, rg)
	} else {
		output, err = virtualNetworkClient.ListAllComplete(ctx)
	}
	if err != nil {
		return err
	}
	g.Resources, err = g.createResources(ctx, output)
	return err
}
