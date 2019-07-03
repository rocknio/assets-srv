package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	asset "github.com/rocknio/micro-assets/assets-srv/proto/asset"
)

type Asset struct{}

func (e *Asset) Handle(ctx context.Context, msg *asset.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *asset.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
