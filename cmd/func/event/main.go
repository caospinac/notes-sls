package main

import (
	"context"
	"net/http"

	"github.com/caospinac/notes-sls/pkg/util"
)

func event(ctx context.Context, event util.EventRequest) (util.Response, util.ApiError) {
	res := util.NewResponse(http.StatusOK).WithBody(event)

	return res, nil
}

func main() {
	util.Start(event)
}
