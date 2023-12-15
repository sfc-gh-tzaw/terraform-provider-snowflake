package sdk

import g "github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/poc/generator"

//go:generate go run ./poc/main.go

var imageRepositoryDbRow = g.DbStruct("imageRepositoryDBRow").
	Field("created_on", "time.Time").
	Field("name", "string").
	Field("database_name", "string").
	Field("schema_name", "string").
	Field("repository_url", "string").
	Field("owner", "string").
	Field("owner_role_type", "sql.NullString").
	Field("comment", "string")

var imageRepository = g.PlainStruct("ImageRepository").
	Field("CreatedOn", "time.Time").
	Field("Name", "string").
	Field("DatabaseName", "string").
	Field("SchemaName", "string").
	Field("RepositoryUrl", "string").
	Field("Owner", "string").
	Field("OwnerRoleType", "string").
	Field("Comment", "string")

var ImageRepositoriesDef = g.NewInterface(
	"ImageRepositories",
	"ImageRepository",
	g.KindOfT[SchemaObjectIdentifier](),
).
	CreateOperation(
		"https://docs.snowflake.com/LIMITEDACCESS/snowpark-containers/reference/image-repository#create-image-repository",
		g.NewQueryStruct("CreateImageRepository").
			Create().
			OrReplace().
			SQL("IMAGE REPOSITORY").
			IfNotExists().
			Name().
			WithValidation(g.ValidIdentifier, "name"),
	).
	DropOperation(
		"https://docs.snowflake.com/LIMITEDACCESS/snowpark-containers/reference/image-repository#drop-image-repository",
		g.NewQueryStruct("DropImageRepository").
			Drop().
			SQL("IMAGE REPOSITORY").
			IfExists().
			Name().
			WithValidation(g.ValidIdentifier, "name"),
	).
	ShowOperation(
		"https://docs.snowflake.com/LIMITEDACCESS/snowpark-containers/reference/image-repository#show-image-repositories",
		imageRepositoryDbRow,
		imageRepository,
		g.NewQueryStruct("ShowImageRepositories").
			Show().
			SQL("IMAGE REPOSITORIES").
			OptionalLike().
			OptionalIn(),
	).
	ShowByIdOperation()
