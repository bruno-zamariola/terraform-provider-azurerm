package simgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SIMGroupClient struct {
	Client *resourcemanager.Client
}

func NewSIMGroupClientWithBaseURI(api environments.Api) (*SIMGroupClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "simgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SIMGroupClient: %+v", err)
	}

	return &SIMGroupClient{
		Client: client,
	}, nil
}
