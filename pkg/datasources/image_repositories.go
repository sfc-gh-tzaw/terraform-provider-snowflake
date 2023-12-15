package datasources

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/snowflake"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var imageRepositoriesSchema = map[string]*schema.Schema{
	"image_repositories": {
		Type:        schema.TypeList,
		Computed:    true,
		Description: "The image repositories",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"database_name": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"schema_name": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"repository_url": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"owner": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"comment": {
					Type:     schema.TypeString,
					Computed: true,
				},
			},
		},
	},
}

func ImageRepositories() *schema.Resource {
	return &schema.Resource{
		Read:   ReadImageRepositories,
		Schema: imageRepositoriesSchema,
	}
}

func ReadImageRepositories(d *schema.ResourceData, meta interface{}) error {
	db := meta.(*sql.DB)
	client := sdk.NewClientFromDB(db)
	ctx := context.Background()

	account, err := snowflake.ReadCurrentAccount(db)
	if err != nil {
		d.SetId("")
		return nil
	}
	d.SetId(fmt.Sprintf("%s.%s", account.Account, account.Region))

	extractedImageRepositories, err := client.ImageRepositories.Show(ctx, sdk.NewShowImageRepositoryRequest())
	if err != nil {
		// If not found, mark resource to be removed from state file during apply or refresh
		log.Printf("[DEBUG] image repositories in id (%s) not found", d.Id())
		d.SetId("")
		return nil
	}

	imageRepositories := make([]map[string]any, len(extractedImageRepositories))
	for i, imageRepository := range extractedImageRepositories {
		imageRepositories[i] = map[string]any{
			"name":           imageRepository.Name,
			"database_name":  imageRepository.DatabaseName,
			"schema_name":    imageRepository.SchemaName,
			"repository_url": imageRepository.RepositoryUrl,
			"owner":          imageRepository.Owner,
			"comment":        imageRepository.Comment,
		}
	}

	return d.Set("image_repositories", imageRepositories)
}
