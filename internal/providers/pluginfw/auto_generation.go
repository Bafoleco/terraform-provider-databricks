// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package pluginfw

import (
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/alert_v2"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/budget_policy"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/database_instance"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// List of resources that are auto generated based on service OpenAPI specs.
var autoGeneratedResources = []func() resource.Resource{
	alert_v2.ResourceAlertV2,
	budget_policy.ResourceBudgetPolicy,
	database_instance.ResourceDatabaseInstance,
}

// List of data sources that are auto generated based on service OpenAPI specs.
var autoGeneratedDataSources = []func() datasource.DataSource{
	alert_v2.DataSourceAlertV2,
	budget_policy.DataSourceBudgetPolicy,
	database_instance.DataSourceDatabaseInstance,
	alert_v2.DataSourceAlertsV2,
	budget_policy.DataSourceBudgetPolicies,
	database_instance.DataSourceDatabaseInstances,
}
