package atn_utils

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"atn-utils_gitlab_package":    dataSourcePackageGitlab(),
			"atn-utils_nexus_package":    dataSourcePackageNexus(),
		},
	}
}

