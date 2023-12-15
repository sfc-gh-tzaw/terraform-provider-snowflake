package sdk

var (
	_ validatable = new(CreateImageRepositoryOptions)
	_ validatable = new(DropImageRepositoryOptions)
	_ validatable = new(ShowImageRepositoryOptions)
)

func (opts *CreateImageRepositoryOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	if !ValidObjectIdentifier(opts.name) {
		errs = append(errs, ErrInvalidObjectIdentifier)
	}
	return JoinErrors(errs...)
}

func (opts *DropImageRepositoryOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	if !ValidObjectIdentifier(opts.name) {
		errs = append(errs, ErrInvalidObjectIdentifier)
	}
	return JoinErrors(errs...)
}

func (opts *ShowImageRepositoryOptions) validate() error {
	if opts == nil {
		return ErrNilOptions
	}
	var errs []error
	return JoinErrors(errs...)
}
