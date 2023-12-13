package sdk

import (
	"context"
	"time"
)

type ComputePools interface {
	Create(ctx context.Context, request *CreateComputePoolRequest) error
	Alter(ctx context.Context, request *AlterComputePoolRequest) error
	Drop(ctx context.Context, request *DropComputePoolRequest) error
	Show(ctx context.Context, request *ShowComputePoolRequest) ([]ComputePool, error)
	ShowByID(ctx context.Context, id SchemaObjectIdentifier) (*ComputePool, error)
	Describe(ctx context.Context, id SchemaObjectIdentifier) (*ComputePool, error)
}

// CreateComputePoolOptions is based on https://docs.snowflake.com/en/LIMITEDACCESS/snowpark-containers/reference/compute-pool#create-compute-pool.
type CreateComputePoolOptions struct {
	create             bool                   `ddl:"static" sql:"CREATE"`
	computePool        bool                   `ddl:"static" sql:"COMPUTE POOL"`
	IfNotExists        *bool                  `ddl:"keyword" sql:"IF NOT EXISTS"`
	name               SchemaObjectIdentifier `ddl:"identifier"`
	MinNodes           int                    `ddl:"parameter,no_quotes" sql:"MIN_NODES"`
	MaxNodes           int                    `ddl:"parameter,no_quotes" sql:"MAX_NODES"`
	InstanceFamily     string                 `ddl:"parameter,no_quotes" sql:"INSTANCE_FAMILY"`
	AutoResume         *bool                  `ddl:"parameter" sql:"AUTO_RESUME"`
	InitiallySuspended *bool                  `ddl:"parameter" sql:"INITIALLY_SUSPENDED"`
	AutoSuspendSecs    *int                   `ddl:"parameter,no_quotes" sql:"AUTO_SUSPEND_SECS"`
	Comment            *string                `ddl:"parameter,single_quotes" sql:"COMMENT"`
}

// AlterComputePoolOptions is based on https://docs.snowflake.com/en/LIMITEDACCESS/snowpark-containers/reference/compute-pool#alter-compute-pool.
type AlterComputePoolOptions struct {
	alter       bool                   `ddl:"static" sql:"ALTER"`
	computePool bool                   `ddl:"static" sql:"COMPUTE POOL"`
	IfExists    *bool                  `ddl:"keyword" sql:"IF EXISTS"`
	name        SchemaObjectIdentifier `ddl:"identifier"`
	Suspend     *bool                  `ddl:"keyword" sql:"SUSPEND"`
	Resume      *bool                  `ddl:"keyword" sql:"RESUME"`
	StopAll     *bool                  `ddl:"keyword" sql:"STOP ALL"`
	Set         *PropertiesToAlter     `ddl:"keyword" sql:"SET"`
}

type PropertiesToAlter struct {
	MinNodes        *int    `ddl:"parameter,no_quotes" sql:"MIN_NODES"`
	MaxNodes        *int    `ddl:"parameter,no_quotes" sql:"MAX_NODES"`
	AutoResume      *bool   `ddl:"parameter" sql:"AUTO_RESUME"`
	AutoSuspendSecs *int    `ddl:"parameter,no_quotes" sql:"AUTO_SUSPEND_SECS"`
	Comment         *string `ddl:"parameter,single_quotes" sql:"COMMENT"`
}

// DropComputePoolOptions is based on https://docs.snowflake.com/en/LIMITEDACCESS/snowpark-containers/reference/compute-pool#drop-compute-pool.
type DropComputePoolOptions struct {
	drop        bool                   `ddl:"static" sql:"DROP"`
	computePool bool                   `ddl:"static" sql:"COMPUTE POOL"`
	IfExists    *bool                  `ddl:"keyword" sql:"IF EXISTS"`
	name        SchemaObjectIdentifier `ddl:"identifier"`
}

// ShowComputePoolOptions is based on https://docs.snowflake.com/en/LIMITEDACCESS/snowpark-containers/reference/compute-pool#show-compute-pools.
type ShowComputePoolOptions struct {
	show         bool       `ddl:"static" sql:"SHOW"`
	computePools bool       `ddl:"static" sql:"COMPUTE POOLS"`
	Like         *Like      `ddl:"keyword" sql:"LIKE"`
	StartsWith   *string    `ddl:"parameter,single_quotes,no_equals" sql:"STARTS WITH"`
	Limit        *LimitFrom `ddl:"keyword" sql:"LIMIT"`
}

type computePoolDBRow struct {
	Name            string    `db:"name"`
	State           string    `db:"state"`
	MinNodes        int       `db:"min_nodes"`
	MaxNodes        int       `db:"max_nodes"`
	InstanceFamily  string    `db:"instance_family"`
	NumServices     int       `db:"num_services"`
	NumJobs         int       `db:"num_jobs"`
	AutoSuspendSecs int       `db:"auto_suspend_secs"`
	AutoResume      bool      `db:"auto_resume"`
	ActiveNodes     int       `db:"active_nodes"`
	IdleNodes       int       `db:"idle_nodes"`
	CreatedOn       time.Time `db:"created_on"`
	ResumedOn       time.Time `db:"resumed_on"`
	UpdatedOn       time.Time `db:"updated_on"`
	Owner           string    `db:"owner"`
	Comment         string    `db:"comment"`
}

type ComputePool struct {
	Name            string
	State           string
	MinNodes        int
	MaxNodes        int
	InstanceFamily  string
	NumServices     int
	NumJobs         int
	AutoSuspendSecs int
	AutoResume      bool
	ActiveNodes     int
	IdleNodes       int
	CreatedOn       time.Time
	ResumedOn       time.Time
	UpdatedOn       time.Time
	Owner           string
	Comment         string
}

// DescribeComputePoolOptions is based on https://docs.snowflake.com/en/LIMITEDACCESS/snowpark-containers/reference/compute-pool#describe-compute-pool.
type DescribeComputePoolOptions struct {
	describe    bool                   `ddl:"static" sql:"DESCRIBE"`
	computePool bool                   `ddl:"static" sql:"COMPUTE POOL"`
	name        SchemaObjectIdentifier `ddl:"identifier"`
}
