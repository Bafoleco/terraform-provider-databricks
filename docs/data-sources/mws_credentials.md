---
subcategory: "Deployment"
---
# databricks_mws_credentials Data Source

Lists all [databricks_mws_credentials](../resources/mws_credentials.md) in Databricks Account.

-> This data source can only be used with an account-level provider!

## Example Usage

Listing all credentials in Databricks Account

```hcl
provider "databricks" {
  // other configuration
  account_id = "<databricks account id>"
}

data "databricks_mws_credentials" "all" {}

output "all_mws_credentials" {
  value = data.databricks_mws_credentials.all.ids
}
```

## Attribute Reference

-> This resource has an evolving interface, which may change in future versions of the provider.

This data source exports the following attributes:

* `ids` - name-to-id map for all of the credentials in the account

## Related Resources

The following resources are used in the same context:

* [Provisioning Databricks on AWS](../guides/aws-workspace.md) guide.
* [databricks_mws_customer_managed_keys](../resources/mws_customer_managed_keys.md) to configure KMS keys for new workspaces within AWS.
* [databricks_mws_log_delivery](../resources/mws_log_delivery.md) to configure delivery of [billable usage logs](https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html) and [audit logs](https://docs.databricks.com/administration-guide/account-settings/audit-logs.html).
* [databricks_mws_networks](../resources/mws_networks.md) to [configure VPC](https://docs.databricks.com/administration-guide/cloud-configurations/aws/customer-managed-vpc.html) & subnets for new workspaces within AWS.
* [databricks_mws_storage_configurations](../resources/mws_storage_configurations.md) to configure root bucket new workspaces within AWS.
* [databricks_mws_workspaces](../resources/mws_workspaces.md) to set up [AWS and GCP workspaces](https://docs.databricks.com/getting-started/overview.html#e2-architecture-1).
