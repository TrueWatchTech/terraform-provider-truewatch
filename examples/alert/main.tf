resource "truewatch_notify_object" "example" {
  type = "simpleHTTPRequest"
  name = "${var.name_prefix}-notify-object"

  opt_set = jsonencode({
    url = var.webhook_url
    headersConfig = {
      isOpen = false
      items  = []
    }
  })

  open_permission_set = false
}

resource "truewatch_alert_policy_notice_date" "example" {
  name                     = "${var.name_prefix}-notice-date"
  skip_ref_check_on_delete = false

  notice_dates = var.notice_dates
}

resource "truewatch_alert_policy" "example" {
  name          = "${var.name_prefix}-alert-policy"
  desc          = "Alert policy managed by Terraform example"
  rule_timezone = var.timezone

  alert_opt = {
    alert_type                      = "status"
    agg_type                        = "byFields"
    agg_interval                    = 60
    agg_fields                      = ["df_monitor_checker_id"]
    agg_labels                      = ["service"]
    agg_send_first                  = true
    ignore_ok                       = true
    silent_timeout                  = 300
    silent_timeout_by_status_enable = true

    silent_timeout_by_status = [{
      status         = "critical,error"
      silent_timeout = 600
    }]

    alert_target = [{
      name              = "Business hours"
      custom_date_uuids = [truewatch_alert_policy_notice_date.example.uuid]
      custom_start_time = "09:00:00"
      custom_duration   = 28800

      targets = [{
        to            = [truewatch_notify_object.example.uuid]
        status        = "critical,error,warning"
        filter_string = var.alert_filter

        upgrade_targets = [{
          to       = [truewatch_notify_object.example.uuid]
          duration = 900
        }]
      }]
    }]
  }
}

resource "truewatch_mute" "maintenance" {
  name        = "${var.name_prefix}-maintenance-mute"
  description = "One-time maintenance mute managed by Terraform example"
  type        = "alertPolicy"
  timezone    = var.timezone

  mute_ranges = [{
    name              = truewatch_alert_policy.example.name
    alert_policy_uuid = truewatch_alert_policy.example.uuid
  }]

  repeat_time_set = 0
  start_time      = var.mute_start_time
  end_time        = var.mute_end_time

  notify_time_str = var.mute_notify_time
  notify_message  = "Terraform alert example maintenance window starts soon."

  notify_targets = [{
    type = "notifyObject"
    to   = [truewatch_notify_object.example.uuid]
  }]
}

data "truewatch_notify_object" "example" {
  name = truewatch_notify_object.example.name

  depends_on = [truewatch_notify_object.example]
}

data "truewatch_alert_policy_notice_date" "example" {
  name = truewatch_alert_policy_notice_date.example.name

  depends_on = [truewatch_alert_policy_notice_date.example]
}

data "truewatch_alert_policy" "example" {
  name = truewatch_alert_policy.example.name

  depends_on = [truewatch_alert_policy.example]
}

data "truewatch_mute" "maintenance" {
  name = truewatch_mute.maintenance.name

  depends_on = [truewatch_mute.maintenance]
}

