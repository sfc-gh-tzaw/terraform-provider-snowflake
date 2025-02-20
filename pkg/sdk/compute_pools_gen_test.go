package sdk

import "testing"

func TestComputePools_Create(t *testing.T) {
	id := RandomAccountObjectIdentifier()

	// Minimal valid CreateComputePoolOptions
	defaultOpts := func() *CreateComputePoolOptions {
		return &CreateComputePoolOptions{
			name:           id,
			MinNodes:       1,
			MaxNodes:       1,
			InstanceFamily: "standard_1",
		}
	}

	sql := "MIN_NODES = 1 MAX_NODES = 1 INSTANCE_FAMILY = standard_1"

	t.Run("validation: nil options", func(t *testing.T) {
		var opts *CreateComputePoolOptions = nil
		assertOptsInvalidJoinedErrors(t, opts, ErrNilOptions)
	})

	t.Run("validation: valid identifier for [opts.name]", func(t *testing.T) {
		opts := defaultOpts()
		opts.name = NewAccountObjectIdentifier("")
		assertOptsInvalidJoinedErrors(t, opts, ErrInvalidObjectIdentifier)
	})

	t.Run("basic", func(t *testing.T) {
		opts := defaultOpts()
		assertOptsValidAndSQLEquals(t, opts, "CREATE COMPUTE POOL %s %s", id.FullyQualifiedName(), sql)
	})

	// t.Run("all options", func(t *testing.T) {
	// 	opts := defaultOpts()
	// 	// TODO: fill me
	// 	assertOptsValidAndSQLEquals(t, opts, "TODO: fill me")
	// })
}

func TestComputePools_Alter(t *testing.T) {
	id := RandomAccountObjectIdentifier()

	// Minimal valid AlterComputePoolOptions
	defaultOpts := func() *AlterComputePoolOptions {
		return &AlterComputePoolOptions{
			name: id,
		}
	}

	t.Run("validation: nil options", func(t *testing.T) {
		var opts *AlterComputePoolOptions = nil
		assertOptsInvalidJoinedErrors(t, opts, ErrNilOptions)
	})

	t.Run("validation: valid identifier for [opts.name]", func(t *testing.T) {
		opts := defaultOpts()
		// TODO: fill me
		assertOptsInvalidJoinedErrors(t, opts, ErrInvalidObjectIdentifier)
	})

	t.Run("validation: exactly one field from [opts.Suspend opts.Resume opts.StopAll opts.Set opts.Unset] should be present", func(t *testing.T) {
		opts := defaultOpts()
		// TODO: fill me
		assertOptsInvalidJoinedErrors(t, opts, errExactlyOneOf("AlterComputePoolOptions", "Suspend", "Resume", "StopAll", "Set", "Unset"))
	})

	t.Run("validation: at least one of the fields [opts.Set.MinNodes opts.Set.MaxNodes opts.Set.AutoResume opts.Set.AutoSuspendSecs opts.Set.Comment] should be set", func(t *testing.T) {
		opts := defaultOpts()
		// TODO: fill me
		assertOptsInvalidJoinedErrors(t, opts, errAtLeastOneOf("AlterComputePoolOptions.Set", "MinNodes", "MaxNodes", "AutoResume", "AutoSuspendSecs", "Comment"))
	})

	t.Run("validation: at least one of the fields [opts.Unset.AutoSuspendSecs opts.Unset.AutoResume opts.Unset.Comment] should be set", func(t *testing.T) {
		opts := defaultOpts()
		// TODO: fill me
		assertOptsInvalidJoinedErrors(t, opts, errAtLeastOneOf("AlterComputePoolOptions.Unset", "AutoSuspendSecs", "AutoResume", "Comment"))
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

func TestComputePools_Drop(t *testing.T) {
	id := RandomAccountObjectIdentifier()

	// Minimal valid DropComputePoolOptions
	defaultOpts := func() *DropComputePoolOptions {
		return &DropComputePoolOptions{
			name: id,
		}
	}

	t.Run("validation: nil options", func(t *testing.T) {
		var opts *DropComputePoolOptions = nil
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

func TestComputePools_Show(t *testing.T) {
	// Minimal valid ShowComputePoolOptions
	defaultOpts := func() *ShowComputePoolOptions {
		return &ShowComputePoolOptions{}
	}

	t.Run("validation: nil options", func(t *testing.T) {
		var opts *ShowComputePoolOptions = nil
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

func TestComputePools_Describe(t *testing.T) {
	id := RandomAccountObjectIdentifier()

	// Minimal valid DescribeComputePoolOptions
	defaultOpts := func() *DescribeComputePoolOptions {
		return &DescribeComputePoolOptions{
			name: id,
		}
	}

	t.Run("validation: nil options", func(t *testing.T) {
		var opts *DescribeComputePoolOptions = nil
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
