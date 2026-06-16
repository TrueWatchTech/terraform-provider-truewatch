
terraform {
  required_version = ">= 1.0"

  required_providers {
    truewatch = {
      source = "TrueWatchTech/truewatch"
    }
  }
}

provider "truewatch" {
  region = "singapore"
}
