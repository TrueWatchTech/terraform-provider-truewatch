package provider

import (
	"context"

	"github.com/TrueWatchTech/terraform-provider-truewatch/internal/datasources/alert_policy"
	"github.com/TrueWatchTech/terraform-provider-truewatch/internal/datasources/alert_policy_notice_date"
	"github.com/TrueWatchTech/terraform-provider-truewatch/internal/datasources/members"
	"github.com/TrueWatchTech/terraform-provider-truewatch/internal/datasources/monitor"
	"github.com/TrueWatchTech/terraform-provider-truewatch/internal/datasources/monitors"
	"github.com/TrueWatchTech/terraform-provider-truewatch/internal/datasources/mute"
	"github.com/TrueWatchTech/terraform-provider-truewatch/internal/datasources/notify_object"
	"github.com/TrueWatchTech/terraform-provider-truewatch/internal/datasources/permissions"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// DataSources defines the data sources implemented in the provider.
func (p *truewatchProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		members.NewMembersDataSource,
		permissions.NewPermissionsDataSource,
		monitor.NewMonitorDataSource,
		monitors.NewMonitorsDataSource,
		notify_object.NewNotifyObjectDataSource,
		alert_policy.NewAlertPolicyDataSource,
		alert_policy_notice_date.NewAlertPolicyNoticeDateDataSource,
		mute.NewMuteDataSource,
		// default_region.NewDefaultRegionDataSource,
	}
}
