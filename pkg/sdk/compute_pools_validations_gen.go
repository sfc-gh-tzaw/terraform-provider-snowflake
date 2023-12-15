package sdk

var (
	_ validatable = new(CreateComputePoolOptions)
	_ validatable = new(AlterComputePoolOptions)
	_ validatable = new(DropComputePoolOptions)
	_ validatable = new(ShowComputePoolOptions)
	_ validatable = new(DescribeComputePoolOptions)
)

func (opts *CreateComputePoolOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	if !ValidObjectIdentifier(opts.name) {
		errs = append(errs, ErrInvalidObjectIdentifier)
	}
	return JoinErrors(errs...)
}

func (opts *AlterComputePoolOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	if !ValidObjectIdentifier(opts.name) {
		errs = append(errs, ErrInvalidObjectIdentifier)
	}
	if !exactlyOneValueSet(opts.Suspend, opts.Resume, opts.StopAll, opts.Set, opts.Unset) {
		errs = append(errs, errExactlyOneOf("AlterComputePoolOptions", "Suspend", "Resume", "StopAll", "Set", "Unset"))
	}
	if valueSet(opts.Set) {
		if !anyValueSet(opts.Set.MinNodes, opts.Set.MaxNodes, opts.Set.AutoResume, opts.Set.AutoSuspendSecs, opts.Set.Comment) {
			errs = append(errs, errAtLeastOneOf("AlterComputePoolOptions.Set", "MinNodes", "MaxNodes", "AutoResume", "AutoSuspendSecs", "Comment"))
		}
	}
	if valueSet(opts.Unset) {
		if !anyValueSet(opts.Unset.AutoSuspendSecs, opts.Unset.AutoResume, opts.Unset.Comment) {
			errs = append(errs, errAtLeastOneOf("AlterComputePoolOptions.Unset", "AutoSuspendSecs", "AutoResume", "Comment"))
		}
	}
	return JoinErrors(errs...)
}

func (opts *DropComputePoolOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	if !ValidObjectIdentifier(opts.name) {
		errs = append(errs, ErrInvalidObjectIdentifier)
	}
	return JoinErrors(errs...)
}

func (opts *ShowComputePoolOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	return JoinErrors(errs...)
}

func (opts *DescribeComputePoolOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	if !ValidObjectIdentifier(opts.name) {
		errs = append(errs, ErrInvalidObjectIdentifier)
	}
	return JoinErrors(errs...)
}
