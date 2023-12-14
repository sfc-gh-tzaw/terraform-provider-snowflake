package sdk

import (
	"context"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/internal/collections"
)

var _ ComputePools = (*computePools)(nil)

type computePools struct {
	client *Client
}

func (v *computePools) Create(ctx context.Context, request *CreateComputePoolRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *computePools) Alter(ctx context.Context, request *AlterComputePoolRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *computePools) Drop(ctx context.Context, request *DropComputePoolRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *computePools) Show(ctx context.Context, request *ShowComputePoolRequest) ([]ComputePool, error) {
	opts := request.toOpts()
	dbRows, err := validateAndQuery[computePoolDBRow](v.client, ctx, opts)
	if err != nil {
		return nil, err
	}
	resultList := convertRows[computePoolDBRow, ComputePool](dbRows)
	return resultList, nil
}

func (v *computePools) ShowByID(ctx context.Context, id AccountObjectIdentifier) (*ComputePool, error) {
	// TODO: adjust request if e.g. LIKE is supported for the resource
	computePools, err := v.Show(ctx, NewShowComputePoolRequest().WithLike(&Like{String(id.Name())}))
	if err != nil {
		return nil, err
	}
	return collections.FindOne(computePools, func(r ComputePool) bool { return r.Name == id.Name() })
}

func (v *computePools) Describe(ctx context.Context, id AccountObjectIdentifier) (*ComputePool, error) {
	opts := &DescribeComputePoolOptions{
		name: id,
	}
	result, err := validateAndQueryOne[computePoolDBRow](v.client, ctx, opts)
	if err != nil {
		return nil, err
	}
	return result.convert(), nil
}

func (r *CreateComputePoolRequest) toOpts() *CreateComputePoolOptions {
	opts := &CreateComputePoolOptions{
		IfNotExists:        r.IfNotExists,
		name:               r.name,
		MinNodes:           r.MinNodes,
		MaxNodes:           r.MaxNodes,
		InstanceFamily:     r.InstanceFamily,
		AutoResume:         r.AutoResume,
		InitiallySuspended: r.InitiallySuspended,
		AutoSuspendSecs:    r.AutoSuspendSecs,
		Comment:            r.Comment,
	}
	return opts
}

func (r *AlterComputePoolRequest) toOpts() *AlterComputePoolOptions {
	opts := &AlterComputePoolOptions{
		IfExists: r.IfExists,
		name:     r.name,
		Suspend:  r.Suspend,
		Resume:   r.Resume,
		StopAll:  r.StopAll,
	}
	if r.Set != nil {
		opts.Set = &PropertiesToAlter{
			MinNodes:        r.Set.MinNodes,
			MaxNodes:        r.Set.MaxNodes,
			AutoResume:      r.Set.AutoResume,
			AutoSuspendSecs: r.Set.AutoSuspendSecs,
			Comment:         r.Set.Comment,
		}
	}
	return opts
}

func (r *DropComputePoolRequest) toOpts() *DropComputePoolOptions {
	opts := &DropComputePoolOptions{
		IfExists: r.IfExists,
		name:     r.name,
	}
	return opts
}

func (r *ShowComputePoolRequest) toOpts() *ShowComputePoolOptions {
	opts := &ShowComputePoolOptions{
		Like:       r.Like,
		StartsWith: r.StartsWith,
		Limit:      r.Limit,
	}
	return opts
}

func (r computePoolDBRow) convert() *ComputePool {
	// TODO: Mapping - maybe need to change bool <-> string
	computePool := ComputePool{
		Name:            r.Name,
		State:           r.State,
		MinNodes:        r.MinNodes,
		MaxNodes:        r.MaxNodes,
		InstanceFamily:  r.InstanceFamily,
		NumServices:     r.NumServices,
		NumJobs:         r.NumJobs,
		AutoSuspendSecs: r.AutoSuspendSecs,
		AutoResume:      r.AutoResume,
		ActiveNodes:     r.ActiveNodes,
		IdleNodes:       r.IdleNodes,
		CreatedOn:       r.CreatedOn,
		ResumedOn:       r.ResumedOn,
		UpdatedOn:       r.UpdatedOn,
		Owner:           r.Owner,
	}

	if r.Comment.Valid {
		computePool.Comment = r.Comment.String
	}

	return &computePool
}

func (r *DescribeComputePoolRequest) toOpts() *DescribeComputePoolOptions {
	opts := &DescribeComputePoolOptions{
		name: r.name,
	}
	return opts
}
