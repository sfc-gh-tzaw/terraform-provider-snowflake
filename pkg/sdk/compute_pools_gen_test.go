package sdk

import (
	"fmt"
	"testing"
)

func TestComputePools_Create(t *testing.T) {
	id := RandomAccountObjectIdentifier()

	fmt.Println(id.FullyQualifiedName())
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
		opts.name = NewSchemaObjectIdentifier("", "", "")
		assertOptsInvalidJoinedErrors(t, opts, ErrInvalidObjectIdentifier)
	})

	t.Run("basic", func(t *testing.T) {
		opts := defaultOpts()
		assertOptsValidAndSQLEquals(t, opts, "CREATE COMPUTE POOL %s %s", id.FullyQualifiedName(), sql)
	})

	t.Run("all options", func(t *testing.T) {
		opts := defaultOpts()
		// TODO: fill me
		assertOptsValidAndSQLEquals(t, opts, "TODO: fill me")
	})
}

func TestComputePools_Alter(t *testing.T) {
	id := RandomSchemaObjectIdentifier()

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

	t.Run("validation: exactly one field from [opts.Suspend opts.Resume opts.StopAll opts.Set] should be present", func(t *testing.T) {
		opts := defaultOpts()
		// TODO: fill me
		assertOptsInvalidJoinedErrors(t, opts, errExactlyOneOf("AlterComputePoolOptions", "Suspend", "Resume", "StopAll", "Set"))
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
	id := RandomSchemaObjectIdentifier()

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
	id := RandomSchemaObjectIdentifier()

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
