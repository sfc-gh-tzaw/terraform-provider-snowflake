package sdk

//go:generate go run ./dto-builder-generator/main.go

var (
	_ optionsProvider[CreateComputePoolOptions]   = new(CreateComputePoolRequest)
	_ optionsProvider[AlterComputePoolOptions]    = new(AlterComputePoolRequest)
	_ optionsProvider[DropComputePoolOptions]     = new(DropComputePoolRequest)
	_ optionsProvider[ShowComputePoolOptions]     = new(ShowComputePoolRequest)
	_ optionsProvider[DescribeComputePoolOptions] = new(DescribeComputePoolRequest)
)

type CreateComputePoolRequest struct {
	IfNotExists        *bool
	name               AccountObjectIdentifier // required
	MinNodes           int                     // required
	MaxNodes           int                     // required
	InstanceFamily     string                  // required
	AutoResume         *bool
	InitiallySuspended *bool
	AutoSuspendSecs    *int
	Comment            *string
}

type AlterComputePoolRequest struct {
	IfExists *bool
	name     AccountObjectIdentifier // required
	Suspend  *bool
	Resume   *bool
	StopAll  *bool
	Set      *AlterSetPropertiesRequest
	Unset    *AlterUnsetPropertiesRequest
}

type AlterSetPropertiesRequest struct {
	MinNodes        *int
	MaxNodes        *int
	AutoResume      *bool
	AutoSuspendSecs *int
	Comment         *string
}

type AlterUnsetPropertiesRequest struct {
	AutoSuspendSecs *bool
	AutoResume      *bool
	Comment         *bool
}

type DropComputePoolRequest struct {
	IfExists *bool
	name     AccountObjectIdentifier // required
}

type ShowComputePoolRequest struct {
	Like       *Like
	StartsWith *string
	Limit      *LimitFrom
}

type DescribeComputePoolRequest struct {
	name AccountObjectIdentifier // required
}
