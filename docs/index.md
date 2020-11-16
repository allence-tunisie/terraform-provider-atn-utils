# atn-utils Provider

The atn-utils Provider

## Example Usage
```hcl
terraform {
  required_version = ">= 0.13"

  required_providers {
    atn-utils = {
      source = "allence-tunisie/atn-utils"
    }
  }
}

variable "nexus_package_url" {}
variable "gitlab_package_url" {}
variable "gitlab_api_token" {}

provider "atn-utils" {
  
}

data "atn-utils_gitlab_package" "zip" {
  repository_url = var.gitlab_package_url
  access_token = var.gitlab_api_token
  output_path = "package-gitlab.zip"
}
data "atn-utils_nexus_package" "zip" {
  repository_url = var.nexus_package_url
  output_path = "package-nexus.zip"
}
```

