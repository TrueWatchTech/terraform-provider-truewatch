package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/TrueWatchTech/terraform-provider-truewatch/internal/resources/alert_policy"
	"github.com/TrueWatchTech/terraform-provider-truewatch/internal/resources/alert_policy_notice_date"
	"github.com/TrueWatchTech/terraform-provider-truewatch/internal/resources/blacklist"
	"github.com/TrueWatchTech/terraform-provider-truewatch/internal/resources/dashboard"
	"github.com/TrueWatchTech/terraform-provider-truewatch/internal/resources/membergroup"
	"github.com/TrueWatchTech/terraform-provider-truewatch/internal/resources/monitor"
	"github.com/TrueWatchTech/terraform-provider-truewatch/internal/resources/monitor_json"
	"github.com/TrueWatchTech/terraform-provider-truewatch/internal/resources/mute"
	"github.com/TrueWatchTech/terraform-provider-truewatch/internal/resources/notify_object"
	"github.com/TrueWatchTech/terraform-provider-truewatch/internal/resources/pipeline"
	"github.com/TrueWatchTech/terraform-provider-truewatch/internal/resources/role"
)

// Resources defines the resources implemented in the provider.
func (p *truewatchProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		alert_policy.NewAlertPolicyResource,
		alert_policy_notice_date.NewAlertPolicyNoticeDateResource,
		blacklist.NewBlackListResource,
		// custom_region.NewCustomRegionResource,
		dashboard.NewDashboardResource,
		membergroup.NewMemberGroupResource,
		monitor.NewMonitorResource,
		monitor_json.NewMonitorJsonResource,
		mute.NewMuteResource,
		notify_object.NewNotifyObjectResource,
		pipeline.NewPipelineResource,
		role.NewRoleResource,
		// slo.NewSloResource,
		// synthetics_test.NewSyntheticsTestResource,
	}
}
