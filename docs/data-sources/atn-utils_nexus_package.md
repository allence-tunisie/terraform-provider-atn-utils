# atn-utils_nexus_package

Use this data source to get a package from nexus.

## Example Usage

```hcl
data "atn-utils_nexus_package" "zip" {
  repository_url = "https://exemple-repository.com/repository/maven-releases/com/atn/webapp/0.0.0/webapp-0.0.0.zip"
  output_path = "webapp.zip"
}
```

## Argument Reference

The following arguments are supported:
* `repository_url` - (Required) The repository package url.
* `output_path` - (Required) The host path where the file will be created. 


## Attributes Reference

In addition to all arguments above, the package file will be exported in the output path.


