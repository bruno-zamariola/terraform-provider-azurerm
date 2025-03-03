---
subcategory: "Cognitive Services"
layout: "azurerm"
page_title: "Azure Resource Manager: azurerm_cognitive_deployment"
description: |-
  Manages a Cognitive Services Account Deployment.
---

# azurerm_cognitive_deployment

Manages a Cognitive Services Account Deployment.

## Example Usage

```hcl
resource "azurerm_resource_group" "example" {
  name     = "example-resources"
  location = "West Europe"
}

resource "azurerm_cognitive_account" "example" {
  name                = "example-ca"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  kind                = "OpenAI"
  sku_name            = "S0"
}

resource "azurerm_cognitive_deployment" "example" {
  name                 = "example-cd"
  cognitive_account_id = azurerm_cognitive_account.example.id
  model {
    format  = "OpenAI"
    name    = "text-curie-001"
    version = "1"
  }

  scale {
    type = "Standard"
  }
}

```

## Arguments Reference

The following arguments are supported:

* `name` - (Required) The name of the Cognitive Services Account Deployment. Changing this forces a new resource to be created.

* `cognitive_account_id` - (Required) The ID of the Cognitive Services Account. Changing this forces a new resource to be created.

* `model` - (Required) A `model` block as defined below. Changing this forces a new resource to be created.

* `scale` - (Required) A `scale` block as defined below. Changing this forces a new resource to be created.

* `rai_policy_name` - (Optional) The name of RAI policy. Changing this forces a new resource to be created.

---

A `model` block supports the following:

* `format` - (Required) The format of the Cognitive Services Account Deployment model. Changing this forces a new resource to be created. Possible value is `OpenAI`.

* `name` - (Required) The name of the Cognitive Services Account Deployment model. Changing this forces a new resource to be created.

* `version` - (Required) The version of Cognitive Services Account Deployment model.

---

A `scale` block supports the following:

* `type` - (Required) The name of the SKU. Ex - `Standard` or `P3`. It is typically a letter+number code. Changing this forces a new resource to be created.

* `tier` - (Optional) Possible values are `Free`, `Basic`, `Standard`, `Premium`, `Enterprise`. Changing this forces a new resource to be created.

* `size` - (Optional) The SKU size. When the name field is the combination of tier and some other value, this would be the standalone code. Changing this forces a new resource to be created.

* `family` - (Optional) If the service has different generations of hardware, for the same SKU, then that can be captured here. Changing this forces a new resource to be created.

* `capacity` - (Optional) If the SKU supports scale out/in then the capacity integer should be included. If scale out/in is not possible for the resource this may be omitted. Possible values are between `1` and `10000` and the default value is `1`. Changing this forces a new resource to be created.

## Attributes Reference

In addition to the Arguments listed above - the following Attributes are exported:

* `id` - The ID of the Deployment for Azure Cognitive Services Account.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 30 minutes) Used when creating the Cognitive Services Account Deployment.
* `read` - (Defaults to 5 minutes) Used when retrieving the Cognitive Services Account Deployment.
* `update` - (Defaults to 30 minutes) Used when updating the Cognitive Services Account Deployment.
* `delete` - (Defaults to 30 minutes) Used when deleting the Cognitive Services Account Deployment.

## Import

Cognitive Services Account Deployment can be imported using the `resource id`, e.g.

```shell
terraform import azurerm_cognitive_deployment.example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.CognitiveServices/accounts/account1/deployments/deployment1
```
