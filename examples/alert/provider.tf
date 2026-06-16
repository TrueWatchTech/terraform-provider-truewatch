terraform {
  required_version = ">= 1.0"

  required_providers {
    guance = {
      source = "TrueWatchTech/truewatch"
    }
  }
}

provider "guance" {
  # You can set your API key here or use the TRUEWATCH_ACCESS_TOKEN environment variable.
  # access_token = "your-api-key"
}
