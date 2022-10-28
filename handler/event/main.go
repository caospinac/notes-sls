package main

import (
	"context"

	"github.com/caospinac/notes-sls/domain"
	"github.com/caospinac/notes-sls/util"
)

func event(ctx context.Context, event domain.EventRequest) (util.Response, util.ApiError) {
	res := util.NewResponse().WithBody(event)

	return res, nil
}

func main() {
	util.Start(event)
}
