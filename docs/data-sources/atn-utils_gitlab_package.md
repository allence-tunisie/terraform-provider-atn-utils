# atn-utils_gitlab_package

Use this data source to get a package from gitlab.

## Example Usage

```hcl

data "atn-utils_gitlab_package" "zip" {
  repository_url = "https://gitlab.com/api/v4/projects/xxxxxxx/xxx/xxx/xx/xx/webapp/1.1.1/webapp-1.1.1-distribution.zip"
  access_token = "" // here you can put your api token
  output_path = "webapp.zip"
}
```

## Argument Reference

The following arguments are supported:
* `repository_url` - (Required) The repository package url.
* `access_token` - (Required) The api read token.
* `output_path` - (Required) The host path where the file will be created. 


## Attributes Reference

In addition to all arguments above, the package file will be exported in the output path.


