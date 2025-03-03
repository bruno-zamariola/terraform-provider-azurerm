package volumesreplication

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VolumesDeleteReplicationOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// VolumesDeleteReplication ...
func (c VolumesReplicationClient) VolumesDeleteReplication(ctx context.Context, id VolumeId) (result VolumesDeleteReplicationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/deleteReplication", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	result.Poller, err = resourcemanager.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// VolumesDeleteReplicationThenPoll performs VolumesDeleteReplication then polls until it's completed
func (c VolumesReplicationClient) VolumesDeleteReplicationThenPoll(ctx context.Context, id VolumeId) error {
	result, err := c.VolumesDeleteReplication(ctx, id)
	if err != nil {
		return fmt.Errorf("performing VolumesDeleteReplication: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after VolumesDeleteReplication: %+v", err)
	}

	return nil
}
