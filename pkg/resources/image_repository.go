package resources

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/helpers"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var imageRepositorySchema = map[string]*schema.Schema{
	"name": {
		Type:        schema.TypeString,
		Required:    true,
		ForceNew:    true,
		Description: "Specifies the identifier for the image repository.",
	},
	"database": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Database for the image repository.",
	},
	"schema": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Schema for the image repository.",
	},
	"repository_url": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "URL for the image repository.",
	},
	"comment": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Comment for the image repository.",
	},
	"created_on": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Time image repository created.",
	},
	"owner": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Owner of the image repository.",
	},
}

func ImageRepository() *schema.Resource {
	return &schema.Resource{
		Create: CreateImageRepository,
		Read:   ReadImageRepository,
		Update: UpdateImageRepository,
		Delete: DeleteImageRepository,

		Schema: imageRepositorySchema,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// CreateImageRepository implements schema.CreateFunc.
func CreateImageRepository(d *schema.ResourceData, meta interface{}) error {
	db := meta.(*sql.DB)
	client := sdk.NewClientFromDB(db)
	ctx := context.Background()

	databaseName := d.Get("database").(string)
	schemaName := d.Get("schema").(string)
	name := d.Get("name").(string)
	objectIdentifier := sdk.NewSchemaObjectIdentifier(databaseName, schemaName, name)

	createRequest := sdk.NewCreateImageRepositoryRequest(objectIdentifier)

	if err := client.ImageRepositories.Create(ctx, createRequest); err != nil {
		return err
	}

	d.SetId(helpers.EncodeSnowflakeID(objectIdentifier))
	return ReadImageRepository(d, meta)
}

// ReadImageRepository implements schema.ReadFunc.
func ReadImageRepository(d *schema.ResourceData, meta interface{}) error {
	db := meta.(*sql.DB)
	client := sdk.NewClientFromDB(db)
	ctx := context.Background()

	id := helpers.DecodeSnowflakeID(d.Id()).(sdk.SchemaObjectIdentifier)

	i, err := client.ImageRepositories.ShowByID(ctx, id)
	if err != nil {
		return err
	}

	if err = d.Set("name", i.Name); err != nil {
		return err
	}
	if err = d.Set("database", i.DatabaseName); err != nil {
		return err
	}
	if err = d.Set("schema", i.SchemaName); err != nil {
		return err
	}
	if err = d.Set("repository_url", i.RepositoryUrl); err != nil {
		return err
	}
	if err = d.Set("comment", i.Comment); err != nil {
		return err
	}
	if err = d.Set("created_on", i.CreatedOn.Format("2006-01-02T16:04:05.000 -0700")); err != nil {
		return err
	}
	if err = d.Set("owner", i.Owner); err != nil {
		return err
	}

	return nil
}

// UpdateImageRepository implements schema.UpdateFunc.
func UpdateImageRepository(d *schema.ResourceData, meta interface{}) error {
	// NOTE: no alter docs available for image repositories, unclear what can be altered/set
	db := meta.(*sql.DB)
	_ = sdk.NewClientFromDB(db)
	_ = context.Background()

	_ = helpers.DecodeSnowflakeID(d.Id()).(sdk.SchemaObjectIdentifier)

	return ReadImageRepository(d, meta)
}

// DeleteImageRepository implements schema.DeleteFunc.
func DeleteImageRepository(d *schema.ResourceData, meta interface{}) error {
	db := meta.(*sql.DB)
	client := sdk.NewClientFromDB(db)
	ctx := context.Background()

	id := helpers.DecodeSnowflakeID(d.Id()).(sdk.SchemaObjectIdentifier)

	if err := client.ImageRepositories.Drop(ctx, sdk.NewDropImageRepositoryRequest(id)); err != nil {
		return fmt.Errorf("error deleting image repository %v err = %w", d.Id(), err)
	}

	d.SetId("")

	return nil
}
