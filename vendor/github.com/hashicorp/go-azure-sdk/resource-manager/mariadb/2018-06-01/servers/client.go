package servers

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServersClient struct {
	Client *resourcemanager.Client
}

func NewServersClientWithBaseURI(api environments.Api) (*ServersClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "servers", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServersClient: %+v", err)
	}

	return &ServersClient{
		Client: client,
	}, nil
}
