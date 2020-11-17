package atn_utils

import (
	"archive/zip"
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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
			"with_extract": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func dataSourcePackageRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var repositoryURL = data.Get("repository_url").(string)
	var outputPath = data.Get("output_path").(string)
	var token = data.Get("access_token").(string)
	var unzip = data.Get("with_extract").(bool)
	key := "PRIVATE-TOKEN"
	err := DownloadFile(outputPath, repositoryURL, key, token , unzip)
	if err != nil {
		return diag.Errorf("Gitlab download failure : %s", err)
	}
	return diags
}

func DownloadFile(filepath string, url string, tokenKey string, tokenValue string, unzip bool) error {

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
		return errors.New(" status " + res.Status)
	}
	// Create the file

	if unzip {
		zipSource := "tmp.zip"
		out, err := os.Create(zipSource)
		if err != nil {
			return err
		}
		defer out.Close()
		_, err = io.Copy(out, res.Body)
		if err != nil {
			return err
		}
		_, err = Unzip(zipSource, filepath)
		if err != nil {
			return err
		}

	} else {
		out, err := os.Create(filepath)
		if err != nil {
			return err
		}
		defer out.Close()
		_, err = io.Copy(out, res.Body)
		if err != nil {
			return err
		}
	}
	// Write the body to file
	return err

}

func Unzip(src string, dest string) ([]string, error) {

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()
		if err != nil {
			return filenames, err
		}

	}
	err = os.Remove("tmp.zip")
	if err != nil {
		return filenames, err
	}
	return filenames, nil
}
