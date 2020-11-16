package atn_utils

import (
	"context"
	"errors"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"io"
	"net/http"
	"os"
)

func dataSourcePackageGitlab() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePackageRead,
		Schema: map[string]*schema.Schema{
			"repository_url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"access_token": {
				Type:     schema.TypeString,
				Required: true,
			},
			"output_path": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourcePackageRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var repositoryURL = data.Get("repository_url").(string)
	var outputPath = data.Get("output_path").(string)
	var token = data.Get("access_token").(string)
	key := "PRIVATE-TOKEN"
		err := DownloadFile(outputPath, repositoryURL , key , token)
		if err != nil {
			return diag.Errorf("Gitlab download failure : %s" , err)
		}
	return diags
}

func DownloadFile(filepath string, url string , tokenKey string, tokenValue string) error {

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set(tokenKey, tokenValue)

	// Get the data
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.Status != "200 OK" {
		return errors.New(" status "+res.Status)
	}
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, res.Body)
	return err
}