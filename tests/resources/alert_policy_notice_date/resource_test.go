package AlertPolicyNoticeDate_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/TrueWatchTech/terraform-provider-truewatch/internal/provider"
)

func TestAccAlertPolicyNoticeDate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: provider.Config + `
resource "truewatch_alert_policy_notice_date" "demo" {
  name                     = "oac-alert-policy-notice-date-demo"
  skip_ref_check_on_delete = false

  notice_dates = [
    "2026/06/10",
    "2026/06/11",
  ]
}

data "truewatch_alert_policy_notice_date" "demo" {
  name = truewatch_alert_policy_notice_date.demo.name

  depends_on = [truewatch_alert_policy_notice_date.demo]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("truewatch_alert_policy_notice_date.demo", "name", "oac-alert-policy-notice-date-demo"),
					resource.TestCheckResourceAttr("truewatch_alert_policy_notice_date.demo", "notice_dates.0", "2026/06/10"),
					resource.TestCheckResourceAttr("truewatch_alert_policy_notice_date.demo", "notice_dates.1", "2026/06/11"),
					resource.TestCheckResourceAttr("truewatch_alert_policy_notice_date.demo", "skip_ref_check_on_delete", "false"),
					resource.TestCheckResourceAttrSet("data.truewatch_alert_policy_notice_date.demo", "uuid"),
				),
			},
		},
	})
}
