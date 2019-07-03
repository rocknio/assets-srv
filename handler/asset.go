package handler

import (
	"context"

	asset "github.com/rocknio/micro-assets/assets-srv/proto/asset"
)

type Asset struct{}

func (e *Asset) New(context.Context, *asset.Asset, *asset.Response) error {
	panic("implement me")
}

func (e *Asset) QueryByAssetName(context.Context, *asset.Asset, *asset.Response) error {
	panic("implement me")
}

func (e *Asset) UpdateByAssetId(context.Context, *asset.Asset, *asset.Response) error {
	panic("implement me")
}

func (e *Asset) DeleteByAssetId(context.Context, *asset.Asset, *asset.Response) error {
	panic("implement me")
}
