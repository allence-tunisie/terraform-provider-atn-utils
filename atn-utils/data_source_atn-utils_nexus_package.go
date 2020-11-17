package atn_utils

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePackageNexus() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePackageNexusRead,
		Schema: map[string]*schema.Schema{
			"repository_url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"output_path": {
				Type:     schema.TypeString,
				Required: true,
			},
			"with_extract": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

		},
	}
}

func dataSourcePackageNexusRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var repositoryURL = data.Get("repository_url").(string)
	var outputPath = data.Get("output_path").(string)
	var unzip = data.Get("with_extract").(bool)
	key := "PRIVATE-TOKEN"
	err := DownloadFile(outputPath, repositoryURL , key , "token" , unzip)
	if err != nil {
		return diag.Errorf("Nexus download failure : %s" , err)
	}
	return diags
}
