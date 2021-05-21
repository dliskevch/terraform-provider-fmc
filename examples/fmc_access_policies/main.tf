terraform {
  required_providers {
    fmc = {
      source = "hashicorp.com/edu/fmc"
    }
  }
}

provider "fmc" {
  fmc_username = "api"
  fmc_password = "CXsecurity!@34"
  fmc_host = "10.106.107.228"
  fmc_insecure_skip_verify = true
}

data "fmc_ips_policies" "ips_policy" {
    name = "Connectivity Over Security"
}

data "fmc_syslog_alerts" "syslog_alert" {
    name = "Testing Syslog"
}

resource "fmc_access_policies" "access_policy" {
    name = "Terraform Access Policy"
    # default_action = "block" # Cannot have block with base IPS policy
    default_action = "permit"
    default_action_base_intrusion_policy_id = data.fmc_ips_policies.ips_policy.id
    default_action_send_events_to_fmc = "true"
    default_action_log_end = "true"
    default_action_syslog_config_id = data.fmc_syslog_alerts.syslog_alert.id
}

output "new_fmc_access_policy" {
    value = fmc_access_policies.access_policy
}