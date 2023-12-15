package sdk

import "testing"

func TestImageRepositories_Create(t *testing.T) {
	id := RandomSchemaObjectIdentifier()

	// Minimal valid CreateImageRepositoryOptions
	defaultOpts := func() *CreateImageRepositoryOptions {
		return &CreateImageRepositoryOptions{
			name: id,
		}
	}

	t.Run("validation: nil options", func(t *testing.T) {
		var opts *CreateImageRepositoryOptions = nil
		assertOptsInvalidJoinedErrors(t, opts, ErrNilOptions)
	})

	t.Run("validation: valid identifier for [opts.name]", func(t *testing.T) {
		opts := defaultOpts()
		// TODO: fill me
		assertOptsInvalidJoinedErrors(t, opts, ErrInvalidObjectIdentifier)
	})

	t.Run("basic", func(t *testing.T) {
		opts := defaultOpts()
		// TODO: fill me
		assertOptsValidAndSQLEquals(t, opts, "TODO: fill me")
	})

	t.Run("all options", func(t *testing.T) {
		opts := defaultOpts()
		// TODO: fill me
		assertOptsValidAndSQLEquals(t, opts, "TODO: fill me")
	})
}

func TestImageRepositories_Drop(t *testing.T) {
	id := RandomSchemaObjectIdentifier()

	// Minimal valid DropImageRepositoryOptions
	defaultOpts := func() *DropImageRepositoryOptions {
		return &DropImageRepositoryOptions{
			name: id,
		}
	}

	t.Run("validation: nil options", func(t *testing.T) {
		var opts *DropImageRepositoryOptions = nil
		assertOptsInvalidJoinedErrors(t, opts, ErrNilOptions)
	})

	t.Run("validation: valid identifier for [opts.name]", func(t *testing.T) {
		opts := defaultOpts()
		// TODO: fill me
		assertOptsInvalidJoinedErrors(t, opts, ErrInvalidObjectIdentifier)
	})

	t.Run("basic", func(t *testing.T) {
		opts := defaultOpts()
		// TODO: fill me
		assertOptsValidAndSQLEquals(t, opts, "TODO: fill me")
	})

	t.Run("all options", func(t *testing.T) {
		opts := defaultOpts()
		// TODO: fill me
		assertOptsValidAndSQLEquals(t, opts, "TODO: fill me")
	})
}

func TestImageRepositories_Show(t *testing.T) {
	// Minimal valid ShowImageRepositoryOptions
	defaultOpts := func() *ShowImageRepositoryOptions {
		return &ShowImageRepositoryOptions{}
	}

	t.Run("validation: nil options", func(t *testing.T) {
		var opts *ShowImageRepositoryOptions = nil
		assertOptsInvalidJoinedErrors(t, opts, ErrNilOptions)
	})

	t.Run("basic", func(t *testing.T) {
		opts := defaultOpts()
		// TODO: fill me
		assertOptsValidAndSQLEquals(t, opts, "TODO: fill me")
	})

	t.Run("all options", func(t *testing.T) {
		opts := defaultOpts()
		// TODO: fill me
		assertOptsValidAndSQLEquals(t, opts, "TODO: fill me")
	})
}
