package sdk

import g "github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/poc/generator"

//go:generate go run ./poc/main.go

var propertiesToAlter = g.NewQueryStruct("PropertiesToAlter").
	OptionalNumberAssignment("MIN_NODES", g.ParameterOptions().NoQuotes()).
	OptionalNumberAssignment("MAX_NODES", g.ParameterOptions().NoQuotes()).
	OptionalBooleanAssignment("AUTO_RESUME", g.ParameterOptions()).
	OptionalNumberAssignment("AUTO_SUSPEND_SECS", g.ParameterOptions().NoQuotes()).
	OptionalTextAssignment("COMMENT", g.ParameterOptions().SingleQuotes())

var computePoolDbRow = g.DbStruct("computePoolDbRow").
	Field("name", "string").
	Field("state", "string").
	Field("min_nodes", "int").
	Field("max_nodes", "int").
	Field("instance_family", "string").
	Field("num_services", "int").
	Field("num_jobs", "int").
	Field("auto_suspend_secs", "int").
	Field("auto_resume", "bool").
	Field("active_nodes", "int").
	Field("idle_nodes", "int").
	Field("created_on", "time.Time").
	Field("resumed_on", "time.Time").
	Field("updated_on", "time.Time").
	Field("owner", "string").
	Field("comment", "string")

var computePool = g.PlainStruct("ComputePool").
	Field("Name", "string").
	Field("State", "string").
	Field("MinNodes", "int").
	Field("MaxNodes", "int").
	Field("InstanceFamily", "string").
	Field("NumServices", "int").
	Field("NumJobs", "int").
	Field("AutoSuspendSecs", "int").
	Field("AutoResume", "bool").
	Field("ActiveNodes", "int").
	Field("IdleNodes", "int").
	Field("CreatedOn", "time.Time").
	Field("ResumedOn", "time.Time").
	Field("UpdatedOn", "time.Time").
	Field("Owner", "string").
	Field("Comment", "string")

var ComputePoolDef = g.NewInterface(
	"ComputePools",
	"ComputePool",
	g.KindOfT[SchemaObjectIdentifier](),
).CreateOperation(
	"https://docs.snowflake.com/en/LIMITEDACCESS/snowpark-containers/reference/compute-pool#create-compute-pool",
	g.NewQueryStruct("CreateComputePool").
		Create().
		SQL("COMPUTE POOL").
		IfNotExists().
		Name().
		NumberAssignment("MIN_NODES", g.ParameterOptions().NoQuotes().Required()).
		NumberAssignment("MAX_NODES", g.ParameterOptions().NoQuotes().Required()).
		TextAssignment("INSTANCE_FAMILY", g.ParameterOptions().NoQuotes().Required()).
		OptionalBooleanAssignment("AUTO_RESUME", g.ParameterOptions()).
		OptionalBooleanAssignment("INITIALLY_SUSPENDED", g.ParameterOptions()).
		OptionalNumberAssignment("AUTO_SUSPEND_SECS", g.ParameterOptions().NoQuotes()).
		OptionalTextAssignment("COMMENT", g.ParameterOptions().SingleQuotes()).
		WithValidation(g.ValidIdentifier, "name"),
).AlterOperation(
	"https://docs.snowflake.com/en/LIMITEDACCESS/snowpark-containers/reference/compute-pool#alter-compute-pool",
	g.NewQueryStruct("AlterComputePool").
		Alter().
		SQL("COMPUTE POOL").
		IfExists().
		Name().
		OptionalSQL("SUSPEND").
		OptionalSQL("RESUME").
		OptionalSQL("STOP ALL").
		OptionalQueryStructField(
			"Set",
			propertiesToAlter,
			g.KeywordOptions().SQL("SET"),
		).
		WithValidation(g.ValidIdentifier, "name").
		WithValidation(g.ExactlyOneValueSet, "Suspend", "Resume", "StopAll", "Set"),
).
	DropOperation(
		"https://docs.snowflake.com/en/LIMITEDACCESS/snowpark-containers/reference/compute-pool#drop-compute-pool",
		g.NewQueryStruct("DropComputePool").
			Drop().
			SQL("COMPUTE POOL").
			IfExists().
			Name().
			WithValidation(g.ValidIdentifier, "name"),
	).
	ShowOperation(
		"https://docs.snowflake.com/en/LIMITEDACCESS/snowpark-containers/reference/compute-pool#show-compute-pools",
		computePoolDbRow,
		computePool,
		g.NewQueryStruct("ShowComputePools").
			Show().
			SQL("COMPUTE POOLS").
			OptionalLike().
			OptionalStartsWith().
			OptionalLimit(),
	).
	ShowByIdOperation().
	DescribeOperation(
		g.DescriptionMappingKindSingleValue,
		"https://docs.snowflake.com/en/LIMITEDACCESS/snowpark-containers/reference/compute-pool#describe-compute-pool",
		computePoolDbRow,
		computePool,
		g.NewQueryStruct("DescribeComputePool").
			Describe().
			SQL("COMPUTE POOL").
			Name().
			WithValidation(g.ValidIdentifier, "name"),
	)
