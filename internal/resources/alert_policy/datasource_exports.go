package alert_policy

import "github.com/TrueWatchTech/terraform-provider-truewatch/internal/api"

type AlertOptModel = alertOptModel

func AlertOptFromContentForDataSource(content *api.AlertOptContent) *AlertOptModel {
	return alertOptFromContentForDataSource(content)
}
