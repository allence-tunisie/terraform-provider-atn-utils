terraform {
  required_providers {
    atn-utils = {
      source = "allence-tunisie/atn-utils"
      version = "0.0.2"
    }
  }
}

provider "atn-utils" {
  # Configuration options
}



data "atn-utils_gitlab_package" "zip" {
  repository_url = "https://gitlab.com/api/v4/projects/xxxxxxxx/packages/maven/com/atn/webapp-demo/1.1.1/webapp-demo-1.1.1-distribution.zip"
  access_token = "xxxxxxxxxxxxxxxxxx" // api token here
  output_path = "api-gitlab.zip"
}
data "atn-utils_nexus_package" "zip" {
  repository_url = "https://example.com/repository/maven-releases/com/atn/aws-webapp-nodejs/0.0.0/aws-webapp-nodejs-0.0.0.zip"
  output_path = "api"
  with_extract = true
}