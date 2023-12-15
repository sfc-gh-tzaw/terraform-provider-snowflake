package sdk

import (
	"context"
	"database/sql"
	"time"
)

type ImageRepositories interface {
	Create(ctx context.Context, request *CreateImageRepositoryRequest) error
	Drop(ctx context.Context, request *DropImageRepositoryRequest) error
	Show(ctx context.Context, request *ShowImageRepositoryRequest) ([]ImageRepository, error)
	ShowByID(ctx context.Context, id SchemaObjectIdentifier) (*ImageRepository, error)
}

// CreateImageRepositoryOptions is based on https://docs.snowflake.com/LIMITEDACCESS/snowpark-containers/reference/image-repository#create-image-repository.
type CreateImageRepositoryOptions struct {
	create          bool                   `ddl:"static" sql:"CREATE"`
	OrReplace       *bool                  `ddl:"keyword" sql:"OR REPLACE"`
	imageRepository bool                   `ddl:"static" sql:"IMAGE REPOSITORY"`
	IfNotExists     *bool                  `ddl:"keyword" sql:"IF NOT EXISTS"`
	name            SchemaObjectIdentifier `ddl:"identifier"`
}

// DropImageRepositoryOptions is based on https://docs.snowflake.com/LIMITEDACCESS/snowpark-containers/reference/image-repository#drop-image-repository.
type DropImageRepositoryOptions struct {
	drop            bool                   `ddl:"static" sql:"DROP"`
	imageRepository bool                   `ddl:"static" sql:"IMAGE REPOSITORY"`
	IfExists        *bool                  `ddl:"keyword" sql:"IF EXISTS"`
	name            SchemaObjectIdentifier `ddl:"identifier"`
}

// ShowImageRepositoryOptions is based on https://docs.snowflake.com/LIMITEDACCESS/snowpark-containers/reference/image-repository#show-image-repositories.
type ShowImageRepositoryOptions struct {
	show              bool  `ddl:"static" sql:"SHOW"`
	imageRepositories bool  `ddl:"static" sql:"IMAGE REPOSITORIES"`
	Like              *Like `ddl:"keyword" sql:"LIKE"`
	In                *In   `ddl:"keyword" sql:"IN"`
}

type imageRepositoryDBRow struct {
	CreatedOn     time.Time      `db:"created_on"`
	Name          string         `db:"name"`
	DatabaseName  string         `db:"database_name"`
	SchemaName    string         `db:"schema_name"`
	RepositoryUrl string         `db:"repository_url"`
	Owner         string         `db:"owner"`
	OwnerRoleType sql.NullString `db:"owner_role_type"`
	Comment       string         `db:"comment"`
}

type ImageRepository struct {
	CreatedOn     time.Time
	Name          string
	DatabaseName  string
	SchemaName    string
	RepositoryUrl string
	Owner         string
	OwnerRoleType string
	Comment       string
}
