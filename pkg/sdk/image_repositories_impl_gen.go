package sdk

import (
	"context"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/internal/collections"
)

var _ ImageRepositories = (*imageRepositories)(nil)

type imageRepositories struct {
	client *Client
}

func (v *imageRepositories) Create(ctx context.Context, request *CreateImageRepositoryRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *imageRepositories) Drop(ctx context.Context, request *DropImageRepositoryRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (v *imageRepositories) Show(ctx context.Context, request *ShowImageRepositoryRequest) ([]ImageRepository, error) {
	opts := request.toOpts()
	dbRows, err := validateAndQuery[imageRepositoryDBRow](v.client, ctx, opts)
	if err != nil {
		return nil, err
	}
	resultList := convertRows[imageRepositoryDBRow, ImageRepository](dbRows)
	return resultList, nil
}

func (v *imageRepositories) ShowByID(ctx context.Context, id SchemaObjectIdentifier) (*ImageRepository, error) {
	imageRepositories, err := v.Show(ctx, NewShowImageRepositoryRequest().WithLike(&Like{String(id.Name())}))
	if err != nil {
		return nil, err
	}
	return collections.FindOne(imageRepositories, func(r ImageRepository) bool { return r.Name == id.Name() })
}

func (r *CreateImageRepositoryRequest) toOpts() *CreateImageRepositoryOptions {
	opts := &CreateImageRepositoryOptions{
		OrReplace:   r.OrReplace,
		IfNotExists: r.IfNotExists,
		name:        r.name,
	}
	return opts
}

func (r *DropImageRepositoryRequest) toOpts() *DropImageRepositoryOptions {
	opts := &DropImageRepositoryOptions{
		IfExists: r.IfExists,
		name:     r.name,
	}
	return opts
}

func (r *ShowImageRepositoryRequest) toOpts() *ShowImageRepositoryOptions {
	opts := &ShowImageRepositoryOptions{
		Like: r.Like,
		In:   r.In,
	}
	return opts
}

func (r imageRepositoryDBRow) convert() *ImageRepository {
	imageRepository := ImageRepository{
		CreatedOn:     r.CreatedOn,
		Name:          r.Name,
		DatabaseName:  r.DatabaseName,
		SchemaName:    r.SchemaName,
		RepositoryUrl: r.RepositoryUrl,
		Owner:         r.Owner,
		Comment:       r.Comment,
	}

	if r.OwnerRoleType.Valid {
		imageRepository.OwnerRoleType = r.OwnerRoleType.String
	}

	return &imageRepository
}
